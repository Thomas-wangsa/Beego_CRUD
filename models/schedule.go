package models

import (
	"Beego_CRUD/util"
	"fmt"
	"math/rand"
	"time"

	"errors"

	"github.com/astaxie/beego/orm"
)

type Schedule struct {
	Id          int    `orm:"pk;auto"`
	Title       string `orm:"size(100)"`
	Description string `orm:"size(100)"`
	Attendes    int
	Location    string    `orm:"size(100);type(text)"`
	Uuid        string    `orm:"size(100);type(text)"`
	Schedule_at time.Time `orm:"auto_now;type(datetime)"`
	Created_at  time.Time `orm:"type(datetime)"`
	Updated_at  time.Time `orm:"auto_now;type(datetime);null"`
}

type AddSchedulesStruct struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Attendes      int    `json:"attendes"`
	Location      string `json:"location"`
	Schedule_at   string `json:"schedule_at"`
	Schedule_time time.Time
}

func init() {
	orm.RegisterModel(new(Schedule))
}

var (
	ScheduleList map[string]*Schedule
)

type DataSchedule struct {
	Uuid  string `orm:"size(100);type(text)"`
	Title string `orm:"size(100)"`
	Date  string
	Time  string
}

type ReadOneSchedule struct {
	Title       string `orm:"size(100)"`
	Description string `orm:"size(100)"`
	Attendes    int
	Date        string
	Time        string
	Location    string `orm:"size(100);type(text)"`
}

func GetAllSchedule() []DataSchedule {
	var ReadDataSchedule []DataSchedule
	var o = orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("schedule").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			var each = DataSchedule{}
			each.Uuid = m["Uuid"].(string)
			each.Title = m["Title"].(string)
			each.Date = m["Schedule_at"].(time.Time).Format("2006-01-02")
			each.Time = m["Schedule_at"].(time.Time).Format("15:04:05")
			ReadDataSchedule = append(ReadDataSchedule, each)
		}
	}
	return ReadDataSchedule
}

func GetSchedule(uid string) (ReadOneSchedule, error) {
	o := orm.NewOrm()
	var err error
	var schedule Schedule
	var readoneschedule ReadOneSchedule
	err = o.QueryTable("schedule").Filter("uuid", uid).One(&schedule)

	if err == orm.ErrMultiRows {
		// Have multiple records
		fmt.Printf("Returned Multi Rows Not One")
		return readoneschedule, errors.New("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// No result
		fmt.Printf("Not row found")
		return readoneschedule, errors.New("Not row found")
	}

	readoneschedule.Title = schedule.Title
	readoneschedule.Description = schedule.Description
	readoneschedule.Attendes = schedule.Attendes
	readoneschedule.Date = schedule.Schedule_at.Format("2006-01-02")
	readoneschedule.Time = schedule.Schedule_at.Format("15:04:05")
	readoneschedule.Location = schedule.Location
	return readoneschedule, nil
}

func AddSchedule(addSchedulesStruct AddSchedulesStruct) int64 {

	o := orm.NewOrm()
	var data Schedule
	data.Title = addSchedulesStruct.Title
	data.Description = addSchedulesStruct.Description
	data.Attendes = addSchedulesStruct.Attendes
	data.Location = addSchedulesStruct.Location
	data.Schedule_at = addSchedulesStruct.Schedule_time
	data.Created_at = time.Now()
	data.Uuid = util.RandStringBytes(20) + "-" + util.RandStringBytes(rand.Intn(15)) + "-" + util.RandStringBytes(rand.Intn(15))

	id, err := o.Insert(&data)
	if err != nil {
		panic(err)
	}

	return id
}

type UpdateScheduleStruct struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Attendes      int    `json:"attendes"`
	Location      string `json:"location"`
	Schedule_at   string `json:"schedule_at"`
	Schedule_time time.Time
}

func UpdateSchedule(uid string, addSchedulesStruct *AddSchedulesStruct) (bool, error) {
	o := orm.NewOrm()
	exist := o.QueryTable("schedule").Filter("uuid", uid).Exist()

	if exist {

		num, err := o.QueryTable("schedule").Filter("uuid", uid).Update(orm.Params{
			"Title":       addSchedulesStruct.Title,
			"Description": addSchedulesStruct.Description,
			"Attendes":    addSchedulesStruct.Attendes,
			"Location":    addSchedulesStruct.Location,
			"Schedule_at": addSchedulesStruct.Schedule_time,
		})
		if err != nil {
			return false, err
		}
		fmt.Println(num)
		return true, nil
	} else {
		return false, errors.New("Data not exist")
	}

}

func DeleteSchedule(uid string) error {
	o := orm.NewOrm()
	exist := o.QueryTable("schedule").Filter("uuid", uid).Exist()

	if exist {
		fmt.Println("DATA exist")
		num, err := o.QueryTable("schedule").Filter("uuid", uid).Delete()

		if err != nil {
			fmt.Println("DATA failed")
			return err
		}

		if num < 1 {
			fmt.Println("DATA is not exist")
			return errors.New("Delete Failed!")
		} else {
			fmt.Println("Delete success")
			return nil
		}
	} else {
		fmt.Println("DATA not exist")
		return errors.New("Data not exist")
	}
}
