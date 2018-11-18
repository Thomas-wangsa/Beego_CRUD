package main

import (
	"Beego_CRUD/credentials"
	"Beego_CRUD/database/faker"
	_ "Beego_CRUD/routers"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

/*
|--------------------------------------------------------------------------
| Function main()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: void
| @init credentials 	: set connection database and migration table
| @init data			: populate data faker
| @Run 					: Execute Beego Frameworks
|
*/

func main() {
	credentials.Init_Credentials()
	faker.InitData()
	beego.Run()
}
