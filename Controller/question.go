package Controller

import (
	"Question/Dao"
	"Question/Mapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

var link = Dao.DataBase{}
var mapper Mapper.TodosMapper
var scanColumns = []interface{}{&mapper.Id, &mapper.Title, &mapper.Status}

func (c *Controller) GetAll(ctx *gin.Context) {
	result := map[string]interface{}{
		"code":    200,
		"message": "",
		"data":    []map[string]interface{}{},
	}

	link := link.Connect()
	rows, err := link.Query("select * from todos") //使用*查询全部 则需要使用Scan绑定全部字段
	defer link.Close()
	if err != nil {
		panic("Query error...")
	}

	i := 0
	var data = make([]interface{}, 20)

	for rows.Next() {
		err := rows.Scan(scanColumns...) //使用Scan绑定字段
		if err != nil {
			continue
		}
		data[i] = map[string]interface{}{
			"id":     mapper.Id,
			"title":  mapper.Title,
			"status": mapper.Status,
		}
		i++
	}
	rows.Close()

	data = data[:i]
	result["data"] = data

	ctx.JSON(http.StatusOK, result)
}
