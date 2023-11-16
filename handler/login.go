package handler

import (
	"MusicManager/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func LoginHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"code": 200,
		"msg":  "",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
	}

	identity := viper.GetString("user.Identity")
	name := viper.GetString("user.Name")
	pwd := viper.GetString("user.Password")

	fmt.Println(username + "\n" + password)
	fmt.Println(name + "\n" + pwd)

	if strings.Compare(name, username) != 0 || strings.Compare(password, pwd) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "conf Error",
		})
		return
	}

	token, err := util.GenerateToken(identity, name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Generate Token Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
