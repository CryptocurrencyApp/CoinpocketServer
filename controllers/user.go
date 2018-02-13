package controllers

import (
	"encoding/json"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"github.com/astaxie/beego"
	"net/http"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

//@Title CreateUser
//@Description create users
//@Param	body		body 	models.User	true		"body for user content"
//@Success 200 {int} models.User.Id
//@Failure 403 body is empty
//@router / [post]
func (u *UserController) Post() {
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid, err := models.AddUser(&user)
	if err != nil {
		u.Ctx.Output.Status = http.StatusBadRequest
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = map[string]string{"uid": uid}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":id")
	if uid != "" {
		user, err := models.GetUserById(uid)
		if err != nil {
			u.Ctx.Output.Status = http.StatusNotFound
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]

func (u *UserController) Put() {
	uid := u.GetString(":id")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		user.Id = uid
		err := models.UpdateUserById(&user)
		if err != nil {
			u.Ctx.Output.Status = http.StatusBadRequest
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]

//func (u *UserController) Login() {
//	username := u.GetString("username")
//	password := u.GetString("password")
//	if models.Login(username, password) {
//		u.Data["json"] = "login success"
//	} else {
//		u.Data["json"] = "user not exist"
//	}
//	u.ServeJSON()
//}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]

//func (u *UserController) Logout() {
//	u.Data["json"] = "logout success"
//	u.ServeJSON()
//}
