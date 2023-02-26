package main

import (
	"fmt"
	"gamewatching/docs"
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func router(conf Config) *gin.Engine{
	r := gin.Default()

	v1:= r.Group("/v1")
	{
		api:=v1.Group("/api")
		{
			game:=api.Group("/game")
			{
				game.POST("start", Start)
			}
		}
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", conf.Address, conf.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return r
}

var envConf Config

func main(){
	envConf, err := loadConfig()
	if err != nil {
		log.Panicf("server failed with err : %s", err.Error())
	}
	router(envConf).Run(fmt.Sprintf(":%d", envConf.Port))
}