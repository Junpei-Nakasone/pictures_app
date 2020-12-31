package waf

import (
	api002 "pictures_app/api/api002/infra/web"
)

type Handler struct {
	api002 api002.Handler
}

func NewHandler(api002 api002.Handler) Handler {
	return Handler{api002: api002}
}
