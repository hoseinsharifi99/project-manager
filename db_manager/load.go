package db_manager

import "amani/model"

func (d *DbManager) GetUserByUsername(userName string) (*model.User, error) {
	user := new(model.User)
	if err := d.db.First(user, model.User{Username: userName}).Error; err != nil {
		return nil, err
	}
	return user, nil
}
