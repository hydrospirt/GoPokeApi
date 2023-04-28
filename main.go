package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hydrospirt/GoPokeApi/pkg/common/db"
	"github.com/hydrospirt/GoPokeApi/pkg/pokemon"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbURL := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbURL)
	pokemon.RegisterRoutes(r, h)
	r.Run(port)
}
