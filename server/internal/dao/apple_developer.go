package dao

import (
	"supersign/internal/model"

	"gorm.io/gorm"
)

func newAppleDeveloper(db *gorm.DB) *appleDeveloper {
	return &appleDeveloper{db}
}

type appleDeveloper struct {
	db *gorm.DB
}

var _ model.AppleDeveloperStore = (*appleDeveloper)(nil)

func (a *appleDeveloper) Create(appleDeveloper *model.AppleDeveloper) error {
	return a.db.Create(appleDeveloper).Error
}

func (a *appleDeveloper) Del(iss string) error {
	return a.db.Where("iss = ?", iss).Delete(&model.AppleDeveloper{}).Error
}

func (a *appleDeveloper) AddCount(iss string, num int) error {
	return a.db.Model(&model.AppleDeveloper{}).
		Where("iss = ?", iss).
		UpdateColumn("count", gorm.Expr("count + ?", num)).Error
}

func (a *appleDeveloper) UpdateCount(iss string, count int) error {
	return a.db.Model(&model.AppleDeveloper{}).
		Where("iss = ?", iss).
		Update("count", count).Error
}

func (a *appleDeveloper) UpdateLimit(iss string, limit int) error {
	return a.db.Model(&model.AppleDeveloper{}).
		Where("iss = ?", iss).
		Update("`limit`", limit).Error
}

func (a *appleDeveloper) Enable(iss string, enable bool) error {
	return a.db.Model(&model.AppleDeveloper{}).
		Where("iss = ?", iss).
		Update("enable", enable).Error
}

func (a *appleDeveloper) Query(iss string) (*model.AppleDeveloper, error) {
	var appleDeveloper model.AppleDeveloper
	err := a.db.Where("iss = ?", iss).Take(&appleDeveloper).Error
	if err != nil {
		return nil, err
	}
	return &appleDeveloper, nil
}

func (a *appleDeveloper) GetUsable() (*model.AppleDeveloper, error) {
	var appleDeveloper model.AppleDeveloper
	err := a.db.Where("`limit` - count > 0 And count < ? And enable = ?", 100, true).
		Take(&appleDeveloper).Error
	if err != nil {
		return nil, err
	}
	return &appleDeveloper, nil
}

func (a *appleDeveloper) List(content string, page, pageSize *int) ([]model.AppleDeveloper, int64, error) {
	var (
		appleDevelopers []model.AppleDeveloper
		total           int64
	)
	if content == "" {
		err := a.db.Model(&model.AppleDeveloper{}).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		err = a.db.Scopes(paginate(page, pageSize)).Find(&appleDevelopers).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := a.db.Model(&model.AppleDeveloper{}).
			Where("iss LIKE ?", "%"+content+"%").
			Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		err = a.db.Scopes(paginate(page, pageSize)).
			Where("iss LIKE ?", "%"+content+"%").
			Find(&appleDevelopers).Error
		if err != nil {
			return nil, 0, err
		}
	}
	return appleDevelopers, total, nil
}