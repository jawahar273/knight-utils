package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type DefaultQuery struct {
	PageNumber int64
	PerPage    int64
}

func DefaultQueryString(c *gin.Context) DefaultQuery {

	pageNumber, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 16, 64)
	perPage, _ := strconv.ParseInt(c.DefaultQuery("perPage", "20"), 16, 64)

	return DefaultQuery{
		PageNumber: pageNumber,
		PerPage:    perPage,
	}

}
