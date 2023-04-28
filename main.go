package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hydrospirt/GoPokeApi/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbURL := viper.Get("DB_URL").(string)

	r := gin.Default()
	db.Init(dbURL)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"port":  port,
			"dbUrl": dbURL,
		})
	})
	r.Run(port)
}
