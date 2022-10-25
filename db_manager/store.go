package db_manager

import "amani/model"

func (d *DbManager) AddUser(user *model.User) error {
	if err := d.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *DbManager) AddProject(prj *model.Project) error {
	return s.db.Create(prj).Error
}
