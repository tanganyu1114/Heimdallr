package models

import "github.com/astaxie/beego/orm"

// TableName 设置Env表名
func (a *Env) TableName() string {
	return EnvTBName()
}

// RoleQueryParam 用于搜索的类
type EnvQueryParam struct {
	BaseQueryParam
	EnvLike string
}

type Env struct {
	Id           int
	Seq          int
	EnvName      string
	EnvDescript  string
	Ipaddr       string
	Port         int
	Token        string
	RelationPath string
	Status       int
}

type EnvBtn struct {
	Id          int
	Seq         int
	EnvName     string
	EnvDescript string
	Status      int
}

// RolePageList 获取分页数据
func EnvPageList(params *EnvQueryParam) ([]*Env, int64) {
	query := orm.NewOrm().QueryTable(EnvTBName())
	data := make([]*Env, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("envname__istartswith", params.EnvLike)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// EnvOne 根据id获取单条
/*func EnvOne(id int) (*Env, error) {
	o := orm.NewOrm()
	m := Env{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}*/

func EnvBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(EnvTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
