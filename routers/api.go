/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package routers

import (
	"DigitalVillage/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.RouterGroup) {

	apiGroup := r.Group("/api/")
	apiGroup.GET("course/recommend", api.CourseController{}.Recommend)
	apiGroup.GET("course/all", api.CourseController{}.All)
	apiGroup.GET("course/list", api.CourseController{}.List)
	apiGroup.POST("course/catalogue", api.CourseController{}.Catalogue)

}
