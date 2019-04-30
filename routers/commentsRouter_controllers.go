package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:PlayController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:PlayController"],
        beego.ControllerComments{
            Method: "AddPlay",
            Router: `/addplay`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:PlayController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:PlayController"],
        beego.ControllerComments{
            Method: "GetPlayInfo",
            Router: `/getplayinfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "AppCheck",
            Router: `/appcheck/:mobile/:code`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"] = append(beego.GlobalControllerRouter["yukee/webapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "PressCode",
            Router: `/presscode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
