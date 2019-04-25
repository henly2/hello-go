package controller

import (
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject2/common"
	"github.com/henly2/hello-go/myproject2/models"
	"log"
	"net/http"
	"time"
)

type AppClaims struct {
	UserId int `json:"uid"`
	jwt.StandardClaims
}

//用户注册
func UserRegist(c *gin.Context) {
	var (
		response Response

		request models.UserInfo
	)

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("BindJSON err=", err.Error())

		response.Err = common.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}

	//把注册信息加入到数据库
	o := orm.NewOrm()
	user := models.Userorm{}
	user.Username = request.Username
	user.Password = request.Password
	user.Sex = request.Sex
	user.Age = request.Age

	userId, err := o.Insert(&user)
	if err != nil {
		log.Println("Insert err=", err.Error())

		response.Err = common.ErrCode_InternalErr
		response.ErrMsg = "注册失败"
		c.JSON(http.StatusOK, response)
		return
	}

	log.Println("注册成功，userid=", userId)

	// 将新生成的userid返回给用户
	response.Result = userId
	c.JSON(http.StatusOK, response)
}

//登录
func UserLogin(c *gin.Context) {
	var (
		response Response

		request models.User
	)

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("BindJSON err=", err.Error())

		response.Err = common.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}

	//从数据库查询数据，
	//查询所有的数据
	o := orm.NewOrm()
	user := models.Userorm{}
	user.Username = request.Name
	err = o.Read(&user, "username")
	if err != nil {
		log.Println("Read err=", err.Error())

		response.Err = common.ErrCode_UserNotExist
		response.ErrMsg = "用户不存在"
		c.JSON(http.StatusOK, response)
		return
	}

	if user.Password != request.Password {
		response.Err = common.ErrCode_PasswordErr
		response.ErrMsg = "密码错误"
		c.JSON(http.StatusOK, response)
		return
	}

	claims := AppClaims{}
	claims.UserId = user.Id
	claims.ExpiresAt = time.Now().Add(time.Hour * 48).Unix() //设置过期时间，过期需要重新获取
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.SecretKey)) //使用自定义字符串进行加密
	if err != nil {
		log.Println("SignedString err=", err.Error())

		response.Err = common.ErrCode_InternalErr
		response.ErrMsg = "服务内部错误"
		c.JSON(http.StatusOK, response)
		return
	}

	// 设置token
	c.Header("token", tokenString)

	// 将用户数据返回给用户，但是需要将敏感信息过滤
	user.Password = "******"
	response.Result = user
	c.JSON(http.StatusOK, "登录成功")
	c.JSON(http.StatusOK, response)
}

//修改密码
func UserChangePassword(c *gin.Context) {
	var (
		response Response

		request models.ChangePwd
	)

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("BindJSON err=", err.Error())

		response.Err = common.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}

	res_info := c.GetHeader("token")
	token, err := jwt.Parse(res_info, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.SecretKey), nil
	})
	if err != nil {
		log.Println("Parse err=", err.Error())

		response.Err = common.ErrCode_IllegalErr
		response.ErrMsg = "非法数据"
		c.JSON(http.StatusOK, response)
		return
	}

	claims, ok := token.Claims.(*AppClaims)
	if !ok {
		log.Println("not *jwt.MapClaims")

		response.Err = common.ErrCode_IllegalErr
		response.ErrMsg = "非法数据"
		c.JSON(http.StatusOK, response)
		return
	}

	userId := claims.UserId
	if userId == 0 {
		log.Println("userid == 0")
		response.Err = common.ErrCode_IllegalErr
		response.ErrMsg = "非法数据"
		c.JSON(http.StatusOK, response)
		return
	}

	log.Println("token-userId=", userId)

	o := orm.NewOrm()
	user := models.Userorm{}
	user.Id = userId
	user.Username = request.Name
	err = o.Read(&user, "username")
	if err != nil {
		log.Println("Read err=", err.Error())

		response.Err = common.ErrCode_UserNotExist
		response.ErrMsg = "用户不存在"
		c.JSON(http.StatusOK, response)
		return
	}

	if user.Password != request.OldPassword {
		response.Err = common.ErrCode_PasswordErr
		response.ErrMsg = "密码错误"
		c.JSON(http.StatusOK, response)
		return
	}

	user.Password = request.NewPassword
	_, err = o.Update(&user)
	if err != nil {
		log.Println("Update err=", err.Error())

		response.Err = common.ErrCode_InternalErr
		response.ErrMsg = "更新失败"
		c.JSON(http.StatusOK, response)
		return
	}

	// 将用户数据返回给用户，但是需要将敏感信息过滤
	user.Password = "******"
	response.Result = user
	c.JSON(http.StatusOK, response)
}
