package handler

import (
	"encoding/json"
	"fmt"
	"intelligent-investor/internal/app/model"
	store "intelligent-investor/internal/app/store/mysql"
	errors "intelligent-investor/internal/pkg/errors"
	"intelligent-investor/internal/pkg/log"
	"intelligent-investor/internal/pkg/response"
	redisService "intelligent-investor/internal/pkg/service"
	"intelligent-investor/internal/pkg/token"
	"intelligent-investor/pkg/encrypt"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @Description 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body model.LoginRequest true "登录请求"
// @Success 200 {object} response.Response{data=nil} "登录成功"
// @Failure 400 {object} response.Response{data=nil} "请求参数错误"
// @Failure 401 {object} response.Response{data=nil} "用户名或密码错误"
// @Router /user/login [post]
func LoginHandler(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.Errorw("Invalid login params", "error", err)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Invalid login params"))
	}
	log.Debugw("LoginRequest", "req", req)
	user, error := store.GetUserByUsername(req.Username)
	if error != nil {
		log.Errorw("User not found", "error", error)
		ctx.JSON(response.CodeFailed, errors.ErrPageNotFound.WithMessage("User not found").KeyAndValues("username", req.Username))
	}
	// 校验密码
	if err := encrypt.ComparePassword(user.Password, req.Password); err != nil {
		log.Errorw("Invalid password", "username", req.Username)
		ctx.JSON(response.CodeFailed, errors.ErrAuthorizationFailed.WithMessage("Invalid password").KeyAndValues("password", req.Password))
	}
	// 生成 JWT 令牌
	tokenString, expireAt, err := token.CreateToken(req.Username)
	if err != nil {
		log.Errorw("Failed to create token", "error", err)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Failed to create token"))
	}
	log.Infow("Login Success", "username", req.Username)
	// 缓存 JWT 令牌
	loginResponse := model.LoginResponse{
		Username: req.Username,
		Token:    tokenString,
		ExpireAt: expireAt,
	}
	jsonData, _ := json.Marshal(loginResponse)
	if err := redisService.SetAndExpire(fmt.Sprintf("token:%s", tokenString), jsonData, time.Until(expireAt)); err != nil {
		log.Errorw("Failed to cache token", "error", err)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Failed to cache token"))
	}
	response.Success(ctx, loginResponse)
}

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body model.RegisterRequest true "注册请求"
// @Success 200 {object} response.Response{data=nil} "注册成功"
// @Failure 400 {object} response.Response{data=nil} "请求参数错误"
// @Failure 409 {object} response.Response{data=nil} "用户名已存在"
// @Router /user/register [put]
func RegisterHandler(ctx *gin.Context) {
	var req model.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.Errorw("Invalid register params", "error", err)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Invalid register params"))
	}
	log.Debugw("RegisterRequest", "req", req)
	// 校验用户名是否已存在
	if _, err := store.GetUserByUsername(req.Username); err == nil {
		log.Errorw("Username already exists", "username", req.Username)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Username already exists").KeyAndValues("username", req.Username))
	}
	// 校验邮箱是否已存在
	if _, err := store.GetUserByEmail(req.Email); err == nil {
		log.Errorw("Email already exists", "email", req.Email)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Email already exists").KeyAndValues("email", req.Email))
	}
	// 加密密码
	password, err := encrypt.EncryptPassword(req.Password)
	if err != nil {
		log.Errorw("Failed to encrypt password", "error", err)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Failed to encrypt password"))
	}
	// 创建用户
	user := &model.User{
		Username: req.Username,
		Password: password,
		Email:    req.Email,
	}
	if err := store.CreateUser(user); err != nil {
		log.Errorw("Failed to create user", "error", err)
		ctx.JSON(response.CodeFailed, errors.ErrParamInvalid.WithMessage("Failed to create user"))
	}
	response.Success(ctx, nil)
}
