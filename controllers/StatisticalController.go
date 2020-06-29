package controllers

import (
	"Heimdallr/enums"
	"Heimdallr/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"time"
)

//StatisticalController 统计信息
type StatisticalController struct {
	BaseController
}

type BifrostStatus struct {
	Enable int
	Env    string
	*models.StatusMsg
}

//Prepare 参考beego官方文档说明
func (c *StatisticalController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	c.checkLogin()
}

// index page infomation
func (c *StatisticalController) SysInfo() {
	cpupct, err := cpu.Percent(time.Second, false)
	if err != nil {
		logs.Error(err)
	}
	mem, err := mem.VirtualMemory()
	if err != nil {
		logs.Error(err)
	}
	diskinfo, err := disk.Usage("/")
	if err != nil {
		logs.Error(err)
	}
	platform, _, version, err := host.PlatformInformation()
	if err != nil {
		logs.Error(err)
	}
	curTime := time.Now().Format("2006/01/02 15:04:05")
	info := map[string]interface{}{
		"Cpu":     fmt.Sprintf("%.2f", cpupct[0]),
		"Mem":     fmt.Sprintf("%.2f", mem.UsedPercent),
		"Disk":    fmt.Sprintf("%.2f", diskinfo.UsedPercent),
		"Timer":   curTime,
		"Osinfo":  platform + "  " + version,
		"Version": enums.Heimdallr_Version,
	}
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *StatisticalController) Index() {

	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["headcssjs"] = "course/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "statistical/index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("CourseController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("CourseController", "Delete")
}

func (c *StatisticalController) Bifrost() {
	var p models.EnvQueryParam

	data, _ := models.EnvPageList(&p)

	res := make([]*BifrostStatus, 0)
	for i := 0; i < len(data); i++ {
		res = append(res, &BifrostStatus{
			Enable: data[i].Status,
			Env:    data[i].EnvName,
			StatusMsg: func() *models.StatusMsg {
				if data[i].Status == 1 {
					s, err := models.GetStatus(data[i])
					if err != nil {
						logs.Error(err)
						s.Status = "TimeOut"
					}
					return s
				} else {
					return nil
				}
			}(),
		})
	}

	c.Data["json"] = res
	c.ServeJSON()
}
