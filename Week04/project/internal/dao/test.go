package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gogeekbang/internal/model"
)

func (d *Dao) GetRow(id int) (test *model.Test, err error) {
	err = d.db.Where("id = ?", id).First(&test).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.Wrapf(ErrNoRows, "dao:GetRow id = %d", id)
		}
		return nil, errors.Wrapf(err, "dao:GetRow id = %d", id)
	}
	return test, nil
}
