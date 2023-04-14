package main

import (
	"fmt"
	"io/ioutil"
	Config "mygo/cmd/myapp/day9/readyaml"
	_ "time"

	"gopkg.in/yaml.v2"
)

func main() {
	// 读取yaml文件
	//data, err := ioutil.ReadFile("D:\\pro\\mygo\\cmd\\myapp\\day9\\readyaml\\config.yaml")
	data, err := ioutil.ReadFile("..\\config.yaml")
	if err != nil {
		panic(err)
	}

	// 解析yaml文件到结构体中
	var config Config.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// 打印解析出来的结果
	fmt.Printf("Server: %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("Database: %s:%d, name=%s, timeout=%v\n", config.Database.Host, config.Database.Port, config.Database.Name, config.Database.Options.Timeout)
}
