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

type UnitController struct {
}

type unitListForm struct {
	CourseId int64 `form:"course_id"`
}

// List 列表
func (con UnitController) List(c *gin.Context) {
	var data unitListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var list []models.CourseUnit

	var count int64

	ay.Db.Where("course_id = ?", data.CourseId).Find(&list)

	ay.Db.Model(models.CourseUnit{}).Count(&count)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

// Detail 详情
func (con UnitController) Detail(c *gin.Context) {
	var data DetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.CourseUnit

	ay.Db.First(&res)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

// Option 添加 编辑
func (con UnitController) Option(c *gin.Context) {
	type optionForm struct {
		Id       int64  `form:"id"`
		Name     string `form:"name"`
		Sort     int    `form:"sort"`
		CourseId int64  `form:"course_id"`
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

	var res models.CourseUnit
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Name = data.Name
		res.Sort = data.Sort
		res.CourseId = data.CourseId

		ay.Db.Save(&res)
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	} else {
		ay.Db.Create(&models.CourseUnit{
			Name:     data.Name,
			Sort:     data.Sort,
			CourseId: data.CourseId,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con UnitController) Delete(c *gin.Context) {
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
		var res models.CourseUnit
		ay.Db.Delete(&res, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}
