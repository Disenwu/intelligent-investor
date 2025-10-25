package response

import (
	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = 200
	CodeFailed  = 500
)

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	RequestID string      `json:"requestId"`
}

type PageResponse struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
}

func RequestId(ctx *gin.Context) string {
	requestID := ctx.Request.Context().Value("X-Request-ID")
	return requestID.(string)
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(CodeSuccess, Response{
		Code:      CodeSuccess,
		Data:      data,
		Msg:       "Success",
		RequestID: RequestId(ctx),
	})
}

func PageSuccess(ctx *gin.Context, list interface{}, total int64, page, pageSize int) {
	pageResponse := PageResponse{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		List:     list,
	}
	ctx.JSON(CodeSuccess, Response{
		Code:      CodeSuccess,
		Data:      pageResponse,
		Msg:       "Success",
		RequestID: RequestId(ctx),
	})
}
