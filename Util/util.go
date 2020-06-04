package Util

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

//字符串转数字
func StrToInt(str string, defaultInt ...int) (int,error) {
	result, err := strconv.Atoi(str)
	if err == nil {
		return result, nil
	} else if len(defaultInt) > 0 && err != nil {
		return defaultInt[0], nil
	} else {
		return 0, errors.New("转换失败")
	}
}

//字符串数组转Int
func StringSliceToInt(slice []string) []int  {
	var result []int
	for _, value := range slice {
		strInt, err := StrToInt(value)
		if err != nil {
		   continue
		}
		result = append(result, strInt)
	}
	return result
}

//字符串数组转Int
func StringMapToInt(strmap map[interface{}]string) map[interface{}]int {
	var result = make(map[interface{}]int)
	for index, value := range strmap {
		strInt, err := StrToInt(value)
		if err == nil {
			continue
		}
		result[index] = strInt
	}
	return result
}

//分隔符分隔的数字字符串转为数组
//第一个参数：处理结果是否成功
//第二个参数：返回值类型，true:int | false:[]int
//第三个参数：返回值
func StringSplitToInt(str string, symbol string) (bool,bool,interface{}){
	if len(str) == 0 || len(symbol) == 0 {
	   return false, false, []int{}
	}
	if strings.Contains(str, symbol) {
		strSlice := strings.Split(str, symbol)
		strArr := StringSliceToInt(strSlice)
		if len(strSlice) == len(strArr) {
			if len(strArr) == 1 {
				return true, true, strArr[0]
			} else {
				return true, false, strArr
			}
		}
	} else {
		strInt, _ := StrToInt(str, 0)
		if strInt > 0 {
			return true, true, strInt
		}
	}
	return false, false, []int{}
}

//SQL预处理的占位符
func PrepareIntSymbol(arr []int) string {
	count := reflect.ValueOf(arr).Len()
	if count == 0 {
	    return "";
	}
	var str string
	for j := 0; j < count; j++ {
		str += "?,"
	}
	return strings.TrimRight(str, ",")
}

//SQL预处理的占位符
func PrepareStringSymbol(arr []string) string {
	count := reflect.ValueOf(arr).Len()
	if count == 0 {
		return "";
	}
	var str string
	for j := 0; j < count; j++ {
		str += "?,"
	}
	return strings.TrimRight(str, ",")
}

//数组合并
func StringSliceMerge(slice1 []interface{}, slice2 []string) []interface{} {
	for j := 0; j < len(slice2); j++ {
		if IsContain(slice1, slice2[j]) == false {
			slice1 = append(slice1, slice2[j])
		}
	}
	return slice1
}

//数组合并
func IntSliceMerge(slice1 []interface{}, slice2 []int) []interface{} {
	for j := 0; j < len(slice2); j++ {
		if IsContain(slice1, slice2[j]) == false {
			slice1 = append(slice1, slice2[j])
		}
	}
	return slice1
}




//================================================ IS ==========================================

//map,slice,array中查找
//haystack为map slice array
//needle为要查找的内容
func IsContain(haystack interface{}, needle interface{}) bool {
	targetValue := reflect.ValueOf(haystack)
	switch reflect.TypeOf(haystack).Kind() {
		case reflect.Slice:
		case reflect.Array:
			for i := 0; i < targetValue.Len(); i++ {
				if targetValue.Index(i).Interface() == needle {
					return true
				}
			}
		case reflect.Map:
			if targetValue.MapIndex(reflect.ValueOf(needle)).IsValid() {
				return true
			}
	}
	return false
}

//int类型的in_array
func IsInIntSlice(haystack []int, needle int) bool {
	for _, value := range haystack {
		if value == needle {
		    return true;
		}
	}
	return false;
}

//string类型的in_array
func IsInStringSlice(haystack []string, needle string) bool {
	if sort.SearchStrings(haystack, needle) < len(haystack) {
		return true
	}
	return false
}

//有效年份
func IsEnableYear(year int) bool{
	nowYear := time.Now().Year()
	if year > 2000 && year < nowYear + 2 {
		return true
	}
	return false
}
