package controller_a

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henly2/hello-go/myproject3/common_a"
	"github.com/henly2/hello-go/myproject3/models_a"
	log4 "github.com/jeanphorn/log4go"
	"log"
	"net/http"
	"time"
)

type AppClaims struct {
	UserId int `json:"uid"`
	jwt.StandardClaims
}

var (
	response     Response
	adminrequest models_a.AdminInfo
	request      models_a.UserInfo
)

//普通用户注册
func UserRegist(c *gin.Context) {
	timeUnix := time.Now().Unix()                                         //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05") //获取当前时间

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log4.LOGGER("Test").Info("解析数据失败，%s", err.Error())
		response.Err = common_a.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}
	//把注册信息加入到数据库
	o := orm.NewOrm()
	user := models_a.Userorm{}
	user.Username = request.Username
	user.Password = request.Password
	//将密码hash加密
	data := []byte(user.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	user.Password = md5str1
	user.Sex = request.Sex
	user.Age = request.Age
	user.Registtime = formatTimeStr
	// 查询用户
	userId, err := o.Insert(&user)
	if err != nil {
		log4.LOGGER("Test").Info("失败,错误是%s", err.Error())
		response.Err = common_a.ErrCode_InternalErr
		response.ErrMsg = "注册失败"
		c.JSON(http.StatusOK, response)
		return
	}
	log.Println("注册成功，userid=", userId)

	// 将新生成的userid返回给用户
	response.Result = userId
	c.JSON(http.StatusOK, response)

}

//管理员注册
func AdminRegist(c *gin.Context) {
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")

	// 解析数据
	err := c.BindJSON(&adminrequest)
	if err != nil {
		log4.LOGGER("Test").Info("解析数据失败，%s", err.Error())
		response.Err = common_a.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}

	//明文密码进行加密
	data := []byte(adminrequest.Adminpassword)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	users := &models_a.Admingrom{Root: "adminroot", Username: adminrequest.Adminname, Password: md5str1, Sex: adminrequest.Adminsex, Age: adminrequest.Adminage, Registtime: formatTimeStr}
	models_a.Db.Create(users)

	log.Println("注册成功，userid=", users.Id)
	log4.LOGGER("Test").Info("%s注册成功", users.Username)

	// 将新生成的users返回给用户
	response.Result = users
	c.JSON(http.StatusOK, response)
}

//普通用户登录
func UserLogin(c *gin.Context) {
	var (
		response Response

		request models_a.User
	)

	timeUnix := time.Now().Unix()                                         //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05") //当前登录时间

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("BindJSON err=", err.Error())
		response.Err = common_a.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}

	//从数据库查询数据，
	//查询所有的数据
	o := orm.NewOrm()
	user := models_a.Userorm{}
	user.Username = request.Name
	err = o.Read(&user, "username")
	if err != nil {
		log.Println("Read err=", err.Error())

		response.Err = common_a.ErrCode_UserNotExist
		response.ErrMsg = "用户不存在"
		c.JSON(http.StatusOK, response)
		return
	}
	data := []byte(request.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	request.Password = md5str1
	fmt.Println(request)

	if user.Password != request.Password {
		response.Err = common_a.ErrCode_PasswordErr
		response.ErrMsg = "密码错误"
		c.JSON(http.StatusOK, response)
		return
	}

	claims := AppClaims{}
	claims.UserId = user.Id
	claims.ExpiresAt = time.Now().Add(time.Hour * 48).Unix() //设置过期时间，过期需要重新获取
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common_a.SecretKey)) //使用自定义字符串进行加密
	if err != nil {
		log.Println("SignedString err=", err.Error())

		response.Err = common_a.ErrCode_InternalErr
		response.ErrMsg = "服务内部错误"
		c.JSON(http.StatusOK, response)
		return
	}

	log4.LoadConfiguration("./logininfo.json")
	log4.LOGGER("Test").Info("%s Login successful", user.Username)

	// 设置token
	c.Header("token", tokenString)
	// 将用户数据返回给用户，但是需要将敏感信息过滤
	user.Password = "******"
	response.Result = user
	c.JSON(http.StatusOK, response)

	//把注册信息加入到数据库
	users := &models_a.Userlog{Username: user.Username, Userlogintime: formatTimeStr}
	models_a.Db.Create(users)
	//models_a.Db.Create(models_a.Userlog{Username:user.Username,Userlogintime:formatTimeStr})

	log.Println("保存成功，userid=%", user.Id)
	//// 将新生成的userid返回给用户
	response.Result = users
	c.JSON(http.StatusOK, response)

}

//管理员登录
func AdminLogin(c *gin.Context) {
	var (
		response Response

		request models_a.Admin
	)

	timeUnix := time.Now().Unix()                                         //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05") //当前登录时间

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("BindJSON err=", err.Error())
		response.Err = common_a.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}

	user := models_a.Admingrom{}
	err = models_a.DB.Model(models_a.Admingrom{}).Where("username = ?", request.Name).Find(&user).Error
	if err != nil {
		log.Println("Read err=", err.Error())

		response.Err = common_a.ErrCode_UserNotExist
		response.ErrMsg = "用户不存在"
		c.JSON(http.StatusOK, response)
		return
	}
	//if models_a.DB.Find(&user) == nil {
	//	log.Println("Read err=", err.Error())
	//
	//	response.Err = common_a.ErrCode_UserNotExist
	//	response.ErrMsg = "用户不存在"
	//	c.JSON(http.StatusOK, response)
	//	return
	//}
	user.Username = request.Name
	data := []byte(request.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	request.Password = md5str1

	if user.Password != request.Password {
		response.Err = common_a.ErrCode_PasswordErr
		response.ErrMsg = "密码错误"
		c.JSON(http.StatusOK, response)
		return
	}

	claims := AppClaims{}
	claims.UserId = user.Id
	claims.ExpiresAt = time.Now().Add(time.Hour * 48).Unix() //设置过期时间，过期需要重新获取
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common_a.SecretKey)) //使用自定义字符串进行加密
	if err != nil {
		log.Println("SignedString err=", err.Error())

		response.Err = common_a.ErrCode_InternalErr
		response.ErrMsg = "服务内部错误"
		c.JSON(http.StatusOK, response)
		return
	}

	log4.LoadConfiguration("./logininfo.json")
	log4.LOGGER("Test").Info("%s Login successful", user.Username)

	// 设置token
	c.Header("token", tokenString)
	// 将用户数据返回给用户，但是需要将敏感信息过滤
	user.Password = "******"
	response.Result = user
	c.JSON(http.StatusOK, response)

	//把注册信息加入到数据库
	users := &models_a.Userlog{Username: user.Username, Userlogintime: formatTimeStr}
	models_a.Db.Create(users)
	//models_a.Db.Create(models_a.Userlog{Username:user.Username,Userlogintime:formatTimeStr})

	log.Println("保存成功，userid=%", user.Id)
	//// 将新生成的userid返回给用户
	response.Result = users
	c.JSON(http.StatusOK, response)
}

//查询普通用户登录状态
func QueryAdmin(c *gin.Context) {
	var (
		response Response
		request  models_a.InforMation
	)

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("BindJSON err=", err.Error())

		response.Err = common_a.ErrCode_DataErr
		response.ErrMsg = "错误数据格式不对"
		c.JSON(http.StatusOK, response)
		return
	}
	//token解密
	res_info := c.GetHeader("token")
	token, err := jwt.ParseWithClaims(res_info, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common_a.SecretKey), nil
	})
	if err != nil {
		log.Println("Parse err=", err.Error())

		response.Err = common_a.ErrCode_IllegalErr
		response.ErrMsg = "非法数据"
		c.JSON(http.StatusOK, response)
		return
	}

	claims, ok := token.Claims.(*AppClaims)
	if !ok {
		log.Println("not *jwt.MapClaims")

		response.Err = common_a.ErrCode_IllegalErr
		response.ErrMsg = "非法数据"
		c.JSON(http.StatusOK, response)
		return
	}

	userId := claims.UserId
	if userId == 0 {
		log.Println("userid == 0")
		response.Err = common_a.ErrCode_IllegalErr
		response.ErrMsg = "非法数据"
		c.JSON(http.StatusOK, response)
		return
	}

	log.Println("token-userId=", userId)

	o := orm.NewOrm()
	user := models_a.Userorm{}
	user.Id = userId
	user.Username = request.Name
	err = o.Read(&user, "username")
	if err != nil {
		log.Println("Read err=", err.Error())

		response.Err = common_a.ErrCode_UserNotExist
		response.ErrMsg = "用户不存在"
		c.JSON(http.StatusOK, response)
		return
	}

	// 将用户数据返回给用户，但是需要将敏感信息过滤
	user.Password = "******"
	response.Result = user
	c.JSON(http.StatusOK, response)
}

//查询用户登录日志
func LogAdminLogin(c *gin.Context) {
	var (
		response Response
		request  models_a.Loginlog
	)

	// 解析数据
	err := c.BindJSON(&request)
	if err != nil {
		return
	}
	//从数据库查询数据，
	var userlog []models_a.Userlog

	models_a.Db.Model(&models_a.Userlog{}).Where("username=?", request.Name).Find(&userlog)

	for _, a := range userlog {
		if a.Username == request.Name {
			c.JSON(http.StatusOK, "查询成功")
		}
	}

	response.Result = userlog
	c.JSON(http.StatusOK, response)
}
