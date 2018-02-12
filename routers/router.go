// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/CryptocurrencyApp/CoinpocketServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	users := beego.NewNamespace("/users",
		beego.NSRouter("/", &controllers.UserController{}, "post:Post"),
		beego.NSRouter("/:id", &controllers.UserController{}, "get:Get"),
		beego.NSRouter("/:id", &controllers.UserController{}, "put:Put"),
	)
	assets := beego.NewNamespace("/assets",
		beego.NSRouter("", &controllers.AssetsController{}),
		beego.NSRouter("/:id", &controllers.AssetsController{}, "get:GetAll"),
	)
	articles := beego.NewNamespace("/articles",
		beego.NSRouter("/", &controllers.ArticleController{}, "post:Post"),
		beego.NSRouter("/", &controllers.ArticleController{}, "get:GetAll"),
		beego.NSRouter("/:id", &controllers.ArticleController{}, "get:GetOne"),
		beego.NSRouter("/:id/good", &controllers.ArticleController{}, "put:ToggleGood"),
		beego.NSRouter("/:id/bad", &controllers.ArticleController{}, "put:ToggleBad"),
		beego.NSRouter("/:id", &controllers.ArticleController{}, "delete:Delete"),
		beego.NSRouter("/users/:uid/articles", &controllers.ArticleController{}, "get:GetUsersAll"),
	)
	coinIds := beego.NewNamespace("/coinIds",
		beego.NSRouter("", &controllers.CoinIdController{}, "get:GetAll"),
	)
	rates := beego.NewNamespace("/rates",
		beego.NSRouter("", &controllers.RateController{}, "get:GetAll"),
	)

	beego.AddNamespace(users, assets, articles, coinIds, rates)
}
