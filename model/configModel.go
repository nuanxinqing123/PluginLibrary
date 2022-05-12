// -*- coding: utf-8 -*-
// @Time    : 2022/5/12 15:03
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : configModel.go

package model

type Config struct {
	// 运行端口
	Port int `json:"Port"`
	// 上传密码
	PassWord string `json:"PassWord"`
	// 网站公告
	Notice string `json:"Notice"`
}
