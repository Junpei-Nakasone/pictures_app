package waf

import (
	api001 "pictures_app/api/api001/infra/web"
	api002 "pictures_app/api/api002/infra/web"
	api003 "pictures_app/api/api003/infra/web"
	api004 "pictures_app/api/api004/infra/web"
	api005 "pictures_app/api/api005/infra/web"
	api006 "pictures_app/api/api006/infra/web"
	api007 "pictures_app/api/api007/infra/web"
	api008 "pictures_app/api/api008/infra/web"
	api009 "pictures_app/api/api009/infra/web"
	api010 "pictures_app/api/api010/infra/web"
)

type Handler struct {
	api001 api001.Handler
	api002 api002.Handler
	api003 api003.Handler
	api004 api004.Handler
	api005 api005.Handler
	api006 api006.Handler
	api007 api007.Handler
	api008 api008.Handler
	api009 api009.Handler
	api010 api010.Handler
}

func NewHandler(
	api001 api001.Handler,
	api002 api002.Handler,
	api003 api003.Handler,
	api004 api004.Handler,
	api005 api005.Handler,
	api006 api006.Handler,
	api007 api007.Handler,
	api008 api008.Handler,
	api009 api009.Handler,
	api010 api010.Handler,
) Handler {
	return Handler{
		api001: api001,
		api002: api002,
		api003: api003,
		api004: api004,
		api005: api005,
		api006: api006,
		api007: api007,
		api008: api008,
		api009: api009,
		api010: api010,
	}
}
