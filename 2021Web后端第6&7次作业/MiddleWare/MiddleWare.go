package MiddleWare

import (
	"LV2+LV3/ControlPanel"
	"LV2+LV3/Mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)
 //RETURN 接受所有的错误类型
type RETURN struct {
	err error
	str string
}
var R RETURN

func Register(c *gin.Context){        //注册账号
	var u Mysql.User
	R.err = c.ShouldBind(&u)
	if R.err == nil{
		R.str = ControlPanel.Contrast(u)
		if R.str == "注册成功"{
			Mysql.UForAll = append(Mysql.UForAll,u)  //同步信息
			Mysql.Sum++
			R.err = Mysql.WriteIn(u)
		}
	}
	c.JSON(http.StatusOK,gin.H{      //返回相关数据
		"Result":R.str,
	})
}

func LogIn(c *gin.Context){                  //登录账号
	var u Mysql.User
	R.err = c.ShouldBind(&u)
	if R.err == nil{
		R.str = ControlPanel.ToLogIn(u)        //对比账号和密码
	}
	c.JSON(http.StatusOK,gin.H{         //返回相关数据
		"Result":R.str,
	})
}

func AnswerQ(c *gin.Context){
	var u Mysql.User
	//给前端传输问题
	c.JSON(200,gin.H{
		"question1":ControlPanel.U.Question1,
		"question2":ControlPanel.U.Question2,
		"question3":ControlPanel.U.Question3,
	})   //若返回的是空，则前端不做输出
	_ = c.ShouldBind(&u)               //获得输入
	R.str=ControlPanel.ToAnswerQ(u)               //比对正确性
	c.JSON(200,gin.H{
		"Result":R.str,
	})
}

func ChangePassword(c *gin.Context){     //更改密码
	var u Mysql.User
	R.err = c.ShouldBind(&u)             //获取修改后的密码
	if R.err==nil{
		R.err,R.str = ControlPanel.ToChangePassword(u) //检验修改后的密码并储存
		c.JSON(http.StatusOK,gin.H{
			"Result":R.str,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"Result":R.err,
		})
	}
}
func SetQuestions(c *gin.Context){       //创建密保
	var u Mysql.User
	R.err = c.ShouldBind(&u)                //获取密保
	if R.err==nil{
		R.err,R.str = ControlPanel.ToSetQuestions(u)      //储存问题和答案
	}
	if R.err!=nil{
		c.JSON(200,gin.H{
			"Result":R.err,
		})
	}else{
		c.JSON(200,gin.H{
			"Result":R.str,
		})
	}
}
func PostComments(c *gin.Context){      //发送留言
	var co Mysql.Comment
	_ = c.ShouldBind(&co)                        //前端返回 相应留言的order或者他人的用户名receiver,和发送的内容(content)
	R.err,R.str = ControlPanel.ToPostComments(co)    //找到目标 写入表格
	if R.err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"result":R.err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"result":R.str,
		})
	}
}
func ChangeComments(c *gin.Context){    //更改留言
	var co Mysql.Comment
	_ = c.ShouldBind(&co)                            //获取需修改的句子的id,修改后的句子content
	R.str=ControlPanel.ToChangeComments(co)	         //找到特定的留言并更新那一条留言再同步留言
	c.JSON(200,gin.H{
		"Result":R.str,
	})
}
//DeleteComments :删除留言
func DeleteComments(c *gin.Context){
	var co Mysql.Comment
	_ = c.ShouldBind(&co)                             //获取order
	R.str = ControlPanel.ToDeleteComments(co)
	c.JSON(http.StatusOK,gin.H{
		"Result":R.str,
	})
}

func ShowComments(c *gin.Context){       //显示留言(只能显示自己host_name的)
	for i:=0;i<Mysql.CSum;i++ {
		x:=ControlPanel.ToShowComments(i)        //直接显示留言
		c.JSON(200,gin.H{
			"Layer":x.Layer,
			"id":x.Id,
			"FatherId":x.FatherId,
			"Poster":x.Poster,
			"Receiver":x.Receiver,
			"Content":x.Content,
		})
	}

}

