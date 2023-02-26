package main

func genGameRoom() Room {
	return Room{Accessible: true}
}

func GenGameMap(mi MapInfo) Map {
	var gameMap Map
	for f:=0; f<mi.Floor; f++ {
		var floor Floor
		for v:=0; v<mi.Vertical; v++ {
			var horizontalRooms []Room
			for h:=0; h<mi.Horizontal; h++ {
				horizontalRooms = append(horizontalRooms, genGameRoom())
			}
			floor.rooms = append(floor.rooms, horizontalRooms)
		}
		floor.Accessible = true
		gameMap.Floor = append(gameMap.Floor, floor)
	}
	return gameMap
}