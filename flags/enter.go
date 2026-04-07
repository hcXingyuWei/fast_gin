package flags

import (
	"fast_gin/global"
	"flag"
	"fmt"
	"os"
)

type FlagOptions struct {
	File    string
	DB      bool
	Version bool
	Menu    string //菜单
	Type    string //类型
}

var Opts FlagOptions

func Run() (ok bool) {
	if Opts.DB {
		MigrateDB()
		os.Exit(0)
	}
	if Opts.Version {
		fmt.Printf("当前版本%s", global.Version)
		os.Exit(0)
	}
	if Opts.Menu == "user" {
		var user User
		switch Opts.Type {
		case "create":
			user.Create()
		case "list":
			user.List()
		}
		os.Exit(0)
	}
	return false
}

// Parse 配置命令行
func Parse() {
	flag.StringVar(&Opts.File, "f", "setting.yaml", "配置文件路径")
	flag.BoolVar(&Opts.Version, "v", false, "打印当前版本")
	flag.BoolVar(&Opts.DB, "db", false, "迁移表结构")
	flag.StringVar(&Opts.Menu, "m", "", "菜单")
	flag.StringVar(&Opts.Type, "t", "", "类型")
	flag.Parse()
}
