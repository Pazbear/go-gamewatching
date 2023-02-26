package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MoveTicker struct {
	C chan time.Time
	stopc chan struct{}

}

type Player struct {
	Loc Location
	PlayerInfo
	IsAlive bool
	MoveTicker
}



func NewMoveTicker() *MoveTicker {
	mt := &MoveTicker{
		C:     make(chan time.Time),
		stopc: make(chan struct{}),
	}
	return mt
}

func (mt *MoveTicker) nextInterval(spd int) time.Duration{
	interval := rand.Intn(10-spd)+1
	return time.Duration(interval) * time.Second
}

func (p *Player) collocate(mi MapInfo){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randLoc := Location{
		X: r.Intn(mi.Horizontal-1),
		Y: r.Intn(mi.Vertical-1),
		Z: r.Intn(mi.Floor-1),
	}
	p.Loc = randLoc
}

//살아있는 동안 spd에 따른 tick마다 액션 취함
func (p *Player) StartMove(){
	t := time.NewTimer(p.MoveTicker.nextInterval(p.PlayerInfo.Stat.Spd))
	for p.IsAlive{
		select {
		case <-p.MoveTicker.stopc:
			t.Stop()
			return
		case <-t.C:
			t.Stop()

			//Player Action
			fmt.Printf("player %s get action", p.Name)

			t = time.NewTimer(p.MoveTicker.nextInterval(p.PlayerInfo.Stat.Spd))
		}
	}
}

func (p *Player) StopMove() {
	close(p.MoveTicker.stopc)
	close(p.MoveTicker.C)
}

func GenPlayer(pi PlayerInfo, mi MapInfo) Player{
	var newPlayer Player
	newPlayer.PlayerInfo = pi
	newPlayer.IsAlive = true
	newPlayer.MoveTicker = *NewMoveTicker()
	newPlayer.collocate(mi)
	return newPlayer
}