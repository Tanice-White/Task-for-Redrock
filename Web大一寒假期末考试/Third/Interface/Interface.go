package Interface

import (
	"Third/global"
	"Third/handlers"

	"github.com/gin-gonic/gin"
)

func StartEngines(){
	global.Engines = gin.Default()

	global.Engines.GET("/help",handlers.GetHelp)                   //帮助
	global.Engines.POST("/register",handlers.Register)             //注册(name account password 均必写)
	global.Engines.POST("/logIn",handlers.LogIn)                   //登录(account password 均必写)

	UserGroup := global.Engines.Group("/user")
	{                                                                         //先登录一次再进行以下操作

		UserGroup.POST("/reset_password",handlers.ChangePassword)   //修改密码(输入更改的密码password)
		UserGroup.POST("/charge",handlers.Charge)                   //充值入口(输入金额money)
		UserGroup.POST("/transfer",handlers.Transfer)               //转账入口(输入对方账号account,金额money)
	}
	_ = global.Engines.Run()
}
