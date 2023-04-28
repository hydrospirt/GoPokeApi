package pokemon

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := r.Group("/pokemon")
	routes.POST("/", h.AddPokemon)
	routes.GET("/", h.GetAllPokemon)
	routes.GET("/:id", h.GetPokemon)
	routes.PUT("/:id", h.UpdatePokemon)
	routes.DELETE("/:id", h.DeletePokemon)
}
