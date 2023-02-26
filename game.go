package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnterMap(newMap Map, players []Player){
	fmt.Printf("EnterMap\n")
	for _, p := range players {
		p:=p //send to pointer receiver
		go p.StartMove()
	}
}

func GameStartImpl(ctx context.Context) {
	fmt.Printf("GameStartImpl\n")
	req, err := ReqContextWithType[ReqContext[GameStartReq]](ctx)
	if err != nil {
		log.Printf("GameType1StartImpl err - %s", err.Error())
	}

	newMap := GenGameMap(req.Request.MapInfo)
	fmt.Printf("Map Gen.\n")
	var players []Player
	for _, pi := range req.Request.GameInfo.Players {
		players = append(players, GenPlayer(pi, req.Request.MapInfo))
	}
	fmt.Printf("Players Gen.\n")
	EnterMap(newMap, players)
	
}

func doGameStart(ctx context.Context) error {
	fmt.Printf("doGameStart\n")
	req, err := ReqContextWithType[ReqContext[GameStartReq]](ctx)
	if err != nil {
		return err
	}
	switch req.Request.GameInfo.Type {
	case 1:
		go GameStartImpl(ctx)
	default:
		return fmt.Errorf("game type %d does not exist", req.Request.GameInfo.Type)
	}
	return nil
}

// @Summary     게임 시작
// @Description 받은 명세서에 알맞게 게임 시작
// @Tags        game
// @Accept      json
// @Produce     json
// @Param       request       body   GameStartReq true "게임 명세"
// @Router      /v1/api/game/start [post]
// @Success      202
func Start(c *gin.Context){
	ctx, err :=GinContext[GameStartReq](c, envConf)
	if err != nil {
		fmt.Printf("err -%s\n", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if err = doGameStart(ctx); err != nil {
		fmt.Printf("err -%s\n", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusAccepted, nil)
}