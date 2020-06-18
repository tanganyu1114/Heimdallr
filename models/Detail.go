package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type NgConfResult struct {
	Message string
	Status  string
}

func GetDetail(m *Env) *string {
	var n NgConfResult
	url := fmt.Sprintf("http://%s:%d/%s", m.Ipaddr, m.Port, m.RelationPath)

	req := httplib.Get(url)
	req.Param("token", m.Token)
	req.Param("type", "string")

	b, err := req.Bytes()
	if err != nil {
		logs.Error(err)
	}
	json.Unmarshal(b, &n)
	return &n.Message

}
