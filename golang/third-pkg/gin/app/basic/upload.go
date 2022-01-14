package basic

import (
	"gin-example/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Routes) Upload(c *gin.Context) {
	file, err := c.FormFile("img")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	common.IsExistFolder("img")

	err = c.SaveUploadedFile(file, "img/"+file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fileName": file.Filename,
		"size":     file.Size,
		"header":   file.Header,
	})
}

// 多图上传
func (r *Routes) UploadAll(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	files := form.File["img[]"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "请上传图片",
		})
		return
	}

	common.IsExistFolder("img")
	for _, v := range files {
		err = c.SaveUploadedFile(v, "img/"+v.Filename)
		if err != nil {
			break
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err":    nil,
		"lenght": len(files),
	})
}
