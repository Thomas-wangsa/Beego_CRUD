package controllers

import (
	"Beego_CRUD/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

// Operations about Users
type ScheduleController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (sch *ScheduleController) GetAll() {
	schedules := models.GetAllSchedule()
	sch.Data["json"] = schedules
	sch.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Schedule
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *ScheduleController) Get() {
	uid := u.GetString(":uid")
	fmt.Println("UID = " + uid)
	if uid != "" {
		user, err := models.GetSchedule(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ScheduleController) Post() {
	var addschedulestruct models.AddSchedulesStruct
	json.Unmarshal(u.Ctx.Input.RequestBody, &addschedulestruct)

	if addschedulestruct.Title == "" ||
		addschedulestruct.Description == "" ||
		addschedulestruct.Location == "" ||
		addschedulestruct.Attendes == 0 ||
		addschedulestruct.Schedule_at == "" {
		panic("title,description,location,attendes,schedule_at are required")
	}

	layout := "2006-01-02 15:04:05.999999"
	str := addschedulestruct.Schedule_at + ".00000"
	t, err := time.Parse(layout, str)

	if err != nil {
		panic(err)
	}
	addschedulestruct.Schedule_time = t

	uid := models.AddSchedule(addschedulestruct)
	u.Data["json"] = map[string]int64{"id": uid}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *ScheduleController) Put() {
	uid := u.GetString(":uid")
	fmt.Println("ID = " + uid)
	if uid != "" {
		var addschedulestruct models.AddSchedulesStruct
		json.Unmarshal(u.Ctx.Input.RequestBody, &addschedulestruct)

		if addschedulestruct.Title == "" ||
			addschedulestruct.Description == "" ||
			addschedulestruct.Location == "" ||
			addschedulestruct.Attendes == 0 ||
			addschedulestruct.Schedule_at == "" {
			panic("title,description,location,attendes,schedule_at are required")
		}

		layout := "2006-01-02 15:04:05.999999"
		str := addschedulestruct.Schedule_at + ".00000"
		t, err := time.Parse(layout, str)

		if err != nil {
			panic(err)
		}
		addschedulestruct.Schedule_time = t

		uu, err := models.UpdateSchedule(uid, &addschedulestruct)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = map[string]bool{"update": uu}
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *ScheduleController) Delete() {
	uid := u.GetString(":uid")
	var err error
	err = models.DeleteSchedule(uid)
	if err == nil {
		u.Data["json"] = "delete success!"
	} else {
		u.Data["json"] = "delete failed!"
	}

	u.ServeJSON()
}
