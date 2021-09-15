package model

import "github.com/cellargalaxy/go_common/util"

type Template struct {
	TemplateId string `json:"template_id"`
	Title      string `json:"title"`
}

func (object Template) String() string {
	return util.ToJsonString(object)
}

type TemplateData struct {
	Value interface{} `json:"value"`
}

func (object TemplateData) String() string {
	return util.ToJsonString(object)
}

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (object Tag) String() string {
	return util.ToJsonString(object)
}

type UserInfo struct {
	OpenId    string `json:"openid"`
	Nickname  string `json:"nickname"`
	TagIdList []int  `json:"tagid_list"`
}

func (object UserInfo) String() string {
	return util.ToJsonString(object)
}
