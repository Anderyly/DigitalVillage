/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type Class struct {
	BaseModel
	Name string `json:"name"`
}

func (Class) TableName() string {
	return "dv_class"
}
