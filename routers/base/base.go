// Package routers implemented controller methods of beego.
package base

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// baseRouter implemented global settings for all other routers.
type BaseRouter struct {
	beego.Controller
	i18n.Locale
	//	User orm.Users
	IsLogin bool
}
