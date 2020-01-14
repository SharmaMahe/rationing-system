package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // Indicate RationController.List method to handle GET requests.
	beego.Router("/inventory", &controllers.RationController{}, "get:List")

	// Indicate RationController.Add method to handle GET requests.
	beego.Router("/ration/add", &controllers.RationController{}, "get:Add")

	// Indicate RationController.Post method to handle POST requests.
	beego.Router("/ration/save", &controllers.RationController{}, "post:Post")

	// Indicate RationController.Delete method to handle GET requests.
	beego.Router("/ration/delete/:id", &controllers.RationController{}, "get:Delete")
}
