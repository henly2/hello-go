package routers

import (
	"github.com/astaxie/beego"
	"github.com/henly2/hello-go/myproject/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test_input", &controllers.TestinputController{}, "get:Get;post:Post")
	beego.Router("/test_login", &controllers.TestloginController{}, "get:Login;post:Post")
	//beego.Router("/", &controllers.TestloginController{},"post:Regist")
	beego.Router("/test_model", &controllers.TestmodelController{}, "get:Get")

	//注册
	beego.Router("/regist", &controllers.RegistController{}, "post:Regist")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/changepwd", &controllers.ChangepwdController{}, "post:Changepwd")

}
