package api

import (
	"fast_gin/api/captcha_api"
	"fast_gin/api/image_api"
	"fast_gin/api/user_api"
)

var App = new(Api)

type Api struct {
	UserApi    user_api.UserApi
	ImageApi   image_api.ImageApi
	CaptchaApi captcha_api.CaptchaApi
}
