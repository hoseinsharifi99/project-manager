package db_manager

import "amani/model"

func (d *DbManager) AddUser(user *model.User) error {
	if err := d.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
