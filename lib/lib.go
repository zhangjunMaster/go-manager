package lib

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func QuetoKey(v interface{}) []string {
	fileds := []string{}
	rt := reflect.TypeOf(v)
	num := rt.NumField()
	for i := 0; i < num; i++ { //遍历
		f := rt.Field(i)    //reflect.StructField 类型
		fieldName := f.Name //获取struct的key值
		fileds = append(fileds, fmt.Sprintf("%v%v%v", "`", fieldName, "`"))
	}
	fmt.Printf("%+v", fileds)
	return fileds
}

const (
	tableName   = "company"
	timeFormart = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	//
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = JsonTime(now)
	fmt.Println("---------t:", &*t)
	return err
}

func QuetoValue(v interface{}) []interface{} {
	values := []interface{}{}
	rv := reflect.Indirect(reflect.ValueOf(v))
	typ := reflect.TypeOf(v)
	num := typ.NumField()
	for i := 0; i < num; i++ { //遍历
		value := rv.Field(i).Interface()
		// 1.先判断是不是lib.JsonTime类型
		if reflect.TypeOf(value).String() == "lib.JsonTime" {
			// 2. 类型断言 value.(T),因为是interface类型，所以不能time.Time(t).Format
			// value.(JsonTime) 转换成JsonTime 类型
			// 因为time.Time(t) 中t势time.Time，JsonTime又是time.Time，v又是JsonTime，所以可以作为参数传入
			if t, ok := value.(JsonTime); ok {
				// 3. 将时间转成字符串类型
				values = append(values, time.Time(t).Format("2006-01-02 15:04:05"))
			}
		} else {
			values = append(values, rv.Field(i).Interface())

		}
	}
	return values
}

func join(colNames []string) string {
	return strings.Join(colNames, ", ")
}
