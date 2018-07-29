package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	tableName   = "company"
	timeFormart = "2006-01-02 15:04:05"
	salt        = "zhangjun"
)

type JsonTime time.Time

func QuoteKey(v interface{}) []string {
	fileds := []string{}
	rt := reflect.TypeOf(v)
	num := rt.NumField()
	for i := 0; i < num; i++ { //遍历
		f := rt.Field(i)    //reflect.StructField 类型
		fieldName := f.Name //获取struct的key值
		fileds = append(fileds, fmt.Sprintf("%v%v%v", "`", fieldName, "`"))
	}
	return fileds
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	// 将字符串转为Time类型
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	// 将Time强转成 JsonTime 类型
	*t = JsonTime(now)
	return err
}

func QuoteValue(v interface{}) []interface{} {
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

// quote colNames, placeholders, colValues
func Quote(v interface{}) ([]string, []string, []interface{}) {
	var placeholders []string
	colNames := QuoteKey(v)
	colValues := QuoteValue(v)
	for _, _ = range colNames {
		placeholders = append(placeholders, "?")
	}
	return colNames, placeholders, colValues
}

func join(colNames []string) string {
	return strings.Join(colNames, ", ")
}

// md5加盐
func Md5Salt(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	h.Write([]byte(salt))
	st := h.Sum(nil)
	// 16进制转成字符串
	return hex.EncodeToString(st)
}

func DecodeBase64(encodeString string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		return "", err
	}
	decodeStrig := string(decodeBytes)
	return decodeStrig, nil
}

func EncodeBase64(s string) (string, error) {
	input := []byte(s)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString, nil
}
