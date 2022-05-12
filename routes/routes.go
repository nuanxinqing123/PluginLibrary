// -*- coding: utf-8 -*-
// @Time    : 2022/5/12 14:56
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : routes.go

package routes

import (
	"PluginLibrary/bindata"
	"PluginLibrary/controllers"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

func Setup() *gin.Engine {
	// 创建服务
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 前端静态文件
	{
		// 加载模板文件
		t, err := loadTemplate()
		if err != nil {
			panic(err)
		}
		r.SetHTMLTemplate(t)

		// 加载静态文件
		fs := assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: nil,
			Prefix:    "assets",
		}
		r.StaticFS("/static", &fs)

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	// 路由组
	{
		// 获取网站所有信息
		r.GET("web/all", controllers.GetWebAllData)

		// 下载插件
		r.GET("file/download", controllers.DownloadPlugin)

		// 上传插件
		r.POST("file/upload", controllers.UploadPlugin)
	}

	return r
}

// loadTemplate 加载模板文件
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := bindata.Asset(name)
		if err != nil {
			continue
		}
		name := strings.Replace(name, "assets/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
