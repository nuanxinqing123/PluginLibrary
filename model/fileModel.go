// -*- coding: utf-8 -*-
// @Time    : 2022/5/12 15:20
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : fileModel.go

package model

// FileInfo 读取插件信息
type FileInfo struct {
	// 文件名
	FileName string `json:"FileName"`
	// 插件名
	PluginName string `json:"PluginName"`
	// 文件最后修改时间
	FileUploadTime string `json:"FileUploadTime"`
}
