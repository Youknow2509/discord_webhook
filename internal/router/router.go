package router

// router group
type RouterGroup struct {
	SendMessage SendMessageRouterGroup
}

// var
var RouterGroupApp = new(RouterGroup)