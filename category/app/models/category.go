package models

import (
	orm "gin-blog-category/database"
	"fmt"
)

type Category struct {
	Id         int64     `gorm:"column:id" view:"ID"`
	Name       string    `gorm:"column:name" view:"分类名称"`
	Slug       string    `gorm:"column:slug"`
	OrderNum   int64	 `gorm:"column:order_num" view:"排序"`
}

func (Category) TableName() string {
  return "category"
}

var Categorys []Category

//添加
func (category Category) Insert() (id int64,err error) {
	result := orm.Db.Create(&category)
	id =category.Id
	if result.Error != nil {
		fmt.Println(category)
        err = result.Error
        return
    }
    return
}

//列表
func (category *Category) Categorys() (categorys []Category, err error) {
    if err = orm.Db.Find(&categorys).Error; err != nil {
        return
    }
    return
}



