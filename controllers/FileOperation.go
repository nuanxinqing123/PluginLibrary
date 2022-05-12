// -*- coding: utf-8 -*-
// @Time    : 2022/5/12 15:13
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : FileOperation.go

package controllers

import (
	"PluginLibrary/dataSource"
	"PluginLibrary/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// GetWebAllData 获取网站所有信息
func GetWebAllData(c *gin.Context) {
	// 读取配置文件
	conf := dataSource.LoadConfig()

	// 获取插件目录绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"code": 500,
		})
		return
	}
	// 完整文件
	PluginPath := ExecPath + "/plugin/"

	// 读取目录
	var fl []model.FileInfo
	var fi model.FileInfo
	files, _ := ioutil.ReadDir(PluginPath)
	for _, f := range files {
		// 跳过不是JS的文件
		if !strings.Contains(f.Name(), ".js") {
			continue
		}

		// 读取插件名称
		fd, err := os.Open(PluginPath + f.Name())
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(200, gin.H{
				"code": 500,
			})
			return
		}
		defer fd.Close()

		// 读取插件名
		v, _ := ioutil.ReadAll(fd)
		data := string(v)
		PluginName := ""
		if regs := regexp.MustCompile(`\[name:(.+)]`).FindStringSubmatch(data); len(regs) != 0 {
			PluginName = strings.Trim(regs[1], " ")
		}

		// 读取文件更新日期
		stat, err := fd.Stat()
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(200, gin.H{
				"code": 500,
			})
			return
		}

		fi.FileName = f.Name()
		fi.PluginName = PluginName
		fi.FileUploadTime = SwitchTimeStampToDataYear(stat.ModTime().Unix())
		fl = append(fl, fi)
	}

	// 返回网站信息
	c.JSON(200, gin.H{
		"code":       200,
		"notice":     conf.Notice,
		"pluginList": fl,
	})
}

// DownloadPlugin 下载插件
func DownloadPlugin(c *gin.Context) {
	// 获取文件名
	filename := c.Query("filename")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	// 获取插件目录绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"code": 500,
		})
		return
	}
	// 完整文件
	PluginPath := ExecPath + "/plugin/"

	c.File(PluginPath + filename)
}

// UploadPlugin 上传插件
func UploadPlugin(c *gin.Context) {
	// 获取参数
	pwd := c.Query("pwd")
	file, _ := c.FormFile("file")

	// 读取配置文件
	conf := dataSource.LoadConfig()
	// 判断密码是否正确
	if pwd != conf.PassWord {
		// 密码错误
		c.JSON(200, gin.H{
			"code": 400,
		})
		return
	}

	// 获取插件目录绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"code": 500,
		})
		return
	}
	// 完整文件
	PluginPath := ExecPath + "/plugin/"

	err = c.SaveUploadedFile(file, PluginPath+file.Filename)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"code": 500,
		})
		return
	}

	// 返回网站信息
	c.JSON(200, gin.H{
		"code": 200,
	})
}

// SwitchTimeStampToDataYear 将传入的时间戳转为时间
func SwitchTimeStampToDataYear(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
