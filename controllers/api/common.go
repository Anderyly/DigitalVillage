/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"DigitalVillage/ay"
)

var (
	Token string
)

type CommonController struct {
}

func GetToken(token string) string {
	uid := ay.AuthCode(token, "DECODE", "", 0)
	return uid
}
