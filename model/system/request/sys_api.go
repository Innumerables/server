package request

import (
	"server/model/common/request"
	"server/model/system"
)

type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
