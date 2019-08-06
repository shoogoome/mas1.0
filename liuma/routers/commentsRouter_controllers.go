package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["liuma/controllers:ServerControllers"] = append(beego.GlobalControllerRouter["liuma/controllers:ServerControllers"],
        beego.ControllerComments{
            Method: "SSL",
            Router: `/.well-known/pki-validation/fileauth.txt`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers:ServerControllers"] = append(beego.GlobalControllerRouter["liuma/controllers:ServerControllers"],
        beego.ControllerComments{
            Method: "Active",
            Router: `/server/active`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers:ServerControllers"] = append(beego.GlobalControllerRouter["liuma/controllers:ServerControllers"],
        beego.ControllerComments{
            Method: "Signal",
            Router: `/server/signal`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers:ServerControllers"] = append(beego.GlobalControllerRouter["liuma/controllers:ServerControllers"],
        beego.ControllerComments{
            Method: "ChangeToken",
            Router: `/server/token`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
