// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"yukee/webapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
    
    //public
    
    //路由自动化
	ns := beego.NewNamespace("/public",
		//注册功能模块
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/play",
			beego.NSInclude(
                &controllers.PlayController{},
                //可以多个
			),
		),
	)
	beego.AddNamespace(ns)
}
