package routers

import (
	"Heimdallr/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")
	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	//beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/datareset", &controllers.HomeController{}, "Post:DataReset")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	// 系统信息页面
	beego.Router("/", &controllers.StatisticalController{}, "*:Index")
	beego.Router("/statistical/index", &controllers.StatisticalController{}, "*:Index")
	beego.Router("/statistical/bifrost", &controllers.StatisticalController{}, "Get:Bifrost")
	beego.Router("/Statistical/sysinfo", &controllers.StatisticalController{}, "Get:SysInfo")

	// 环境env
	beego.Router("/env/index", &controllers.EnvController{}, "*:Index")
	beego.Router("/env/datagrid", &controllers.EnvController{}, "Get,Post:DataGrid")
	beego.Router("/env/edit/?:id", &controllers.EnvController{}, "Get,Post:Edit")
	beego.Router("/env/delete", &controllers.EnvController{}, "Post:Delete")
	//beego.Router("/env/aside", &controllers.EnvController{}, "Get,Post:Aside")
	beego.Router("/env/list", &controllers.EnvController{}, "Get,Post:Select2")

	// nginx 信息详情页面
	beego.Router("/detail/index", &controllers.DetailController{}, "*:Index")
	beego.Router("/detail/httpserver", &controllers.DetailController{}, "Get:HttpServer")
	beego.Router("/detail/httpport", &controllers.DetailController{}, "Get:HttpPort")
	beego.Router("/detail/streamport", &controllers.DetailController{}, "Get:StreamPort")

	// nginx 配置
	beego.Router("/settings/global/index", &controllers.NgglobalController{}, "*:Index")
	beego.Router("/settings/global/datagrid", &controllers.NgglobalController{}, "Get,Post:DataGrid")
	beego.Router("/settings/global/edit/?:id", &controllers.NgglobalController{}, "Get,Post:Edit")
	beego.Router("/settings/global/delete", &controllers.NgglobalController{}, "Post:Delete")

}
