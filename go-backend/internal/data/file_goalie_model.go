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
	return (*model.DB)[model.getIndexForGoalieId(id)], nil
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
	i := model.getIndexForGoalieId(goalie.Id)

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

func (model FileGoalieModel) GetGames(goalieId int64, seasonId int64) ([]Game, error) {
	goalieIndex := model.getIndexForGoalieId(goalieId)
	seasonIndex := model.getIndexForSeasonId(seasonId, goalieIndex)

	return (*model.DB)[goalieIndex].Seasons[seasonIndex].Games, nil
}

func (model FileGoalieModel) AddGame(game Game, goalieId int64, seasonId int64) (Game, error) {
	goalieIndex := model.getIndexForGoalieId(goalieId)
	seasonIndex := model.getIndexForSeasonId(seasonId, goalieIndex)

	(*model.DB)[goalieIndex].Seasons[seasonIndex].Games = append((*model.DB)[goalieIndex].Seasons[seasonIndex].Games, game)

	err := model.Save()
	if err != nil {
		return Game{}, err
	}

	return (*model.DB)[goalieIndex].Seasons[seasonIndex].Games[model.getIndexForGameId(game.Id, goalieIndex, seasonIndex)], nil
}

func (model FileGoalieModel) UpdateGame(game Game, goalieId int64, seasonId int64, gameId int64) (Game, error) {
	goalieIndex := model.getIndexForGoalieId(goalieId)
	seasonIndex := model.getIndexForSeasonId(seasonId, goalieIndex)
	gameIndex := model.getIndexForGameId(game.Id, goalieIndex, seasonIndex)

	(*model.DB)[goalieIndex].Seasons[seasonIndex].Games[gameIndex] = game

	return (*model.DB)[goalieIndex].Seasons[seasonIndex].Games[model.getIndexForGameId(game.Id, goalieIndex, seasonIndex)], nil
}

func (model FileGoalieModel) getIndexForGoalieId(id int64) int {
	index := -1
	for i := range *model.DB {
		if (*model.DB)[i].Id == id {
			index = i
			break
		}
	}
	return index
}

func (model FileGoalieModel) getIndexForSeasonId(id int64, goalieIndex int) int {
	index := -1
	for i := range (*model.DB)[goalieIndex].Seasons {
		if (*model.DB)[goalieIndex].Seasons[i].Id == id {
			index = i
			break
		}
	}
	return index
}

func (model FileGoalieModel) getIndexForGameId(id int64, goalieIndex int, seasonIndex int) int {
	index := -1
	for i := range (*model.DB)[goalieIndex].Seasons[seasonIndex].Games {
		if (*model.DB)[goalieIndex].Seasons[seasonIndex].Games[i].Id == id {
			index = i
			break
		}
	}
	return index
}
