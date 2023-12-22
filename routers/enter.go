package routers

import "newbeemall/routers/users"

type RouterGroup struct {
	Mall users.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
