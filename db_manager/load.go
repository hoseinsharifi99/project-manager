package db_manager

import "amani/model"

func (d *DbManager) GetUserByUsername(userName string) (*model.User, error) {
	user := new(model.User)
	if err := d.db.First(user, model.User{Username: userName}).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (d *DbManager) GetUserByID(id uint) (*model.User, error) {
	user := new(model.User)
	if err := d.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DbManager) GetProjectByName(name string) (*model.Project, error) {
	prj := new(model.Project)
	if err := d.db.First(prj, model.Project{Name: name}).Error; err != nil {
		return nil, err
	}
	return prj, nil
}

func (d *DbManager) GetProjectById(id uint) (*model.Project, error) {
	prj := new(model.Project)
	if err := d.db.Where("id = ?", id).Find(&prj).Error; err != nil {
		return nil, err
	}
	return prj, nil
}

func (d *DbManager) GetTaskById(id uint) (*model.Task, error) {
	tsk := new(model.Task)
	if err := d.db.Where("id = ?", id).Find(&tsk).Error; err != nil {
		return nil, err
	}
	return tsk, nil
}

func (d *DbManager) GetTaskByName(name string) (*model.Task, error) {
	tsk := new(model.Task)
	if err := d.db.First(tsk, model.Task{Name: name}).Error; err != nil {
		return nil, err
	}
	return tsk, nil
}

func (d *DbManager) GetUserProjects(userId, prjId, tskId uint) (*model.UserProject, error) {
	prj := new(model.UserProject)
	if err := d.db.Where("user_id = ? AND project_id = ? AND task_id = ?", userId, prjId, tskId).Find(&prj).Error; err != nil {
		return nil, err
	}
	return prj, nil
}

func (d *DbManager) UpdateUserProject(uprj *model.UserProject) error {
	return d.db.Model(uprj).Update(uprj).Error
}

func (d *DbManager) GetProjects() ([]model.Project, error) {
	var projects []model.Project
	if err := d.db.Model(&model.Project{}).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (d *DbManager) GetTaskByUserID(id uint) ([]model.UserProject, error) {
	var userproject []model.UserProject
	if err := d.db.Model((&model.UserProject{})).Where("user_id = ?", id).Find(&userproject).Error; err != nil {
		return nil, err
	}
	return userproject, nil
}
