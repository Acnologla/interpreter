package interpreter

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"strings"
)

var defaultVars = map[string]interface{}{
	"false":             false,
	"null":              nil,
	"true":              true,
	"getContext": func() interface{} {
		return context.Background()
	},
	"toUpper": func(value interface{}) interface{}{
		str, ok := value.(string)
		if !ok{
			return ""
		}
		return strings.ToUpper(str)
	},
	"toLower": func(value interface{}) interface{}{
		str, ok := value.(string)
		if !ok{
			return ""
		}
		return strings.ToLower(str)
	},
	"toInt": func (value interface{}) interface{} {
		str, ok := value.(string)
		if !ok {
			return "NaN"
		}
		val, err := strconv.Atoi(str)
		if err == nil{
			return val
		}
		return "NaN"
    },
	"print": func(values ...interface{}) interface{} {
		fmt.Println(values...)
		return nil
	},
	"sleep": func(values interface{}) interface{} {
		time.Sleep(time.Duration(values.(float64)) * time.Second)
		return nil
	},
	"len": func(values interface{}) interface{} {
		arr, ok := toArrInterface(values)
		if !ok {
			return float64(0)
		}
		return float64(len(arr))
	},
	"append": func(values interface{}, item interface{}) interface{} {
		arr, ok := toArrInterface(values)
		if !ok {
			return nil
		}
		return append(arr, item)
	},
}

func Change(varname string, value interface{}){
	defaultVars[varname] = value
}