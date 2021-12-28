package handlers

import (
	"Third/dao"
	"Third/global"
	"Third/tools"
	"github.com/gin-gonic/gin"

	"net/http"
)

func GetHelp(c *gin.Context){                          //获取帮助
	c.JSON(200,gin.H{
		"帮助":"127.0.0.1:8080/help",
		"注册":"127.0.0.1:8080/register",
		"登录":"127.0.0.1:8080/logIn",
		"密码修改":"127.0.0.1:8080/reset_password",
		"充值":"127.0.0.1:8080/charge",
		"转账":"127.0.0.1:8080/transfer",
	})
}

func Register(c *gin.Context){                          //注册账号
	var u global.User
	global.R.Err= c.ShouldBind(&u)                     //需要得到用户的用户名、账号、密码
	if global.R.Err == nil{
		global.R.Str = tools.CanRegister(u)
		if global.R.Str == "注册成功"{
			global.AllUsers = append(global.AllUsers,u)  //同步信息
			global.Sum++
			global.R.Err = dao.WriteIn(u)
		}
	}
	c.JSON(http.StatusOK,gin.H{                          //返回相关数据
		"Result":global.R.Str,
	})
}

func LogIn(c *gin.Context){                              //登录账号
	var u global.User
	global.R.Err = c.ShouldBind(&u)                      //获取账号和密码
	if global.R.Err == nil{
		global.R.Str = tools.CanLogIn(u)                 //对比账号和密码
	}
	c.JSON(http.StatusOK,gin.H{                          //返回相关数据
		"Result":global.R.Str,
	})
}

func ChangePassword(c *gin.Context){                        //更改密码
	var u global.User
	global.R.Err = c.ShouldBind(&u)                         //获取修改后的密码
	if global.R.Err==nil{
		global.R.Err,global.R.Str = dao.HasChangePassword(u) //检验修改后的密码并储存
		c.JSON(http.StatusOK,gin.H{
			"Result":global.R.Str,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"Result":global.R.Err,
		})
	}
}

func Charge(c *gin.Context){
	var u global.User                                          //充值的金额
	global.R.Err = c.ShouldBind(&u)
	global.R.Err,global.R.Str = dao.HasCharge(u.Money)
	if global.R.Err==nil{
		c.JSON(http.StatusOK,gin.H{
			"Result":global.R.Str,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"Result":global.R.Err,
		})
	}
}

func Transfer(c *gin.Context){
	var u global.User                                          //转出的金额和转入的对象
	global.R.Err = c.ShouldBind(&u)
	global.R.Err,global.R.Str = dao.HasTransfer(u)
	if global.R.Err==nil{
		c.JSON(http.StatusOK,gin.H{
			"Result":global.R.Str,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"Result":global.R.Err,
		})
	}
}