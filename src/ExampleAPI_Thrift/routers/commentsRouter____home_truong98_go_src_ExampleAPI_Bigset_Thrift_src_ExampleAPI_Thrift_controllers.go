package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/controllers:TeamController"] = append(beego.GlobalControllerRouter["ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/controllers:TeamController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/controllers:TeamController"] = append(beego.GlobalControllerRouter["ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/controllers:TeamController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/controllers:TeamController"] = append(beego.GlobalControllerRouter["ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/controllers:TeamController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
