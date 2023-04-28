package pokemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hydrospirt/GoPokeApi/pkg/common/models"
)

type UpdatePokemonRequestBody struct {
	Name        string `json:"name`
	Description string `json:"description"`
}

func (h handler) UpdatePokemon(c *gin.Context) {
	id := c.Param("id")
	body := UpdatePokemonRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var pokemon models.Pokemon

	if result := h.DB.First(&pokemon, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	pokemon.Name = body.Name
	pokemon.Description = body.Description

	h.DB.Save(&pokemon)

	c.JSON(http.StatusOK, &pokemon)

}
