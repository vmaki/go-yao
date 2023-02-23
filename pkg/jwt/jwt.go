package jwt

import (
	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v4"
	"go-yao/common/helpers"
	"go-yao/common/response"
	"go-yao/pkg/global"
	"go-yao/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           = response.New(response.CodeTokenExpired)
	ErrTokenExpiredMaxRefresh = response.New(response.CodeTokenExpiredMaxRefresh)
	ErrTokenMalformed         = response.New(response.CodeTokenMalformed)
	ErrTokenInvalid           = response.New(response.CodeTokenInvalid)
	ErrHeaderEmpty            = response.New(response.CodeHeaderEmpty)
	ErrHeaderMalformed        = response.New(response.CodeHeaderMalformed)
)

type JWT struct {
	SignKey    []byte        // 秘钥
	MaxRefresh time.Duration // 刷新token的最大过期时间
}

// CustomJWTClaims 自定义Payload信息
type CustomJWTClaims struct {
	UserInfo
	ExpireAtTime int64
	jwtLib.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(global.Conf.JWT.Secret),
		MaxRefresh: time.Duration(global.Conf.JWT.MaxRefreshTime) * time.Second,
	}
}

// expireAtTime token 的过期时间
func (j *JWT) expireAtTime() time.Time {
	timezone := helpers.TimeNowInTimezone()
	expireTime := global.Conf.JWT.ExpireTime

	return timezone.Add(time.Duration(expireTime) * time.Second)
}

// IssueToken 生成 token，在登录成功时调用
func (j *JWT) IssueToken(info UserInfo) string {
	expireTime := j.expireAtTime()
	claims := CustomJWTClaims{
		UserInfo: UserInfo{
			UserID: info.UserID,
		},
		ExpireAtTime: expireTime.Unix(),
		RegisteredClaims: jwtLib.RegisteredClaims{
			NotBefore: jwtLib.NewNumericDate(helpers.TimeNowInTimezone()), // 签名生效时间
			IssuedAt:  jwtLib.NewNumericDate(helpers.TimeNowInTimezone()), // 首次签名时间（后续刷新 token 不会更新）
			ExpiresAt: jwtLib.NewNumericDate(expireTime),                  // 签名过期时间
			Issuer:    global.Conf.Application.Name,                       // 签名颁发者
		},
	}

	// 根据 claims 生成 token 对象
	token, err := j.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// createToken 创建 token，内部使用，外部请调用 IssueToken
func (j *JWT) createToken(claims CustomJWTClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	return t.SignedString(j.SignKey)
}

func (j *JWT) getTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}

// parseTokenString 解析 token
func (j *JWT) parseTokenString(token string) (*jwtLib.Token, error) {
	return jwtLib.ParseWithClaims(token, &CustomJWTClaims{}, func(token *jwtLib.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

// ParserToken 解析 token，中间件中调用
func (j *JWT) ParserToken(ctx *gin.Context) (*CustomJWTClaims, error) {
	// 从header获取token
	tokenString, err := j.getTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// 解析token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtLib.ValidationError)
		if ok {
			if validationErr.Errors == jwtLib.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtLib.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 token
func (j *JWT) RefreshToken(ctx *gin.Context) (string, error) {
	// 1. 从 Header 里获取 token
	tokenString, parseErr := j.getTokenFromHeader(ctx)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		// 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
		validationErr, ok := err.(*jwtLib.ValidationError)

		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtLib.ValidationErrorExpired {
			return "", err
		}
	}

	// 3. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*CustomJWTClaims)

	// 4. 检查是否过了『最大允许刷新的时间』
	t := helpers.TimeNowInTimezone().Add(-j.MaxRefresh).Unix()
	// 首次签名时间 > (当前时间 - 最大允许刷新时间)
	if claims.IssuedAt.Unix() > t {
		claims.RegisteredClaims.ExpiresAt = jwtLib.NewNumericDate(j.expireAtTime())
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}
