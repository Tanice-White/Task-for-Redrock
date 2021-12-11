package Interface

import (
	"LV2+LV3/MiddleWare"
	"github.com/gin-gonic/gin"
)

var engines *gin.Engine

func StartEngines(){
	engines = gin.Default()

	engines.POST("/Register",MiddleWare.Register) //注册(name account password 均必写)
	engines.POST("/LogIn",MiddleWare.LogIn)       //登录(account password 均必写)
	engines.POST("/AnswerQ",MiddleWare.AnswerQ)    //若回答正确则显示密码(手动输入账号account)

	UserGroup := engines.Group("/user")
	{
		//先登录一次再进行以下操作!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

		UserGroup.POST("/Reset_Password",MiddleWare.ChangePassword)   //修改密码(输入更改的密码password)
		UserGroup.POST("Reset_Questions",MiddleWare.SetQuestions)     //设置密保(密保1:question1,answer1...最多三个)）
		UserGroup.POST("/Post_Comments",MiddleWare.PostComments)    //发布留言(输入对方的留言number，用户名Receiver和你的留言)
		UserGroup.POST("/Change_Comments",MiddleWare.ChangeComments)//修改留言(输入对方的留言number和你的留言)
		UserGroup.POST("/Delete_Comments",MiddleWare.DeleteComments) //删除留言(手动输入需删除留言的序号number)
		UserGroup.GET("/Show_Comments",MiddleWare.ShowComments)        //查看所有留言
	}
	_ = engines.Run(":8080")
}