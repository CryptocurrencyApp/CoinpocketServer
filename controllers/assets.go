package controllers

import (
	"github.com/astaxie/beego"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"encoding/json"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

// AssetsController operations for Assets
type AssetsController struct {
	beego.Controller
}

type Request struct {
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
	// ここ共通化出来る
	var asset models.Asset
	request := Request{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	asset.CoinId = request.Id
	asset.UserId = request.UserId
	asset.Amount = fmt.Sprint(request.Amount)
	// ここまで

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
		c.Ctx.Output.Status = http.StatusNotFound
		c.Data["json"] = map[string]string{"message": err}
	} else {
		var result []map[string]string

		rates := getRates()

		for _, a := range *asset {
			result = append(result, map[string]string{
				"id":        a.CoinId,
				"amount":    a.Amount,
				"price_jpy": getJpyPrice(a.CoinId, rates),
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
	// ここ共通化出来る
	var asset models.Asset

	request := Request{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	asset.CoinId = request.Id
	asset.UserId = request.UserId
	asset.Amount = fmt.Sprint(request.Amount)
	// ここまで

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
	// ここ共通化出来る
	var asset models.Asset

	request := Request{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	asset.CoinId = request.Id
	asset.UserId = request.UserId
	asset.Amount = fmt.Sprint(request.Amount)
	// ここまで

	err := models.DeleteAsset(&asset)
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.Data["json"] = map[string]string{"message": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"id": asset.CoinId}
	}

	c.ServeJSON()
}

const priceFilePath = "./rateLog/newest.json"

type Rates struct {
	GetAt    string
	InfoList []Rate
}

type Rate struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	PriceUsd         float64 `json:"price_usd"`
	PriceJpy         float64 `json:"price_jpy"`
	PriceBtc         float64 `json:"price_btc"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}

func getJpyPrice(coinId string, rates Rates) (price string) {
	for _, rate := range rates.InfoList {
		if rate.ID == coinId {
			price = fmt.Sprint(rate.PriceJpy)
		}
	}

	return price
}

func getRates() (rates Rates) {
	file, err := os.OpenFile(priceFilePath, os.O_RDONLY, 700)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(raw, &rates)
	if err != nil {
		log.Fatal(err)
	}

	return rates
}
