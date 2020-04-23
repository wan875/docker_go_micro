package models

import (
	orm "gin-blog-post/database"
	"time"
	"fmt"
)

type Post struct {
	Id              int64     `gorm:"column:id" view:"ID"`
	PostTitle       string    `gorm:"column:post_title" valid:"Required; MaxSize(255)" view:"标题"`
	PostExcerpt     string    `gorm:"column:post_excerpt"`
	PostContent     string    `gorm:"column:post_content" view:"内容"`
	Tags            string    `gorm:"column:tags" view:"标签"`
	Categorys       string    `gorm:"column:categorys" view:"分类"`
	IsPublished     int64     `gorm:"column:is_published" view:"是否已发布"`
	Poster          string    `gorm:"column:poster" view:"海报"`
	CreatedAt       time.Time `gorm:"column:created_at`
	//ModifiedAt      time.Time `gorm:"column:modified_at`
	//PublishedAt     time.Time `gorm:"column:published_at" view:"发布时间"`
	Views           int64     `gorm:"column:views" view:"阅读数"`
	PostAuthor      int64     `gorm:"column:post_author" view:"发布者"`
	WordsNum        int64     `gorm:"column:words_num"`
	IsShowFirst     int64     `gorm:"column:is_show_first" view:"是否首页显示（如果是专栏文章，设置首页值为1，否则为0）改字段主要针对专栏设置到首页"`
}

func (Post) TableName() string {
  return "post"
}

var Posts []Post

//添加
func (post Post) Insert() (id int64,err error) {
	result := orm.Db.Create(&post)
	id =post.Id
	if result.Error != nil {
		fmt.Println(post)
        err = result.Error
        return
    }
    return
}

//列表
func (post *Post) Posts(page int64) (posts []Post, err error) {
	fmt.Println(page)
	var limit int64
	limit = 10
	offset := (page-1)*limit
    if err = orm.Db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
        return
    }
    return
}
