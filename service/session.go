/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetSession(r *gin.Engine) *gin.Engine {

	store := cookie.NewStore([]byte("Andreyly"))
	r.Use(sessions.Sessions("sb", store))

	return r

}
