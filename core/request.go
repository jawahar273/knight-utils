package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type DefaultQuery struct {
	PageNumber int32
	PerPage    int32
}

func DefaultQueryString(c *gin.Context) DefaultQuery {

	pageNumber, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	perPage, _ := strconv.ParseInt(c.DefaultQuery("perPage", "20"), 10, 32)
	// limiting to max of 50 at a time
	if perPage > 50 {
		perPage = 50
	}

	return DefaultQuery{
		PageNumber: int32(pageNumber),
		PerPage:    int32(perPage),
	}

}
