package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"time"
)

type NgConfResult struct {
	Message string
	Status  string
}

type StatisticsResult struct {
	Message StatisticsMsg `json:"message"`
	Status  string        `json:"status"`
}

type StatisticsMsg struct {
	HttpPorts     []int            `json:"httpPorts, omitempty"`
	HttpSvrs      map[string][]int `json:"httpSvrs, omitempty"`
	HttpSvrsNum   int              `json:"httpSvrsNum, omitempty"`
	StreamPorts   []int            `json:"streamPorts, omitempty"`
	StreamSvrsNum int              `json:"streamSvrsNum, omitempty"`
}

func GetDetail(m *Env, t string) *string {
	var n NgConfResult
	url := fmt.Sprintf("http://%s:%d/%s", m.Ipaddr, m.Port, m.RelationPath)

	req := httplib.Get(url).SetTimeout(2*time.Second, 3*time.Second)
	req.Param("token", m.Token)
	req.Param("type", t)

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
		n.Message = "获取配置信息详情失败,错误信息请查看日志文件"
	}

	json.Unmarshal(b, &n)
	return &n.Message
}

func GetStatistics(m *Env) *StatisticsMsg {
	var s StatisticsResult
	url := fmt.Sprintf("http://%s:%d/%s/statistics", m.Ipaddr, m.Port, m.RelationPath)

	req := httplib.Get(url)
	req.SetTimeout(2*time.Second, 3*time.Second)
	req.Param("token", m.Token)

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
	}
	json.Unmarshal(b, &s)
	return &s.Message
}
