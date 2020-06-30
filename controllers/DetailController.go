package controllers

import "Heimdallr/models"

type DetailController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *DetailController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	c.checkLogin()
}

func (c *DetailController) confInfo(m *models.Env) {

	s := models.GetStatistics(m)
	nginfo := map[string]interface{}{
		"HttpPorts":     s.HttpPorts,
		"HttpPortNum":   len(s.HttpPorts),
		"HttpSvrs":      s.HttpSvrs,
		"HttpSvrsNum":   s.HttpSvrsNum,
		"StreamPorts":   s.StreamPorts,
		"StreamPortNum": len(s.StreamPorts),
		"StreamSvrsNum": s.StreamSvrsNum,
	}

	c.Data["NgInfo"] = nginfo
}

func (c *DetailController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	// 获取nginx统计信息
	m, b := c.EnvCookie()
	if b {
		c.confInfo(m)
		c.Data["Content"] = *models.GetDetail(m)
	}

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "detail/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "detail/index_footerjs.html"

}

func (c *DetailController) HttpServer() {
	m, b := c.EnvCookie()
	if b {
		c.confInfo(m)
	}
	c.setTpl("detail/http_server.html", "shared/layout_pullbox.html")
}

func (c *DetailController) HttpPort() {
	m, b := c.EnvCookie()
	if b {
		c.confInfo(m)
	}
	c.setTpl("detail/http_port.html", "shared/layout_pullbox.html")
}

func (c *DetailController) StreamPort() {
	m, b := c.EnvCookie()
	if b {
		c.confInfo(m)
	}
	c.setTpl("detail/stream_port.html", "shared/layout_pullbox.html")
}
