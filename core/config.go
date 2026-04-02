package core

import (
	"fast_gin/configs"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig() (cfg *configs.Config) {
	cfg = new(configs.Config)
	byteData, err := os.ReadFile("setting.yaml")
	//fmt.Println(string(byteData))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		fmt.Printf("配置文件格式错误 %S", err)
		return
	}
	return
}
