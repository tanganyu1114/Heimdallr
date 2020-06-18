package controllers

import (
	"Heimdallr/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"html"
	"net/url"
	"time"
)

//StatisticalController 统计信息
type StatisticalController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *StatisticalController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

// index page infomation
func (c *StatisticalController) sysInfo() {
	cpupct, err := cpu.Percent(time.Second, false)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	mem, err := mem.VirtualMemory()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	diskinfo, err := disk.Usage("/")
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	platform, _, version, err := host.PlatformInformation()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	curTime := time.Now().Format("2006/01/02 15:04:05")
	info := map[string]interface{}{
		"Cpu":    fmt.Sprintf("%.2f", cpupct[0]),
		"Mem":    fmt.Sprintf("%.2f", mem.UsedPercent),
		"Disk":   fmt.Sprintf("%.2f", diskinfo.UsedPercent),
		"Timer":  curTime,
		"Osinfo": platform + "  " + version,
	}
	c.Data["Info"] = info
}

func (c *StatisticalController) confInfo(m *models.Env) {

	s := models.GetStatistics(m)
	nginfo := map[string]interface{}{
		"HttpPorts":     s.HttpPorts,
		"HttpPortNum":   len(s.HttpPorts),
		"HttpSvrNames":  s.HttpSvrNames,
		"HttpSvrsNum":   s.HttpSvrsNum,
		"LocationNum":   s.LocNum,
		"StreamPorts":   s.HttpPorts,
		"StreamPortNum": len(s.StreamPorts),
		"StreamSvrsNum": s.StreamSvrsNum,
	}

	c.Data["NgInfo"] = nginfo
}

func (c *StatisticalController) Index() {

	var m models.Env
	m.EnvName = c.Ctx.GetCookie("env")
	fmt.Println(m.EnvName)
	//url.QueryUnescape(m.EnvName)
	fmt.Println(html.UnescapeString(m.EnvName))
	um, err := url.QueryUnescape(m.EnvName)
	if err != nil {
		logs.Error(err)
	}
	m.EnvName = um

	if m.EnvName != "" {
		err := orm.NewOrm().QueryTable(models.EnvTBName()).Filter("EnvName", m.EnvName).One(&m)
		if err != nil {
			logs.Error(err)
		}
	}
	fmt.Println("env info: \n", m)
	c.sysInfo()
	c.confInfo(&m)

	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["headcssjs"] = "course/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "statistical/index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("CourseController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("CourseController", "Delete")
}

// 可以通过修改底层url.QueryEscape代码获得更高的效率，很简单
/*func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}*/
