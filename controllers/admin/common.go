/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"DigitalVillage/ay"
	"DigitalVillage/controllers/api"
	"DigitalVillage/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strings"
	"time"
)

type CommonController struct {
}

func Auth() bool {
	token := ay.AuthCode(api.Token, "DECODE", "", 0)
	if token == "" {
		log.Println("没有token")
		return false
	}
	var user models.Admin
	ay.Db.First(&user, "account = ?", token)
	if user.Id == 0 {
		return false
	} else {
		return true
	}
}

func Upload(c *gin.Context, address string) (int, string) {
	file, err := c.FormFile("file")
	if err != nil {
		return 400, err.Error()
	}

	fileExt := strings.ToLower(path.Ext(file.Filename))
	log.Println(fileExt)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" && fileExt != ".xlsx" && fileExt != ".xls" {
		return 400, "上传失败!只允许png,jpg,gif,jpeg,xls,xlsx文件"
	}
	fileName := ay.MD5(fmt.Sprintf("%s%s", file.Filename, time.Now().String()))
	fileDir := fmt.Sprintf("static/upload/"+address+"/%d-%d/", time.Now().Year(), time.Now().Month())

	err = ay.CreateMutiDir(fileDir)
	if err != nil {
		return 400, err.Error()
	}

	filepath := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
	c.SaveUploadedFile(file, filepath)

	return 200, "/" + filepath

}
