package middleware

import (
	"fast_gin/service/redis_ser"
	"fast_gin/utils/jwts"
	"fast_gin/utils/res"

	"github.com/gin-gonic/gin"
)

// jwt认证
func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	_, err := jwts.CheckToken(token)
	if err != nil {
		res.FailWithMig("认证失败", c)
		c.Abort()
		return
	}
	if redis_ser.HasLogout(token) {
		res.FailWithMig("当前登录已注销", c)
		c.Abort()
		return
	}
	c.Next()
}

// jwt角色认证
func AdminMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	cliams, err := jwts.CheckToken(token)
	if err != nil {
		res.FailWithMig("认证失败", c)
		c.Abort()
		return
	}
	if redis_ser.HasLogout(token) {
		res.FailWithMig("当前登录已注销", c)
		c.Abort()
		return
	}
	if cliams.RoleID != 1 {
		res.FailWithMig("角色认证失败", c)
		c.Abort()
		return
	}
	c.Set("claims", cliams)
	c.Next()
}

func GetAuth(c *gin.Context) (cl *jwts.MyClaims) {
	cl = new(jwts.MyClaims)
	_claims, ok := c.Get("claims")
	if !ok {
		return cl
	}
	cl, ok = _claims.(*jwts.MyClaims)
	return
}
