/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package routers

import (
	"DigitalVillage/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.RouterGroup) {

	apiGroup := r.Group("/admin/")
	apiGroup.POST("login", admin.Controller{}.Login)
	apiGroup.POST("upload", admin.Controller{}.Upload)
	apiGroup.GET("class/all", admin.ClassController{}.All)

	apiGroup.GET("class/list", admin.ClassController{}.List)
	apiGroup.POST("class/detail", admin.ClassController{}.Detail)
	apiGroup.POST("class/delete", admin.ClassController{}.Delete)
	apiGroup.POST("class/option", admin.ClassController{}.Option)

	apiGroup.GET("course/list", admin.CourseController{}.List)
	apiGroup.POST("course/detail", admin.CourseController{}.Detail)
	apiGroup.POST("course/delete", admin.CourseController{}.Delete)
	apiGroup.POST("course/option", admin.CourseController{}.Option)

	apiGroup.GET("unit/list", admin.UnitController{}.List)
	apiGroup.POST("unit/detail", admin.UnitController{}.Detail)
	apiGroup.POST("unit/delete", admin.UnitController{}.Delete)
	apiGroup.POST("unit/option", admin.UnitController{}.Option)

	apiGroup.GET("catalogue/list", admin.CatalogueController{}.List)
	apiGroup.POST("catalogue/detail", admin.CatalogueController{}.Detail)
	apiGroup.POST("catalogue/delete", admin.CatalogueController{}.Delete)
	apiGroup.POST("catalogue/option", admin.CatalogueController{}.Option)

}
