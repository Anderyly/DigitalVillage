/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type CourseUnit struct {
	BaseModel
	CourseId int64  `json:"course_id"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
}

func (CourseUnit) TableName() string {
	return "dv_course_unit"
}
