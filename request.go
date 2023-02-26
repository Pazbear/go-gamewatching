package main

type GameStartReq struct {
	GameInfo GameInfo `json:"game_info"`
	MapInfo MapInfo `json:"map_info"`
	Timeout int `json:"timeout"`
}