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

func (s *DbManager) AddUserProjec(userProject *model.UserProject) error {
	return s.db.Create(userProject).Error
}

func (s *DbManager) AddTask(task *model.Task) error {
	return s.db.Create(task).Error
}
