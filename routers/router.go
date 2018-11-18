// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Beego_CRUD/controllers"

	"github.com/astaxie/beego"
)

/*
|--------------------------------------------------------------------------
| Function init()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: void
| @init 				: set Routing Beego Framework
|
*/

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/schedule",
			beego.NSInclude(
				&controllers.ScheduleController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
