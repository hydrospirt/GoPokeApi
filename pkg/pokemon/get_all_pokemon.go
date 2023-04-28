package pokemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hydrospirt/GoPokeApi/pkg/common/models"
)

func (h handler) GetAllPokemon(c *gin.Context) {
	var allpokemon []models.Pokemon

	if result := h.DB.Find(&allpokemon); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &allpokemon)
}
