/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"DigitalVillage/ay"
	"DigitalVillage/models"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
}

type recommendListForm struct {
	Page  int `form:"page"`
	Count int `form:"count"`
}

// Recommend 获取推荐
func (con CourseController) Recommend(c *gin.Context) {
	var data recommendListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var list []models.Course
	ay.Db.Where("is_recommend = 1 and status = 1").Order("created_at desc").Find(&list)

	if list == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": list,
		})

	}

}

// All 获取类型
func (con CourseController) All(c *gin.Context) {
	var data listForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var list []models.Class
	ay.Db.Find(&list)

	if list == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": list,
		})

	}

}

type listForm struct {
	Page  int `form:"page"`
	Count int `form:"count"`
	Type  int `form:"type"`
}

// List 获取类型下数据
func (con CourseController) List(c *gin.Context) {
	var data listForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var list []models.Course
	ay.Db.Where("cid = ? and status = 1", data.Type).Order("created_at desc").Find(&list)

	if list == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": list,
		})

	}

}

type catalogueForm struct {
	Id int `form:"id"`
}

// Catalogue List 获取目录
func (con CourseController) Catalogue(c *gin.Context) {
	var data catalogueForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var list []map[string]interface{}
	var catalogure []map[string]interface{}

	// 获取所有单元
	var unit []models.CourseUnit
	ay.Db.Where("course_id = ?", data.Id).Order("sort asc").Find(&unit)

	for _, v := range unit {
		var res []models.CourseCatalogue
		ay.Db.Where("unit_id = ?", v.Id).Order("sort asc").Find(&res)
		for _, v1 := range res {
			catalogure = append(catalogure, gin.H{
				"name": v1.Name,
				"link": v1.Link,
			})
		}
		list = append(list, gin.H{
			"unit":      v.Name,
			"catalogue": catalogure,
		})
		catalogure = []map[string]interface{}{}
	}

	if list == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": list,
		})

	}

}
