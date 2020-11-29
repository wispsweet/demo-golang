package Controller

import (
	"Question/Dao"
	"Question/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//分页试题列表
func QuestionPageList(ctx *gin.Context) {
	result := map[string]interface{}{
		"code":      0,
		"page":      1,
		"page_size": 20,
		"count":     0,
		"data":      []interface{}{},
	}
	prepare, params := Model.QuestionCondition(ctx)
	if len(prepare) == 0 {
		ctx.JSON(http.StatusOK, result)
	}

	result["data"] = Dao.QuestionList(prepare, params)
	result["code"] = 200
	ctx.JSON(http.StatusOK, result)
}
