package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "DeleteChuck",
            Router: `/data/chuck`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "GetChuck",
            Router: `/data/chuck`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "SaveShard",
            Router: `/data/shard`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "DeleteShard",
            Router: `/data/shard`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "GetShard",
            Router: `/data/shard`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/download`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "GenerateDownloadToken",
            Router: `/download/token`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "FileInfo",
            Router: `/file/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "ChunkUpload",
            Router: `/upload/chuck`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "Finish",
            Router: `/upload/finish`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "SingleUpload",
            Router: `/upload/single`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"] = append(beego.GlobalControllerRouter["liuma/controllers/file:FileSystemController"],
        beego.ControllerComments{
            Method: "GenerateUploadToken",
            Router: `/upload/token`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
