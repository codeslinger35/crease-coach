package data

// GoalieModel is an interface for a Goalie datastore
type GoalieModel interface {
	Init() error
	Save() error

	// GetAll will return all data for all goalies in the database
	GetAll() ([]Goalie, error)

	// GetGoalie will return all data for a goalie with the id goalieId
	GetGoalie(int64) (Goalie, error)
	// AddGoalie will add a new goalie to the database
	AddGoalie(Goalie) (Goalie, error)
	// UpdateGoalie will update an existing goalie
	UpdateGoalie(Goalie) (Goalie, error)

	// GetGames will return all games for a goalie with id goalieId and season with id seasonId
	GetGames(int64, int64) ([]Game, error)
	// AddGame will add new game data for goalie goalieId in season seasonId
	AddGame(Game, int64, int64) (Game, error)
	// UpdateGame will update all game data for a goalie goalieId in season seasonId
	UpdateGame(Game, int64, int64, int64) (Game, error)
}

type Goalie struct {
	Id        int64    `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Age       string   `json:"age"`
	Team      string   `json:"team"`
	Seasons   []Season `json:"seasons"`
}

type Season struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Team        string `json:"team"`
	Year        string `json:"year"`
	Games       []Game `json:"games"`
}

type Game struct {
	Id       int64    `json:"id"`
	Date     string   `json:"date"`
	Opponent string   `json:"opponent"`
	Started  bool     `json:"started"`
	Pulled   bool     `json:"pulled"`
	Notes    string   `json:"notes"`
	Periods  []Period `json:"periods"`
}

type Period struct {
	Id             int64           `json:"id"`
	PeriodNumber   int64           `json:"periodNumber"`
	Notes          string          `json:"notes"`
	ShotsAgainst   int64           `json:"shotsAgainst"`
	Saves          int64           `json:"saves"`
	CoachingPoints []CoachingPoint `json:"coachingPoints"`
}

type CoachingPoint struct {
	Id    int64    `json:"id"`
	Time  string   `json:"time"`
	Evemt string   `json:"event"`
	Notes string   `json:"notes"`
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
}

type Tag struct {
	Name     string `json:"name"`
	Positive bool   `json:"positive"`
}
