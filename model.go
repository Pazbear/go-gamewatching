package main

type Location struct{
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Stat struct{
	Pow int `json:"pow"`
	Spd int `json:"spd"`
	Int int `json:"int"`
}

type PlayerInfo struct{
	Name string `json:"name"`
	Stat Stat `json:"stat"`
}

type GameInfo struct{
	Type int `json:"type"`
	Players []PlayerInfo `json:"players"`
}

type MapInfo struct {
	Vertical int `json:"vertical"`
	Horizontal int `json:"horizontal"`
	Floor int `json:"floor"`
	LockTime int `json:"locktime"`
}

type Room struct{
	Accessible bool
}

type Floor struct {
	rooms [][]Room
	Accessible bool
}

type Map struct {
	Floor []Floor
}