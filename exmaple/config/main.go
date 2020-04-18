package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type Config struct {
	Mysql struct {
		DriverName string `yaml:"driverName"`
		Address string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}

	App struct {
		Server struct {
			Port int32 `yaml:"port"`
		}
		Name string `yaml:"name"`
	}
}

func GetCurrentFilePath(fileName string, skip int) string {
	//获取当前函数Caller reports，取得当前调用对应的文件
	_, f, _, _ := runtime.Caller(skip)
	//解析出所在目录
	dir := filepath.Dir(f)
	//组装配置文件路径
	file := filepath.Join(dir, fileName)
	return file
}

func main() {
	conf := &Config{}
	confFilePath := GetCurrentFilePath("config.yaml", 1)
	yamlFile, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v\n", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v\n", err)
	}

	fmt.Printf("%v, %+v\n", conf)
}
