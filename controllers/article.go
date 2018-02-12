package controllers

import (
	"github.com/astaxie/beego"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
)

// ArticleController operations for Article
type ArticleController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArticleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Article
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 201 {object} models.Article
// @Failure 403 body is empty
// @router / [post]
func (c *ArticleController) Post() {
	var article models.Article
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &article)
	user, err := models.GetUserById(article.UserId)
	article.UserName = user.Name

	uid, err := models.AddArticle(&article)
	if err != nil {
		c.Ctx.Output.Status = http.StatusBadRequest
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(uid)
		c.Data["json"] = map[string]string{"aid": strconv.Itoa((int)(uid))}
	}
	c.ServeJSON()}

// GetOne ...
// @Title GetOne
// @Description get Article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArticleController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Article
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Article
// @Failure 403
// @router / [get]
func (c *ArticleController) GetAll() {
	articles, err := models.GetAllArticle()
	if err != nil {
		c.Ctx.Output.Status = http.StatusBadRequest
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = articles
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Article
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 200 {object} models.Article
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArticleController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Article
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArticleController) Delete() {

}
