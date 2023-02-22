package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtMod "github.com/golang-jwt/jwt/v4"
	"go-yao/common/helpers"
	"go-yao/pkg/global"
	"go-yao/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	SignKey    []byte        // 秘钥
	MaxRefresh time.Duration // 刷新token的最大过期时间
}

// UserInfo 自定义用户信息
type UserInfo struct {
	UserID uint64 `json:"user_id"`
}

// CustomJWTClaims 自定义Payload信息
type CustomJWTClaims struct {
	UserInfo
	ExpireAtTime int64 `json:"expire_time"` // 过期时间
	jwtMod.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(global.Conf.JWT.Secret),
		MaxRefresh: time.Duration(global.Conf.JWT.MaxRefreshTime) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件中调用
func (j *JWT) ParserToken(ctx *gin.Context) (*CustomJWTClaims, error) {
	// 从header获取token
	tokenString, err := j.getTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// 解析token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtMod.ValidationError)
		if ok {
			if validationErr.Errors == jwtMod.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtMod.ValidationErrorExpired {
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

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (j *JWT) RefreshToken(ctx *gin.Context) (string, error) {
	// 1. 从 Header 里获取 token
	tokenString, parseErr := j.getTokenFromHeader(ctx)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		// 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
		validationErr, ok := err.(*jwtMod.ValidationError)

		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtMod.ValidationErrorExpired {
			return "", err
		}
	}

	// 3. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*CustomJWTClaims)

	// 4. 检查是否过了『最大允许刷新的时间』
	t := helpers.TimenowInTimezone().Add(-j.MaxRefresh).Unix()
	// 首次签名时间 > (当前时间 - 最大允许刷新时间)
	if claims.IssuedAt.Unix() > t {
		claims.RegisteredClaims.ExpiresAt = jwtMod.NewNumericDate(j.expireAtTime())
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成  Token，在登录成功时调用
func (j *JWT) IssueToken(info UserInfo) string {
	// 构造自定义Payload信息
	expireTime := j.expireAtTime()
	claims := CustomJWTClaims{
		// 用户信息
		UserInfo: UserInfo{
			UserID: info.UserID,
		},

		// 过期时间
		ExpireAtTime: expireTime.Unix(),
		RegisteredClaims: jwtMod.RegisteredClaims{
			NotBefore: jwtMod.NewNumericDate(helpers.TimenowInTimezone()), // 签名生效时间
			IssuedAt:  jwtMod.NewNumericDate(helpers.TimenowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtMod.NewNumericDate(expireTime),                  // 签名过期时间
			Issuer:    global.Conf.Application.Name,                       // 签名颁发者
		},
	}

	// 根据 claims 生成token对象
	token, err := j.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (j *JWT) createToken(claims CustomJWTClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwtMod.NewWithClaims(jwtMod.SigningMethodHS256, claims)
	return t.SignedString(j.SignKey)
}

// token过期时间
func (j *JWT) expireAtTime() time.Time {
	timezone := helpers.TimenowInTimezone()
	expireTime := global.Conf.JWT.ExpireTime

	expire := time.Duration(expireTime) * time.Minute
	return timezone.Add(expire)
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

// parseTokenString 解析token
func (j *JWT) parseTokenString(token string) (*jwtMod.Token, error) {
	return jwtMod.ParseWithClaims(token, &CustomJWTClaims{}, func(token *jwtMod.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}
