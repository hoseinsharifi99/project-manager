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

func (d *DbManager) GetUserProjects(userId, prjId uint) (*model.UserProject, error) {
	prj := new(model.UserProject)
	if err := d.db.Where("user_id = ? AND project_id = ?", userId, prjId).Find(&prj).Error; err != nil {
		return nil, err
	}
	return prj, nil
}

func (d *DbManager) UpdateUserProject(uprj *model.UserProject) error {
	return d.db.Model(uprj).Update(uprj).Error
}
