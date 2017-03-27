package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ordering/controllers:MenuController"] = append(beego.GlobalControllerRouter["ordering/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:MenuController"] = append(beego.GlobalControllerRouter["ordering/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:MenuController"] = append(beego.GlobalControllerRouter["ordering/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:MenuController"] = append(beego.GlobalControllerRouter["ordering/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:MenuController"] = append(beego.GlobalControllerRouter["ordering/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:OrderController"] = append(beego.GlobalControllerRouter["ordering/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:OrderController"] = append(beego.GlobalControllerRouter["ordering/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:OrderController"] = append(beego.GlobalControllerRouter["ordering/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:OrderController"] = append(beego.GlobalControllerRouter["ordering/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:OrderController"] = append(beego.GlobalControllerRouter["ordering/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:UserController"] = append(beego.GlobalControllerRouter["ordering/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:UserController"] = append(beego.GlobalControllerRouter["ordering/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:UserController"] = append(beego.GlobalControllerRouter["ordering/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:UserController"] = append(beego.GlobalControllerRouter["ordering/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["ordering/controllers:UserController"] = append(beego.GlobalControllerRouter["ordering/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
