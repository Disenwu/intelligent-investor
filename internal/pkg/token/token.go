package token

import (
	"fmt"
	apiserver "intelligent-investor/cmd/api-server/options"
	"sync"
	
	"time"
	
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var (
	opts = &apiserver.ServerOptions{
		JWTSecret:  "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		Expiration: 2 * time.Hour,
	}
	once sync.Once
)

func Init() {
	once.Do(func() {
		fmt.Println("Start token config init...")
		jwtSecret := viper.GetString("jwt-secret")
		if jwtSecret != "" {
			opts.JWTSecret = jwtSecret
		}
		expiration := viper.GetDuration("expiration")
		if expiration != 0 {
			opts.Expiration = expiration
		}
	})
}

func CreateToken(username string) (string, time.Time, error) {
	expireAt := time.Now().Add(opts.Expiration)
	// 生成 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
		"exp":      expireAt.Unix(),
		"iat":      time.Now().Unix(),
	})
	
	tokenString, err := token.SignedString([]byte(opts.JWTSecret))
	if err != nil {
		return "", time.Time{}, err
	}
	// 签名并返回令牌
	return tokenString, expireAt, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(opts.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
