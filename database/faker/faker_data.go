package faker

import (
	"Beego_CRUD/models"
	"Beego_CRUD/util"
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
)

/*
|--------------------------------------------------------------------------
| Function main()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: void
| @init data			: populate data faker
| @desc 				: iteration/count faker based on app.conf
| for more information  : https://github.com/icrowley/fake
|
*/

func InitData() {
	var err error
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	define_faker, err := beego.AppConfig.Int("TotalFaker")
	if err != nil {
		fmt.Println("GET Total Faker Data Error")
		panic(err)
	}
	data_schedule := PopulateData(define_faker)

	successNums, err := o.InsertMulti(100, data_schedule)
	fmt.Printf("ID: %d, ERR: %v\n", successNums, err)
}

func PopulateData(n int) []models.Schedule {
	total_data := []models.Schedule{}

	for i := 0; i < n; i++ {
		var u models.Schedule
		u.Title = fake.Title()
		u.Description = fake.Street()
		u.Location = fake.StreetAddress()
		u.Attendes = fake.Day()
		u.Schedule_at = time.Now().AddDate(rand.Intn(10), rand.Intn(10), rand.Intn(10))
		u.Created_at = time.Now()
		u.Uuid = util.RandStringBytes(20) + "-" + util.RandStringBytes(rand.Intn(15)) + "-" + util.RandStringBytes(rand.Intn(15))
		total_data = append(total_data, u)
	}
	return total_data
}
