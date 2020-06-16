package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type Request struct {
	Ipaddr string
	Port   string
	Token  string
}

type NgConfRes struct {
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
	LocNum        int      `json:"locNum, omitempty"`
	StreamPorts   []int    `json:"streamPorts, omitempty"`
	StreamSvrsNum int      `json:"streamSvrsNum, omitempty"`
}

var Req Request

func GetAllConf() *string {
	var n NgConfRes
	url := fmt.Sprintf("http://%s:%s/ng_conf", Req.Ipaddr, Req.Port)

	req := httplib.Get(url)
	req.Param("token", Req.Token)
	req.Param("type", "string")

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
	}
	json.Unmarshal(b, &n)
	return &n.Message

}

func GetStatistics() *StatisticsMsg {
	var s StatisticsResult
	url := fmt.Sprintf("http://%s:%s/ng_conf/statistics", Req.Ipaddr, Req.Port)

	req := httplib.Get(url)
	req.Param("token", Req.Token)

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
	}
	json.Unmarshal(b, &s)
	fmt.Println(s)
	return &s.Message
}
