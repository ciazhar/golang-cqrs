package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func RequestParamInt(ctx *gin.Context, paramName string) (int, error) {
	response, err := strconv.Atoi(ctx.Query(paramName))
	if err != nil {
		return 0, err
	}
	return response, nil
}

func RequestParamFloat(ctx *gin.Context, paramName string) (float64, error) {
	response := ctx.Query(paramName)
	responseF, err := strconv.ParseFloat(response, 64)
	if err != nil {
		return 0, err
	}

	return responseF, nil
}

func RequestParamString(ctx *gin.Context, paramName string) string {
	return ctx.Query(paramName)
}

func RequestBody(ctx *gin.Context, v interface{}) error {
	return ctx.BindJSON(v)
}

func RequestHeader(ctx *gin.Context, name string) string {
	return ""
}

func RequestPathVariableInteger(ctx *gin.Context, name string) int {
	id := ctx.Param(name)
	value, _ := strconv.Atoi(id)
	return value
}

func RequestPathVariableString(ctx *gin.Context, name string) string {
	return ctx.Param(name)
}

type GinParam struct {
	Gin   *gin.Context
	Param Param
}

func NewParamGin(c *gin.Context) GinParam {
	return GinParam{
		Gin:   c,
		Param: NewParam(),
	}
}

func (q *GinParam) Include() {
	value := q.Gin.Query("include")
	if value != "" {
		q.Param.Var["include"] = value
	}
}

func (q *GinParam) Paging() {
	offset, _ := RequestParamInt(q.Gin, "offset")
	limit, _ := RequestParamInt(q.Gin, "limit")
	paginate := RequestParamString(q.Gin, "paginate")
	if paginate == "false" {
		offset = 0
		limit = 0
	}
	q.Param.Offset = offset
	q.Param.Limit = limit
}

func (q *GinParam) Add(fieldName string, value interface{}) {
	q.Param.Query[fieldName] = value
}
