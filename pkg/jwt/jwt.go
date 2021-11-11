package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenExpireDuration = 2 * time.Hour
)

var (
	mySecret = []byte(string("🐏杨承翰干爆🐖朱涛涛"))
)

// MyClaims 声明自定义结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含官方字段
// 添加一个额外的username字段，所以要自定义结构体
type MyClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenToken 生成token
// 返回Token err
func GenToken(userId int64, username string) (string, error) {
	c := MyClaims{
		UserId:   userId,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 返回一个使用指定签名并获得完整编码后的字符串Token
	return token.SignedString(mySecret)
}

// ParseToken 解析token string
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	// 将token string解析成MyClaim类型
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	// 解析错误
	if err != nil {
		return nil, err
	}
	//校验token
	if token.Valid {
		return mc, nil
	}
	// 解析失败
	return nil, errors.New("invalid token")
}
