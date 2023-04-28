package pokemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hydrospirt/GoPokeApi/pkg/common/models"
)

func (h handler) DeletePokemon(c *gin.Context) {
	id := c.Param("id")

	var pokemon models.Pokemon

	if result := h.DB.First(&pokemon, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&pokemon)

	c.Status(http.StatusOK)
}