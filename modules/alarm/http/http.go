package http

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/Cepave/open-falcon-backend/modules/alarm/g"
	log "github.com/Sirupsen/logrus"
	_ "net/http/pprof"

	"github.com/astaxie/beego"
)

func configRoutes() {
	beego.Router("/", &MainController{}, "get:Index")
	beego.Router("/version", &MainController{}, "get:Version")
	beego.Router("/health", &MainController{}, "get:Health")
	beego.Router("/workdir", &MainController{}, "get:Workdir")
	beego.Router("/config/reload", &MainController{}, "get:ConfigReload")
	beego.Router("/event", &MainController{}, "get:Event")
	beego.Router("/event/solve", &MainController{}, "post:Solve")
}

func Duration(now, before int64) string {
	d := now - before
	if d <= 60 {
		return "just now"
	}

	if d <= 120 {
		return "1 minute ago"
	}

	if d <= 3600 {
		return fmt.Sprintf("%d minutes ago", d/60)
	}

	if d <= 7200 {
		return "1 hour ago"
	}

	if d <= 3600*24 {
		return fmt.Sprintf("%d hours ago", d/3600)
	}

	if d <= 3600*24*2 {
		return "1 day ago"
	}

	return fmt.Sprintf("%d days ago", d/3600/24)
}

func init() {
	configRoutes()
	beego.AddFuncMap("duration", Duration)
}

func Start() {
	if !g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}

	if g.Config().Debug {
		beego.BConfig.RunMode = "dev"
	} else {
		beego.BConfig.RunMode = "prod"
	}

	beego.Run(addr)

	log.Println("http listening", addr)
}
