package Model

import (
	"Question/Dao"
	"Question/Text"
	"Question/Util"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func QuestionCondition(ctx *gin.Context) (string,[]interface{}) {
	var str string
	var params []interface{}
	//id
	questionIdParam := ctx.Query("question_id")
	questionIdRes, questionIdType, questionIdSplit := Util.StringSplitToInt(questionIdParam, ",")
	if questionIdRes {
	   if questionIdType {
	   		str += "id = ? and "
		   	params = append(params, questionIdSplit.(int))
	   } else {
	   		questionId := questionIdSplit.([]int)
	   		str += "id in (" + Util.PrepareIntSymbol(questionId) +") and "
	   		params = Util.IntSliceMerge(params, questionId)
	   }
	}
	//subject_id
	subjectId, _ := Util.StrToInt(ctx.Query("subject_id"), 0)
	if subjectId > 0 {
		str += "subject_id = ? and "
		params = append(params, subjectId)
	}
	//knowledge_id
	knowledgeId, _ := Util.StrToInt(ctx.Query("knowledge_id"), 0)
	if knowledgeId > 0 {
		str += "knowledge_id = ? and "
		params = append(params, knowledgeId)
	}
	//question_type_id
	questionTypeId, _ := Util.StrToInt(ctx.Query("question_type_id"), 0)
	if questionTypeId > 0 {
		str += "question_type_id = ? and "
		params = append(params, questionTypeId)
	}
	//is_positive_question_type
	positiveType := ctx.Query("is_positive_question_type")
	if positiveType == "1" {
		str += "question_type_id >= 0 and "
	}
	//difficulty
	difficulty, _ := Util.StrToInt(ctx.Query("difficulty"), 0) // ---------------------------
	if Util.IsContain(Text.QuestionDifficulty(), difficulty) {
		str += "difficulty = ? and "
		params = append(params, difficulty)
	}
	//province|数据库查询 获取省份id
	province := ctx.Query("province")
	if len(province) > 0 {
		provinceId := Dao.DistrictIdByName(province)
		if provinceId > 0 {
			str += "province_id = ? and "
			params = append(params, provinceId)
		}
	}
	//city
	city := ctx.Query("city")
	if len(city) > 0 {
		str += "city = ? and "
		params = append(params, city)
	}
	//year
	year, _ := Util.StrToInt(ctx.Query("year"), 2000)
	if year > 2000 {
		yearFlag := Text.YearFlag()
	   if len(strconv.Itoa(year)) == len(strconv.Itoa(yearFlag)) {
		   enableYear := yearFlag - year
		   if Util.IsEnableYear(enableYear) {
			   str += "year >= ? and "
			   params = append(params, enableYear)
		   }
	   } else if Util.IsEnableYear(year) {
		   str += "year = ? and "
		   params = append(params, year)
	   }
	}
	//source
	source := ctx.Query("source")
	if len(source) > 0 && subjectId > 0 {
		str += "source = ? and "
		params = append(params, source)
	}
	//source_id
	sourceId, _ := Util.StrToInt(ctx.Query("source_id"), 0)
	if sourceId > 0 && subjectId > 0 {
		str += "source_id = ? and "
		params = append(params, sourceId)
	}
	//keyword
	keyword := ctx.Query("keyword")
	if len(keyword) > 0 {
		str += "content_text like %?% and "
		params = append(params, keyword)
	}
	//course_type_id
	courseTypeId, _ := Util.StrToInt(ctx.Query("course_type_id"))
	if courseTypeId > 0 {
		str += "course_type_id = ? and "
		params = append(params, courseTypeId)
	}
	//question_model
	questionModel := ctx.Query("question_model")
	modelRes, modelType, modelSplit := Util.StringSplitToInt(questionModel, ",")
	if modelRes {
	   if modelType {
		   str += "question_model = ? and "
		   params = append(params, modelSplit.(int))
	   } else {
		   str += "question_model in (" + Util.PrepareIntSymbol(modelSplit.([]int)) + ") and "
		   params = Util.IntSliceMerge(params, modelSplit.([]int))
	   }
	} else {
		str += "question_model = 0 and "
	}
	//province_id
	provinceId, _ := Util.StrToInt(ctx.Query("province_id"))
	if provinceId > 0 {
		str += "province_id = ? and "
		params = append(params, provinceId)
	}
	//exclude_id
	excludeId := ctx.Query("exclude_id")
	excludeRes, excludeType, excludeSplit := Util.StringSplitToInt(excludeId, ",")
	if excludeRes {
		if excludeType {
			str += "id != ? and "
			params = append(params, excludeSplit.(int))
		} else {
			str += "province_id not in (" + Util.PrepareIntSymbol(excludeSplit.([]int)) + ") and "
			params = Util.IntSliceMerge(params, excludeSplit.([]int))
		}
	}
	//paper_id
	paperId := ctx.Query("paper_id")
	paperRes, paperType, paperSplit := Util.StringSplitToInt(paperId, ",")
	if paperRes {
		if paperType {
			str += "paper_id = ? and "
			params = append(params, paperSplit.(int))
		} else {
			str += "id in (" + Util.PrepareIntSymbol(paperSplit.([]int)) + ") and "
			params = Util.IntSliceMerge(params, paperSplit.([]int))
		}
	}
	//question_level
	level := ctx.Query("question_level")
	levelRes, levelType, levelSplit := Util.StringSplitToInt(level, ",")
	if levelRes {
		if levelType {
			str += "question_level = ? and "
			params = append(params, levelSplit.(int))
		} else {
			str += "question_level >= ? and "
			params = append(params, levelSplit.([]int))
		}
	}

	return strings.TrimRight(str, " and "), params
}