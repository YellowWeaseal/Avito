package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	userCtx    = "userId"
	segmentCtx = "segmentId"
)

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getSegmentId(c *gin.Context) (int, error) {
	id, ok := c.Get(segmentCtx)
	if !ok {
		return 0, errors.New("segment id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("segment id is of invalid type")
	}

	return idInt, nil
}
