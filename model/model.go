package model

// based on official api https://docs.battlesnake.com/api

type Snake struct {
	Id             string        `json:"id"`
	Name           string        `json:"name"`
	Health         int           `json:"health"`
	Body           []Coord       `json:"body"`
	Head           Coord         `json:"head"`
	Length         int           `json:"length"`
	Latency        string        `json:"latency"`
	Shout          string        `json:"shout"`
	Customizations Customization `json:"customizations"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Customization struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

type Board struct {
	Height  int     `json:"height"`
	Width   int     `json:"width"`
	Food    []Coord `json:"food"`
	Hazards []Coord `json:"hazards"`
	Snakes  []Snake `json:"snakes"`
}

type GameReq struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type Game struct {
	Id      string `json:"id"`
	Rules   Rules  `json:"ruleset"`
	Map     string `json:"map"`
	Source  string `json:"source"`
	Timeout int    `json:"timeout"`
}

type Rules struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Settings Settings `json:"settings"`
}

type Settings struct {
	FoodSpawnChance     int `json:"foodSpawnChance"`
	MinimumFood         int `json:"minimumFood"`
	HazardDamagePerTurn int `json:"hazardDamagePerTurn"`
}

type InfoResp struct {
	ApiVersion string `json:"apiversion"`
	Author     string `json:"author,omitempty"`
	Color      string `json:"color,omitempty"`
	Head       string `json:"head,omitempty"`
	Tail       string `json:"tail,omitempty"`
	Version    string `json:"version,omitempty"`
}

type MoveResp struct {
	Move  Move   `json:"move"`
	Shout string `json:"shout,omitempty"`
}
