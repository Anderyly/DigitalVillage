/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type Admin struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (Admin) TableName() string {
	return "dv_admin"
}
