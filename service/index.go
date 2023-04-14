package service

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	log.Print("'get ...'")
	c.JSON(200, gin.H{
		"message": "index",
	})
}

func GetHtml(c *gin.Context) {
	log.Print("gethtml")
	c.HTML(http.StatusOK, "index2.html", gin.H{
		"msg": "hello gin",
	})
}

func UpLoadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Print(file.Header)

	name := c.PostForm("name")
	email := c.PostForm("email")

	// err := c.SaveUploadedFile(file, "D:\\GoWorks\\src\\Stest")
	// if err != nil {
	// 	log.Panic(err)
	// }// Source

	// if err != nil {
	// 	c.String(http.StatusBadRequest, "get form err: %s", err.Error())
	// 	return
	// }

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "./static/upload/"+filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
}
