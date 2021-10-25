package route

import (
	"fmt"
	"server/mysql"
	"server/util"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
)

type registertype struct {
	Phone     string `form:"phone" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Invitcode string `form:"invitcode binding:"required""`
	Code      string `form:"code" binding:"required`
}
type sendsms struct {
	Phone string `form:"phone" binding:"required"`
}
type logintype struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Test(context *gin.Context) {
	istrue := mysql.Selectinvit("8032")
	fmt.Println(istrue)
	context.JSON(200, gin.H{
		"code": 500,
		"msg":  "验证码不存在",
	})
}
func Sendsms(context *gin.Context) {
	var getdata sendsms
	err := context.ShouldBind(&getdata)
	if err != nil {
		fmt.Println(err)
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
		return
	}
	AccessKeyId := tea.String("LTAI5t9NHsixMxxkCPL6sZTr")
	AccessKeySecret := tea.String("nWgOAWjgkngQFQICo9z7hYvlsu2E5E")
	network := tea.String("dysmsapi.aliyuncs.com")
	PhoneNumbers := tea.String(getdata.Phone)
	SignName := tea.String("灰色Bird")
	TemplateCode := tea.String("SMS_225986913")
	smscode := util.RandNum(4)
	TemplateParam := `{"code": ` + smscode + `}`
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: AccessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = network
	// _result = &dysmsapi20170525.Client{}
	client, _err := dysmsapi20170525.NewClient(config)
	if _err != nil {
		panic(_err)
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  PhoneNumbers,
		SignName:      SignName,
		TemplateCode:  TemplateCode,
		TemplateParam: &TemplateParam,
	}

	// 复制代码运行请自行打印 API 的返回值
	sendResp, _err := client.SendSms(sendSmsRequest)
	if _err != nil {
		panic(_err)
	}
	code := sendResp.Body.Code
	if *code != "OK" {
		context.JSON(200, gin.H{
			"code": 500,
			"smg":  "发送失败",
		})
		panic(_err)
		return
	}
	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "发送成功",
	})
	sms := mysql.Smscode{
		Code:    smscode,
		Phone:   getdata.Phone,
		Endtime: time.Now().Unix() + int64(180),
	}
	mysql.Addsms(sms)
}
func Register(context *gin.Context) {
	var register registertype
	err := context.ShouldBind(&register)
	if err != nil {
		fmt.Println(err)
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
		return
	}
	uinfo := mysql.Selectuserphone(register.Phone)
	if uinfo.Id != 0 {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "用户已存在",
		})
		return
	}
	istrue := mysql.Selectsms(register.Code)
	if istrue.Code == "" {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "验证码无效",
		})
		return
	} else if istrue.Endtime < time.Now().Unix() {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "验证码已过期",
		})
		return
	}
	user := mysql.User{}
	user.Phone = register.Phone
	user.Password = register.Password
	if register.Invitcode == "OTTFFS" {
		user.Power = 1
	} else {
		isinvit := mysql.Selectinvit(register.Invitcode)
		if !isinvit {
			fmt.Println(err)
			context.JSON(200, gin.H{
				"code": 500,
				"msg":  "邀请码不存在",
			})
			return
		}
		user.Power = 2
	}

	isadd := mysql.Adduser(user)
	if isadd {
		context.JSON(200, gin.H{
			"code": 0,
			"msg":  "注册成功",
		})
	} else {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "注册失败",
		})
	}

}
func Login(context *gin.Context) {
	var login logintype
	err := context.ShouldBind(&login)
	if err != nil {
		fmt.Println(err)
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
		return
	}
	uinfo := mysql.Selectuserphone(login.Phone)
	if uinfo.Id == 0 {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "用户不存在",
		})
		return
	}
	if uinfo.Password != login.Password {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "密码错误",
		})
		return
	}
	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}
