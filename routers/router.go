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
	coinIds := beego.NewNamespace("/coinIds",
		beego.NSRouter("", &controllers.CoinIdController{}, "get:GetAll"),
	)
	rates := beego.NewNamespace("/rates",
		beego.NSRouter("", &controllers.RateController{}, "get:GetAll"),
	)

	beego.AddNamespace(users, assets, coinIds, rates)
}
