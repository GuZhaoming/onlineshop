package router

import (
	"github.com/gin-gonic/gin"
	"onlineshop-api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("sent_sms", api.SendSms)
	}
}
