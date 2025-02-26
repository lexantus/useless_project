package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexantus/useless_project/internal/errs"
	"github.com/lexantus/useless_project/internal/models"
	"net/http"
)

type ActivityHandler struct {
	repo activityRepo
}

func validate(param interface{}) (int, bool) {
	res, isOk := param.(int)
	return res, isOk
}

type activityRepo interface {
	GetActivity(ctx context.Context, timespanMs int) (models.Activity, error)
}

func NewActivityHandler(repo activityRepo) *ActivityHandler {
	return &ActivityHandler{repo: repo}
}

func (s *ActivityHandler) Handler(c *gin.Context) {
	milliseconds, isOk := validate(c.Param("milliseconds"))
	if !isOk {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong param - milliseconds"})
	}

	a, err := s.repo.GetActivity(c, milliseconds)

	if err != nil {
		switch {
		case errors.Is(err, errs.ActivityNotFoundError):
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, a)
}
