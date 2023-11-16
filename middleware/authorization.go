package middleware

import (
	"MusicManager/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := util.AnalyseToken(auth)

		if err != nil {
			println("err:" + err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "AuthUserCheck Failed1",
			})
			c.Abort()
			return
		}

		if userClaim == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "AuthUserCheck Failed2",
			})
			c.Abort()
			return
		}

		c.Set("user", userClaim)
		c.Next()
	}
}
