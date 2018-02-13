package controllers

import (
	"encoding/json"
	"github.com/CryptocurrencyApp/CoinpocketServer/lib/rate"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"github.com/astaxie/beego"
	"net/http"
)

// AssetsController operations for Assets
type AssetsController struct {
	beego.Controller
}

// URLMapping ...
func (c *AssetsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
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
	json.Unmarshal(c.Ctx.Input.RequestBody, &asset)

	_, err := models.AddAsset(&asset)
	if err != nil {
		c.Ctx.Output.Status = http.StatusBadRequest
		c.Data["json"] = map[string]string{"message": "Already registered. Please use HTTP/PUT method."}
	} else {
		c.Data["json"] = map[string]string{"id": asset.CoinId, "amount": asset.Amount}
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
	userId := c.GetString(":id")

	if userId == "" {
		c.Data["json"] = map[string]string{"message": "You need specified user id"}
		c.ServeJSON()
	}

	asset, err := models.GetAssetById(userId)
	if err != "" {
		//c.Ctx.Output.Status = http.StatusNotFound
		//c.Data["json"] = map[string]string{"message": err}
		// assetsがない場合
		c.Data["json"] = []map[string]string{ }
	} else {
		var result []map[string]string

		rates, _ := rate.GetRates()

		for _, a := range *asset {
			result = append(result, map[string]string{
				"id":        a.CoinId,
				"amount":    a.Amount,
				"price_jpy": rate.GetJpyPrice(a.CoinId, rates),
			})
		}

		c.Data["json"] = result
	}

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
	var asset models.Asset
	json.Unmarshal(c.Ctx.Input.RequestBody, &asset)

	err := models.UpdateAssetById(&asset)
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.Data["json"] = map[string]string{"message": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"id": asset.CoinId, "amount": asset.Amount}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Assets
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AssetsController) Delete() {
	var asset models.Asset
	asset.UserId = c.GetString(":uid")
	asset.CoinId = c.GetString(":cid")
	err := models.DeleteAsset(&asset)
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.Data["json"] = map[string]string{"message": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "Delete Completed"}
	}

	c.ServeJSON()
}