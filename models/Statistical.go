package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type StatisticsResult struct {
	Message StatisticsMsg `json:"message"`
	Status  string        `json:"status"`
}

type StatisticsMsg struct {
	HttpPorts     []int    `json:"httpPorts, omitempty"`
	HttpSvrNames  []string `json:"httpSvrNames, omitempty"`
	HttpSvrsNum   int      `json:"httpSvrsNum, omitempty"`
	LocNum        int      `json:"locNum, omitempty"`
	StreamPorts   []int    `json:"streamPorts, omitempty"`
	StreamSvrsNum int      `json:"streamSvrsNum, omitempty"`
}

func GetStatistics(m *Env) *StatisticsMsg {
	var s StatisticsResult
	url := fmt.Sprintf("http://%s:%d/%s/statistics", m.Ipaddr, m.Port, m.RelationPath)

	req := httplib.Get(url)
	req.Param("token", m.Token)

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
	}
	json.Unmarshal(b, &s)
	fmt.Println(s)
	return &s.Message
}
