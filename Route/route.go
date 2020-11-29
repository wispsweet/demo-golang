package Route

import (
	"Question/Controller"
	"github.com/gin-gonic/gin"
)

var controller Controller.Controller

func InitRoute() *gin.Engine {
	router := gin.Default()

	questionService := router.Group("/questionservice")
	{
		//分页试题列表
		questionService.GET("/getpagequestion", Controller.QuestionPageList)
		//试题列表
		questionService.GET("/getQuestionList", controller.GetAll)
	}

	return router
}
