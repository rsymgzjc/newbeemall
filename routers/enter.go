package routers

import (
	"newbeemall/routers/manager"
	"newbeemall/routers/users"
)

type RouterGroup struct {
	Mall   users.MallRouterGroup
	Manage manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
