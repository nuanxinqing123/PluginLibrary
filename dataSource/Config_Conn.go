// -*- coding: utf-8 -*-
// @Time    : 2022/5/12 15:02
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : Config_Conn.go

package dataSource

import (
	"PluginLibrary/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadConfig() *model.Config {
	// 创建对象
	config := model.Config{}

	// 打开文件
	file, err := os.Open("config/config.json")
	if err != nil {
		// 打开文件时发生错误
		panic(err)
	}
	// 延迟关闭
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 配置读取
	byteData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		// 读取配置时发生错误
		panic(err)
	}

	// 数据绑定
	err3 := json.Unmarshal(byteData, &config)
	if err3 != nil {
		// 数据绑定时发生错误
		panic(err)
	}

	return &config
}
