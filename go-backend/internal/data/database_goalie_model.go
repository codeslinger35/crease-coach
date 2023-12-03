package data

import "database/sql"

type DatabaseGoalieModel struct {
	DB *sql.DB
}

func (model DatabaseGoalieModel) Save() error { return nil }

func (model DatabaseGoalieModel) Init() error { return nil }

func (model DatabaseGoalieModel) GetAll() ([]Goalie, error) {
	return nil, nil
}

func (model DatabaseGoalieModel) GetGoalie(id int64) (Goalie, error) {
	return Goalie{}, nil
}

func (model DatabaseGoalieModel) AddGoalie(goalie Goalie) (Goalie, error) {
	return Goalie{}, nil
}

func (model DatabaseGoalieModel) UpdateGoalie(goalie Goalie) (Goalie, error) {
	return Goalie{}, nil
}

func (model DatabaseGoalieModel) AddGame(game Game, goalieId int64, seasonId int64) (Game, error) {
	return Game{}, nil
}
