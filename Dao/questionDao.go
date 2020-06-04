package Dao

import (
	"database/sql"
	"fmt"
)

//查询一条数据
func QuestionRow(where string, params []interface{}) map[string]interface{} {
	var result map[string]interface{}
	var row *sql.Row

	link := Connect()
	defer link.Close()
	sqlStr := "select id from vip_question where " + where + " limit 1"

	fmt.Println(sqlStr)
	fmt.Println(params)

	if len(params) > 0 {
		row = link.QueryRow(sqlStr, params...)
	} else {
		row = link.QueryRow(sqlStr)
	}

	var id int
	err := row.Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		result = map[string]interface{}{
			"error": 0,
		}
	} else {
		result = map[string]interface{}{
			"id": id,
		}
	}
	return result
}

func QuestionList(where string, params []interface{}) []interface{} {
	var result = make([]interface{}, 20)
	var rows *sql.Rows
	var err error

	link := Connect()
	defer link.Close()
	sqlStr := "select id from vip_question where " + where

	fmt.Println(sqlStr)
	fmt.Println(params)

	if len(params) > 0 {
		rows, err = link.Query(sqlStr, params...)
	} else {
		rows, err = link.Query(sqlStr)
	}

	if err != nil {
		result = []interface{}{}
	} else {
		var id int;
		i := 0
		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				result = []interface{}{}
			}
			result[i] = map[string]interface{}{
				"id": id,
			}
			i++
		}
		result = result[:i]
	}
	return result
}
