package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

func Manage(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"code":    200,
		"message": "",
	})
}

func UploadMusicFile(c *gin.Context) {

	folderPath := viper.GetString("savePath")
	fileName := c.Query("file-name")

	file, err := c.FormFile(fileName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(file, folderPath+"/"+file.Filename); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文件上传成功",
	})
}

func DeleteMusicFile(c *gin.Context) {

	fileName := c.Query("music-name")
	folderPath := viper.GetString("savePath")

	filePath := fmt.Sprintf(folderPath + "/" + fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "要删除的文件不存在，请尝试刷新列表",
		})
		return
	}

	err := os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功!",
	})
}

func GetMusicList(c *gin.Context) {
	folderPath := viper.GetString("savePath")
	length := len(folderPath + "\\")

	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取音乐列表失败",
		})
		return
	}

	for i, f := range files {
		files[i] = f[length:]
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取音乐列表成功",
		"data":    files,
	})
}
