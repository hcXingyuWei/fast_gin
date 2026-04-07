package middleware

import (
	"fast_gin/utils/res"
	"time"

	"github.com/gin-gonic/gin"
)

func LimitMiddleware(limit int) gin.HandlerFunc {
	return NewLimiter(limit, time.Second).Middleware
}

type Limiter struct {
	limit      int
	duration   time.Duration
	timestamps map[string][]int64
}

func NewLimiter(limit int, duration time.Duration) *Limiter {
	return &Limiter{
		limit:      limit,
		duration:   duration,
		timestamps: make(map[string][]int64),
	}
}

func (l *Limiter) Middleware(c *gin.Context) {
	ip := c.ClientIP() //获取客户端ip

	//检查请求时间戳是否存在
	if _, ok := l.timestamps[ip]; !ok {
		l.timestamps[ip] = make([]int64, 0)
	}

	now := time.Now().Unix() //获取当前时间戳

	//移除过期请求时间戳
	for i := 0; i < len(l.timestamps[ip]); i++ {
		if l.timestamps[ip][i] < now-int64(l.duration.Seconds()) {
			l.timestamps[ip] = append(l.timestamps[ip][:i], l.timestamps[ip][i+1:]...)
			i--
		}
	}

	//检查请求数量是否超过限制
	if len(l.timestamps[ip]) >= l.limit {
		res.FailWithMig("Too many requests", c)
		c.Abort()
		return
	}

	//添加当前请求时间戳到切片
	l.timestamps[ip] = append(l.timestamps[ip], now)

	// 继续处理请求
	c.Next()
}
