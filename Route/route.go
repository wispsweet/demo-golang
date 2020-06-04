package Route

import (
	"Question/Controller"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine  {
	router := gin.Default()

	questionService := router.Group("/questionservice")
	{
		//分页试题列表
		questionService.GET("/getpagequestion", Controller.QuestionPageList)
		//试题列表
		questionService.GET("/getQuestionList", Controller.QuestionList)
	}

	return router
}