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

type CatalogueController struct {
}

type catalogueListForm struct {
	UnitId int64 `form:"unit_id"`
}

// List 列表
func (con CatalogueController) List(c *gin.Context) {
	var data catalogueListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var list []models.CourseCatalogue

	var count int64

	ay.Db.Where("unit_id = ?", data.UnitId).Find(&list)

	ay.Db.Model(models.CourseUnit{}).Count(&count)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

// Detail 详情
func (con CatalogueController) Detail(c *gin.Context) {
	var data DetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.CourseCatalogue

	ay.Db.First(&res)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

// Option 添加 编辑
func (con CatalogueController) Option(c *gin.Context) {
	type optionForm struct {
		Id     int64  `form:"id"`
		Name   string `form:"name"`
		Link   string `form:"Link"`
		Sort   int    `form:"sort"`
		UnitId int64  `form:"unit_id"`
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

	var res models.CourseCatalogue
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Name = data.Name
		res.Sort = data.Sort
		res.UnitId = data.UnitId
		res.Link = data.Link

		ay.Db.Save(&res)
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	} else {
		ay.Db.Create(&models.CourseCatalogue{
			Name:   data.Name,
			Sort:   data.Sort,
			UnitId: data.UnitId,
			Link:   data.Link,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con CatalogueController) Delete(c *gin.Context) {
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
		var res models.CourseCatalogue
		ay.Db.Delete(&res, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}
