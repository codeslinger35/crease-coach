package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codeslinger35/ccapi/internal/data"
	"github.com/gorilla/mux"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *application) init(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	app.models.Goalies.Init()
	goalies, err := app.models.Goalies.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, goalies, nil)
}

func (app *application) save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	app.models.Goalies.Save()

	w.WriteHeader(http.StatusCreated)
	w.Write(nil)
}

func (app *application) goalieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		goalies, err := app.models.Goalies.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.writeJSON(w, http.StatusOK, goalies, nil)
	}

	if r.Method == http.MethodPost {
		var newGoalie data.Goalie
		err := app.readJSON(w, r, &newGoalie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.models.Goalies.AddGoalie(newGoalie)

		app.writeJSON(w, http.StatusCreated, newGoalie, nil)
	}
}

func (app *application) goalieByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		goalie, err := app.models.Goalies.GetGoalie(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.writeJSON(w, http.StatusOK, goalie, nil)
	}

	if r.Method == http.MethodPut {
		var changedGoalie data.Goalie
		err := app.readJSON(w, r, &changedGoalie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.models.Goalies.UpdateGoalie(changedGoalie)

		goalie, err := app.models.Goalies.GetGoalie(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.writeJSON(w, http.StatusOK, goalie, nil)
	}

	if r.Method == http.MethodDelete {

	}
}

func (app *application) gameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalieId, err := strconv.ParseInt(vars["goalieId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	seasonId, err := strconv.ParseInt(vars["seasonId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		games, err := app.models.Goalies.GetGames(goalieId, seasonId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.writeJSON(w, http.StatusOK, games, nil)
	}

	if r.Method == http.MethodPost {
		var newGame data.Game
		err := app.readJSON(w, r, &newGame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		outGame, err := app.models.Goalies.AddGame(newGame, goalieId, seasonId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		app.writeJSON(w, http.StatusCreated, outGame, nil)
	}
}

func (app *application) gameByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goalieId, err := strconv.ParseInt(vars["goalieId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	seasonId, err := strconv.ParseInt(vars["seasonId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gameId, err := strconv.ParseInt(vars["seasonId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPut {
		var updatedGame data.Game
		err := app.readJSON(w, r, &updatedGame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		game, err := app.models.Goalies.UpdateGame(updatedGame, goalieId, seasonId, gameId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.writeJSON(w, http.StatusCreated, game, nil)
	}
}
