package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"time"
)

type StatusMsg struct {
	Status  string  `json:"status"`
	Message SysInfo `json:"message"`
}

type SysInfo struct {
	System string `json:"system"`
	Time   string `json:"time"`
	Cpu    string `json:"cpu"`
	Mem    string `json:"mem"`
	Disk   string `json:"disk"`
	Ver    string `json:"bifrost_version"`
}

func GetStatus(m *Env) (*StatusMsg, error) {
	var s StatusMsg
	url := fmt.Sprintf("http://%s:%d/sysinfo", m.Ipaddr, m.Port)

	req := httplib.Get(url)
	req.SetTimeout(2*time.Second, 3*time.Second)
	req.Param("token", m.Token)

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
	}

	json.Unmarshal(b, &s)
	return &s, err
}
