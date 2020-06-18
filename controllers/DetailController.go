package controllers

import "Heimdallr/models"

type DetailController struct {
	BaseController
}

func (c *DetailController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	// 获取nginx统计信息
	m, b := c.EnvCookie()
	if b {
		c.Data["Content"] = *models.GetDetail(m)
	}

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "detail/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "detail/index_footerjs.html"

}
