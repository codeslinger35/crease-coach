package data

import "database/sql"

type Models struct {
	Goalies GoalieModel
}

func NewModels(filename string, db *sql.DB) Models {
	if db == nil {
		return Models{
			Goalies: FileGoalieModel{DB: new([]Goalie), File: filename}}
	}

	return Models{
		Goalies: DatabaseGoalieModel{DB: db}}
}
