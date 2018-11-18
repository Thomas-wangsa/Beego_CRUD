package main

import (
	"Beego_CRUD/credentials"
	"Beego_CRUD/database/faker"
	_ "Beego_CRUD/routers"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func main() {
	credentials.Init_Credentials()
	faker.InitData()
	beego.Run()
}
