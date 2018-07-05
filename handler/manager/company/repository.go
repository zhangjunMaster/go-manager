package company

import (
	"database/sql"
	"go-manager/lib"
	"go-manager/model"
	"time"
)

const (
	tableName   = "company"
	timeFormart = "2006-01-02 15:04:05"
)

type Company struct {
	ID                    string    `json:"id"`
	Name                  string    `json:"name"`
	DomainName            string    `json:"domainName"`
	ManagerServer         string    `json:"managerServer"`
	IsPrivateDevelopment  int       `json:"isPrivateDevelopment"`
	Edition               int       `json:"edition"`
	Type                  int       `json:"type"`
	OriginalManagerServer string    `json:"originalManagerServer"`
	IsDeleted             int       `json:"isDeleted"`
	CreateDate            time.Time `json:"createDate"`
	LastUpdate            time.Time `json:"lastUpdate"`
	CreateDateAtHub       time.Time `json:"createDateAtHub"`
}

type User struct {
	Name     string       `json:"name"`
	Age      int          `json:"age"`
	Birthday lib.JsonTime `json:"birthday"`
}

var userModel = model.Model{TableName: "users"}

func CreateUser(v interface{}) (sql.Result, error) {
	var udb = userModel.Open()
	res, err := udb.Create(v)
	return res, err
}
