package router

import (
	"MusicManager/handler"
	"MusicManager/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter() {
	r := gin.Default()

	r.LoadHTMLGlob("./manage/static/html/*")
	r.Static("/manage/static", "./manage/static")
	r.MaxMultipartMemory = 64 << 20 //设置上传文件最大为64MB

	rManage := r.Group("/manage")
	{
		rManage.GET("/login", handler.LoginHtml)
		rManage.POST("/manage-login", handler.Login)

		rManage.GET("/manage", handler.Manage)
		rManage.GET("/music-list", middleware.AuthCheck(), handler.GetMusicList)
		rManage.POST("/upload-music", middleware.AuthCheck(), handler.UploadMusicFile)
		rManage.DELETE("/delete-music", middleware.AuthCheck(), handler.DeleteMusicFile)
	}

	_ = r.Run()
}
