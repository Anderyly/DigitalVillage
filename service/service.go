/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package service

import (
	"github.com/gin-gonic/gin"
)

func Set(r *gin.Engine) *gin.Engine {
	r.Use(Cors())
	r.Use(Header())
	//r.Use(Pretreatment())
	return r
}
