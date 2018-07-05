package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

const (
	tableName   = "company"
	timeFormart = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	fmt.Println("---------t:", now)
	*t = JsonTime(now)
	fmt.Println("---------t:", JsonTime(now))
	fmt.Println("---------t:", reflect.TypeOf(time.Time(*t).Format("2006-01-02 15:04:05")).String())
	return err
}

type User struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Birthday JsonTime `json:"birthday"`
}

func main() {
	src := `{"Name":"5", "Age":12,"Birthday":"2016-06-30 16:09:51"}`
	p := new(User)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", time.Time(p.Birthday))
}
