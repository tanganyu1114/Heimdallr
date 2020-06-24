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
	HttpPorts     []int    `json:"httpPorts, omitempty"`
	HttpSvrNames  []string `json:"httpSvrNames, omitempty"`
	HttpSvrsNum   int      `json:"httpSvrsNum, omitempty"`
	StreamPorts   []int    `json:"streamPorts, omitempty"`
	StreamSvrsNum int      `json:"streamSvrsNum, omitempty"`
}

func GetDetail(m *Env) *string {
	var n NgConfResult
	url := fmt.Sprintf("http://%s:%d/%s", m.Ipaddr, m.Port, m.RelationPath)

	req := httplib.Get(url).SetTimeout(2*time.Second, 3*time.Second)
	req.Param("token", m.Token)
	req.Param("type", "string")

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
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
