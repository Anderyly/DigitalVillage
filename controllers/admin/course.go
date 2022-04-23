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

type CourseController struct {
}

type courseListForm struct {
	Page  int    `form:"page"`
	Count int    `form:"count"`
	Key   string `form:"key"`
	Type  int64  `form:"type"`
}

// List 列表
func (con CourseController) List(c *gin.Context) {
	var data courseListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	type returnList struct {
		models.Course
		TypeName string `json:"type_name"`
	}
	var list []returnList

	var count int64
	res := ay.Db.Model(models.Course{}).
		Select("dv_course.*,dv_class.name as type_name").
		Joins("left join dv_class on dv_course.cid=dv_class.id").
		Order("id desc").
		Limit(data.Count).
		Offset((data.Page - 1) * data.Count)

	if data.Key != "" {
		res.Where("dv_course.title like ?", "%"+data.Key+"%")
	}

	if data.Type != 0 {
		res.Where("dv_course.cid = ?", data.Type)
	}
	res.Find(&list)

	ay.Db.Model(models.Course{}).Count(&count)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

// Detail 详情
func (con CourseController) Detail(c *gin.Context) {
	var data DetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	type returnList struct {
		models.Course
		TypeName string `json:"type_name"`
	}

	var res returnList

	ay.Db.Model(models.Course{}).
		Select("dv_course.*,dv_class.name as type_name").
		Joins("left join dv_class on dv_course.cid=dv_class.id").
		Where("dv_course.id = ?", data.Id).
		First(&res)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

// Option 添加 编辑
func (con CourseController) Option(c *gin.Context) {
	type optionForm struct {
		Id          int    `form:"id"`
		Cid         int64  `form:"cid"`
		Cover       string `form:"cover"`
		Title       string `form:"title"`
		Label       string `form:"label"`
		IsRecommend int    `form:"is_recommend"`
		Status      int    `form:"status"`
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

	var res models.Course
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Title = data.Title
		res.Cover = data.Cover
		res.Cid = data.Cid
		res.Label = data.Label
		res.Status = data.Status
		res.IsRecommend = data.IsRecommend

		ay.Db.Save(&res)
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	} else {
		ay.Db.Create(&models.Course{
			Title:       data.Title,
			Cover:       data.Cover,
			Cid:         data.Cid,
			Label:       data.Label,
			Status:      data.Status,
			IsRecommend: res.IsRecommend,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con CourseController) Delete(c *gin.Context) {
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
		var res models.Course
		ay.Db.Delete(&res, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}
