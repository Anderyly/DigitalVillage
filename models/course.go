/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type Course struct {
	BaseModel
	Cid         int64  `json:"cid"`
	Cover       string `json:"cover"`
	Title       string `json:"title"`
	Label       string `json:"label"`
	IsRecommend int    `json:"is_recommend"`
	Status      int    `json:"status"`
}

func (Course) TableName() string {
	return "dv_course"
}
