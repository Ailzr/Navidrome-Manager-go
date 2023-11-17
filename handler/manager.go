package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func Manage(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"code":    200,
		"message": "",
	})
}

func UploadMusicFile(c *gin.Context) {

	folderPath := viper.GetString("savePath")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(file, folderPath+"/"+file.Filename); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "存储文件发生错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文件上传成功",
	})
	//form, err := c.MultipartForm()
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    200,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//files := form.File["file"]
	//var failFiles []string
	//Fail := false
	//
	//for _, file := range files {
	//	filename := filepath.Join(folderPath, file.Filename)
	//	if err := c.SaveUploadedFile(file, filename); err != nil {
	//		fmt.Println(err.Error())
	//		failFiles = append(failFiles, file.Filename)
	//		Fail = true
	//		continue
	//	}
	//}

	//if !Fail {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    200,
	//		"message": "文件上传成功",
	//	})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code":    -1,
	//	"message": "部分文件上传失败",
	//	"data":    failFiles,
	//})

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

type FileInfo struct {
	Name    string    `json:"name"`
	ModTime time.Time `json:"mod_time"`
}

func GetMusicList(c *gin.Context) {
	folderPath := viper.GetString("savePath")

	//var files []string
	//err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
	//	if err != nil {
	//		return err
	//	}
	//	if !info.IsDir() {
	//		files = append(files, path)
	//	}
	//	return nil
	//})

	files, err := os.ReadDir(folderPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取音乐列表失败",
		})
		return
	}

	var fileInfos []FileInfo

	for _, f := range files {
		filePath := filepath.Join(folderPath + "/" + f.Name())
		fileStat, err := os.Stat(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    -1,
				"message": "获取音乐列表失败",
			})
			return
		}
		fileInfo := FileInfo{f.Name(), fileStat.ModTime()}
		fileInfos = append(fileInfos, fileInfo)
	}

	sort.Sort(ByModTime(fileInfos))

	var filesName []string
	for _, fi := range fileInfos {
		filesName = append(filesName, fi.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取音乐列表成功",
		"data":    filesName,
	})
}

type ByModTime []FileInfo

func (bmt ByModTime) Len() int           { return len(bmt) }
func (bmt ByModTime) Swap(i, j int)      { bmt[i], bmt[j] = bmt[j], bmt[i] }
func (bmt ByModTime) Less(i, j int) bool { return bmt[i].ModTime.After(bmt[j].ModTime) }
