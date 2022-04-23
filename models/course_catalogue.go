/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type CourseCatalogue struct {
	BaseModel
	UnitId int64  `json:"unit_id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Sort   int    `json:"sort"`
}

func (CourseCatalogue) TableName() string {
	return "dv_course_catalogue"
}
