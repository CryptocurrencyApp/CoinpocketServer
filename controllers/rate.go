package controllers

import (
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"github.com/astaxie/beego"
	"net/http"
)

// RateController operations for Rate
type RateController struct {
	beego.Controller
}

// GetAll ...
// @Title GetAll
// @Description get Rate
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Rate
// @Failure 403
// @router / [get]
func (c *RateController) GetAll() {
	rates, err := models.GetCoinRates()
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.Data["json"] = map[string]string{"message": "Sorry. We can't send response."}
	} else {
		c.Data["json"] = rates
	}
	c.ServeJSON()
}
