package data

import (
	"encoding/json"
	"os"
)

type FileGoalieModel struct {
	DB   *[]Goalie
	File string
}

func (model FileGoalieModel) Save() error {
	b, err := json.Marshal(model.DB)
	if err != nil {
		return err
	}

	err = os.WriteFile(model.File, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (model FileGoalieModel) Init() error {
	content, err := os.ReadFile(model.File)
	if err != nil {
		return err
	}
	var payload []Goalie
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return err
	}

	*model.DB = payload

	return nil
}

func (model FileGoalieModel) GetAll() ([]Goalie, error) {
	return *model.DB, nil
}

func (model FileGoalieModel) GetGoalie(id int64) (Goalie, error) {
	return (*model.DB)[model.getIndexForId(id)], nil
}

func (model FileGoalieModel) AddGoalie(goalie Goalie) (Goalie, error) {
	*model.DB = append((*model.DB), goalie)

	err := model.Save()
	if err != nil {
		return Goalie{}, err
	}

	return goalie, nil
}

func (model FileGoalieModel) UpdateGoalie(goalie Goalie) (Goalie, error) {
	i := model.getIndexForId(goalie.Id)

	(*model.DB)[i].FirstName = goalie.FirstName
	(*model.DB)[i].LastName = goalie.LastName
	(*model.DB)[i].Age = goalie.Age
	(*model.DB)[i].Team = goalie.Team

	err := model.Save()
	if err != nil {
		return Goalie{}, err
	}

	return (*model.DB)[i], nil
}

func (model FileGoalieModel) getIndexForId(id int64) int {
	index := -1
	for i := range *model.DB {
		if (*model.DB)[i].Id == id {
			index = i
			break
		}
	}
	return index
}

func (model *FileGoalieModel) updateModel(newModel []Goalie) {
	*model.DB = newModel
}
