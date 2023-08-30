package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"gofox/controllers"
	_ "gofox/routers"
	"time"
)

const (
	INDEX = "http://127.0.0.1:8082/sys_user/login_form" //后台首页
)

func convertM(in int) (out int64) {
	out = int64(in) / 1000000
	return
}

func convertT(in uint) (out string) {
	tm := time.Unix(int64(in), 0)
	out = tm.Format("2006-01-02 15:04:05")
	return
}

func main() {
	// 定制路径
	customize_path()

	beego.AddFuncMap("convertm", convertM)
	beego.AddFuncMap("convertt", convertT)

	fmt.Printf("%+v\r\n", beego.BConfig.WebConfig)
	fmt.Println(INDEX)

	beego.Run()
}

func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	beego.BConfig.Log.AccessLogs = true
	beego.SetLogger("file", `{"filename":"logs/beego.log"}`)
	//beego.BeeLogger.DelLogger("console")

	//初始化数据库
	init_db()

	beego.ErrorController(&controllers.ErrorController{})
}

func init_db() {
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := beego.AppConfig.String("db_type")
	if len(dbType) <= 0 {
		panic("")
	}
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")

	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
			dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	}

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动建表
	//orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}

}

// 定制路径
func customize_path() {
	//beego.SetStaticPath("/static/", "../static")
	//beego.SetViewsPath("../views")
}
