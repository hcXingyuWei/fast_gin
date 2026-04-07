package core

import (
	"fast_gin/configs"
	"fast_gin/flags"
	"fast_gin/global"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// ReadConfig 读取配置文件
func ReadConfig() (cfg *configs.Config) {
	cfg = new(configs.Config)
	byteData, err := os.ReadFile(flags.Opts.File)
	if err != nil {
		logrus.Fatalf("配置文件读取错误 :%S", err)
		return
	}

	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		logrus.Fatalf("配置文件格式错误 :%S", err)
		return
	}
	return
}

// DumpConfig 写入配置文件
func DumpConfig() {
	byteData, err := yaml.Marshal(&global.Config)
	if err != nil {
		logrus.Fatalf("配置文件转换错误 :%S", err)
		return
	}
	err = os.WriteFile(flags.Opts.File, byteData, os.ModePerm)
	if err != nil {
		logrus.Fatalf("配置文件写入错误 :%S", err)
		return
	}
	logrus.Info("配置文件写入成功")
}
