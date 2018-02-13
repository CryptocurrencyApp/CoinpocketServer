package controllers

import (
	"encoding/json"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"github.com/astaxie/beego"
	"net/http"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

type LoginRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Login() {
	request := LoginRequest{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	user := models.User{
		Mail:     request.Mail,
		Password: request.Password,
	}

	userId, err := models.Login(&user)
	if err != nil {
		c.Ctx.Output.Status = http.StatusForbidden
		c.Data["json"] = map[string]string{"message": "Authenticate failed."}
	} else {
		c.Data["json"] = map[string]string{"id": userId}
	}

	c.ServeJSON()
}
