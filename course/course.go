// 该文件表示我们要访问的资源
package course

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	Name     string `json:"name"form:"name"binding:"required,alphaunicode"`
	Teacher  string `json:"teacher"form:"teacher"binding:"required,alphaunicode"`
	Duration string `json:"duration"form:"duration"binding:"required,alphaunicode"`
}

func Add(c *gin.Context) {
	req := &course{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
	fmt.Println(*req)
}

func Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Update(c *gin.Context) {
	req := &course{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func Delete(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
