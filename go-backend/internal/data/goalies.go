package data

type GoalieModel interface {
	Init() error
	Save() error

	GetAll() ([]Goalie, error)
	GetGoalie(int64) (Goalie, error)
	AddGoalie(Goalie) (Goalie, error)
	UpdateGoalie(Goalie) (Goalie, error)
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
	Note           string          `json:"note"`
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
