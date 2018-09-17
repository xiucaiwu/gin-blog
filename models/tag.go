package models

import (
	"github.com/jinzhu/gorm"
	//"time"
)

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return 
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddTag(name string, state int, createdBy string) bool{
	db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})
	return true
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// 删除所有软删除的数据
func CleanAllTag() error {
	// 硬删除要使用 Unscoped()，这是 GORM 的约定
	err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{}).Error
	if err != nil {
		return err
	}
	return nil
}

//gorm的Callbacks
//gorm所支持的回调方法：
//创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
//更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
//删除：BeforeDelete、AfterDelete
//查询：AfterFind
//func (tag *Tag) BeforeCreate(scope *gorm.Scope) error  {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//	return nil
//}