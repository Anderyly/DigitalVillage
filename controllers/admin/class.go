/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"DigitalVillage/ay"
	"DigitalVillage/models"
	"github.com/gin-gonic/gin"
	"strings"
)

type ClassController struct {
}

type noticeTypeListForm struct {
	Page  int `form:"page"`
	Count int `form:"count"`
}

// List 列表
func (con ClassController) List(c *gin.Context) {
	var data noticeTypeListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var list []models.Class

	var count int64
	ay.Db.Order("id desc").
		Limit(data.Count).
		Offset((data.Page - 1) * data.Count).
		Find(&list)

	ay.Db.Model(models.Class{}).Count(&count)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

type DetailForm struct {
	Id int `form:"id"`
}

// Detail 详情
func (con ClassController) Detail(c *gin.Context) {
	var data DetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.Class

	ay.Db.First(&res, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

// Option 添加 编辑
func (con ClassController) Option(c *gin.Context) {
	type optionForm struct {
		Id   int    `form:"id"`
		Name string `form:"name"`
	}
	var data optionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.Class
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Name = data.Name
		ay.Db.Save(&res)
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	} else {
		ay.Db.Create(&models.Class{
			Name: data.Name,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

type DeleteForm struct {
	Id string `form:"id"`
}

func (con ClassController) Delete(c *gin.Context) {
	var data DeleteForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	idArr := strings.Split(data.Id, ",")

	for _, v := range idArr {
		var res models.Class
		ay.Db.Delete(&res, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con ClassController) All(c *gin.Context) {
	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}
	type list struct {
		Label string `gorm:"column:name" json:"label"`
		Value int64  `gorm:"column:id" json:"value"`
	}
	var l []list
	ay.Db.Model(models.Class{}).Find(&l)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": l,
	})
}
