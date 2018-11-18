package credentials

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
|--------------------------------------------------------------------------
| Function main()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: void
| @init credentials 	: set connection database and migration table
| for more information  : https://beego.me/docs/mvc/controller/config.md
|
*/

func Init_Credentials() {
	orm.RegisterDriver(beego.AppConfig.String("ConnectionType"), orm.DRPostgres)
	orm.RegisterDataBase("default",
		beego.AppConfig.String("ConnectionType"),
		"user="+beego.AppConfig.String("Username")+
			" password="+beego.AppConfig.String("Password")+
			" host="+beego.AppConfig.String("Host")+
			" port="+beego.AppConfig.String("Port")+
			" dbname="+beego.AppConfig.String("DBName")+
			" sslmode='"+beego.AppConfig.String("SSLMode")+"'")

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
