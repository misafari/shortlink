package http

import (
	"github.com/gin-gonic/gin"
	"ir.safari.shortlink/repository"
	"net/http"
)

type Controller struct {
	originalUrlRepository repository.OriginalUrlRepository
}

func NewController(originalUrlRepository repository.OriginalUrlRepository) *Controller {
	return &Controller{
		originalUrlRepository: originalUrlRepository,
	}
}

func (h *Controller) RedirectHandler(c *gin.Context) {
	code := c.Params.ByName("code")
	if code == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	result, fetchErr := h.originalUrlRepository.FetchUrl(code)
	if fetchErr != nil || result == nil {
		//todo log
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if result.OriginalUrl == "" {
		//todo log
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Redirect(301, result.OriginalUrl)
}
