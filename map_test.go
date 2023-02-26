package main

import "testing"

func TestGenGameMap(t *testing.T) {
	var mapInfo MapInfo = MapInfo{Vertical: 2, Horizontal: 3, Floor: 4}
	newMap := GenGameMap(mapInfo)

	if len(newMap.Floor) != mapInfo.Floor {
		t.Logf("expected map floor length is %d, but %d", mapInfo.Floor, len(newMap.Floor))
		t.Fail()
	}else if len(newMap.Floor[0].rooms) != mapInfo.Vertical {
		t.Logf("expected map room vertical length is %d, but %d", mapInfo.Vertical,len(newMap.Floor[0].rooms))
		t.Fail()
	}else if len(newMap.Floor) != mapInfo.Horizontal {
		t.Logf("expected map room horizontal length is %d, but %d", mapInfo.Horizontal, len(newMap.Floor[0].rooms[0]))
		t.Fail()
	}
}