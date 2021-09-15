package model

import "github.com/cellargalaxy/go_common/util"

type ListAllTemplateRequest struct {
}

func (this ListAllTemplateRequest) String() string {
	return util.ToJsonString(this)
}

type ListAllTemplateResponse struct {
	List []Template `json:"list"`
}

func (this ListAllTemplateResponse) String() string {
	return util.ToJsonString(this)
}

type SendTemplateToTagRequest struct {
	TemplateId string                 `form:"template_id" json:"template_id"`
	TagId      int                    `form:"tag_id" json:"tag_id"`
	Url        string                 `form:"url" json:"url"`
	Data       map[string]interface{} `form:"data" json:"data"`
}

func (this SendTemplateToTagRequest) String() string {
	return util.ToJsonString(this)
}

type SendTemplateToTagResponse struct {
	FailOpenIds []string `json:"fail_open_ids"`
}

func (this SendTemplateToTagResponse) String() string {
	return util.ToJsonString(this)
}

type SendTemplateToCommonTagRequest struct {
	Text string `form:"text" json:"text"`
}

func (this SendTemplateToCommonTagRequest) String() string {
	return util.ToJsonString(this)
}

type SendTemplateToCommonTagResponse struct {
	FailOpenIds []string `json:"fail_open_ids"`
}

func (this SendTemplateToCommonTagResponse) String() string {
	return util.ToJsonString(this)
}

type CreateTagRequest struct {
	Tag string `form:"tag" json:"tag"`
}

func (this CreateTagRequest) String() string {
	return util.ToJsonString(this)
}

type CreateTagResponse struct {
	Result bool `json:"result"`
}

func (this CreateTagResponse) String() string {
	return util.ToJsonString(this)
}

type DeleteTagRequest struct {
	TagId int `form:"tag_id" json:"tag_id"`
}

func (this DeleteTagRequest) String() string {
	return util.ToJsonString(this)
}

type DeleteTagResponse struct {
	Result bool `json:"result"`
}

func (this DeleteTagResponse) String() string {
	return util.ToJsonString(this)
}

type AddTagToUserRequest struct {
	TagId  int    `form:"tag_id" json:"tag_id"`
	OpenId string `form:"open_id" json:"open_id"`
}

func (this AddTagToUserRequest) String() string {
	return util.ToJsonString(this)
}

type AddTagToUserResponse struct {
	Result bool `json:"result"`
}

func (this AddTagToUserResponse) String() string {
	return util.ToJsonString(this)
}

type DeleteTagFromUserRequest struct {
	TagId  int    `form:"tag_id" json:"tag_id"`
	OpenId string `form:"open_id" json:"open_id"`
}

func (this DeleteTagFromUserRequest) String() string {
	return util.ToJsonString(this)
}

type DeleteTagFromUserResponse struct {
	Result bool `json:"result"`
}

func (this DeleteTagFromUserResponse) String() string {
	return util.ToJsonString(this)
}

type ListAllTagRequest struct {
}

func (this ListAllTagRequest) String() string {
	return util.ToJsonString(this)
}

type ListAllTagResponse struct {
	List []Tag `json:"list"`
}

func (this ListAllTagResponse) String() string {
	return util.ToJsonString(this)
}

type ListAllUserInfoRequest struct {
}

func (this ListAllUserInfoRequest) String() string {
	return util.ToJsonString(this)
}

type ListAllUserInfoResponse struct {
	List []UserInfo `json:"list"`
}

func (this ListAllUserInfoResponse) String() string {
	return util.ToJsonString(this)
}

type SendTgMsg2ConfigChatIdRequest struct {
	Text string `form:"text" json:"text"`
}

func (this SendTgMsg2ConfigChatIdRequest) String() string {
	return util.ToJsonString(this)
}

type SendTgMsg2ConfigChatIdResponse struct {
	Result bool `json:"result"`
}

func (this SendTgMsg2ConfigChatIdResponse) String() string {
	return util.ToJsonString(this)
}
