package pokemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hydrospirt/GoPokeApi/pkg/common/models"
)

type AddPokemonRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h handler) AddPokemon(c *gin.Context) {
	body := AddPokemonRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pokemon models.Pokemon

	pokemon.Name = body.Name
	pokemon.Description = body.Description

	if result := h.DB.Create(&pokemon); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pokemon)
}
