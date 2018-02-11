package controllers

import (
	"github.com/astaxie/beego"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"encoding/json"
	"log"
	"strconv"
	"fmt"
)

// AssetsController operations for Assets
type AssetsController struct {
	beego.Controller
}

type PostRequest struct {
	Id     string
	UserId string `json:"user_id"`
	Amount float64
}

// URLMapping ...
func (c *AssetsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Assets
// @Param	body		body 	models.Assets	true		"body for Assets content"
// @Success 201 {object} models.Assets
// @Failure 403 body is empty
// @router / [post]
func (c *AssetsController) Post() {
	var asset models.Asset
	request := PostRequest{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	asset.CoinId = request.Id
	asset.UserId = request.UserId
	asset.Amount = fmt.Sprint(request.Amount)

	id, err := models.AddAsset(&asset)
	if err != nil {
		log.Fatal(err)
		c.Data["json"] = err
	} else {
		c.Data["json"] = map[string]string{"id": strconv.Itoa(int(id))}
	}

	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Assets by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Assets
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AssetsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Assets
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Assets
// @Failure 403
// @router / [get]
func (c *AssetsController) GetAll() {

	c.Data["json"] = map[string]string{"example": "example"}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Assets
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Assets	true		"body for Assets content"
// @Success 200 {object} models.Assets
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AssetsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Assets
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AssetsController) Delete() {

}
