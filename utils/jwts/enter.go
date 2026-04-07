package jwts

import (
	"errors"
	"fast_gin/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID uint `json:"userID"`
	RoleID int8 `json:"roleID"`
}

type MyClaims struct {
	Claims
	jwt.RegisteredClaims
}

// SetToken 生成token
func SetToken(data Claims) (string, error) {
	SerClaims := MyClaims{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), //有效时间
			//IssuedAt:  jwt.NewNumericDate(time.Now()),                                                         //签发时间
			//NotBefore: jwt.NewNumericDate(time.Now()),                                                         //生效时间
			Issuer: global.Config.Jwt.Issuer, //签发人
			//Subject:   "somebody",                                                                             //主题
			//ID:        "1",                                                                                    //jwt id 核实jwt
			//Audience:  []string{"somebody"},                                                                   //用户
		},
	}

	//使用指定的加密方式和声明类型创建令牌
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, SerClaims)
	//获得完整令牌
	token, err := tokenStruct.SignedString([]byte(global.Config.Jwt.Key))
	if err != nil {
		logrus.Errorf("获取token令牌错误: %s", err)
		return "", err
	}
	return token, nil
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Key), nil
	})
	if err != nil {
		logrus.Errorf("验证token令牌错误: %s", err)
		return nil, err
	}
	if claims, ok := tokenObj.Claims.(*MyClaims); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}
}
