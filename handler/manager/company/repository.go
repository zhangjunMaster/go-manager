package company

import (
	"database/sql"
	"fmt"
	"go-manager/handler/manager/admin"
	"go-manager/lib"
	"go-manager/model"
	"strings"
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

// 将客户端传来的数据转换成sql的字段
type CompanyModel struct {
	ID                      string    `json:"id"`
	Name                    string    `json:"name"`
	Domain_name             string    `json:"domainName"`
	Manager_server          string    `json:"managerServer"`
	Is_private_development  int       `json:"isPrivateDevelopment"`
	Edition                 int       `json:"edition"`
	Type                    int       `json:"type"`
	Original_manager_server string    `json:"originalManagerServer"`
	Is_deleted              int       `json:"isDeleted"`
	Create_date             time.Time `json:"createDate"`
	Last_update             time.Time `json:"lastUpdate"`
	Create_date_at_hub      time.Time `json:"createDateAtHub"`
}

var companyModel = model.Model{TableName: "company"}

func CreateCompany(companyData CompanyModel, adminData admin.AdminModel) (sql.Result, error) {
	var mdb = companyModel.Open()
	var transactions []model.Transaction
	companyColNames, companyPlaceholders, companyColValues := lib.Quote(companyData)
	adminColNames, adminPlaceholders, adminColValues := lib.Quote(adminData)
	companySqlStr := fmt.Sprintf("INSERT INTO `company` (%v) VALUES (%v)",
		strings.Join(companyColNames, ", "),
		strings.Join(companyPlaceholders, ", "))
	adminSqlStr := fmt.Sprintf("INSERT INTO `admin` (%v) VALUES (%v)",
		strings.Join(adminColNames, ", "),
		strings.Join(adminPlaceholders, ", "))
	companyTransaction := model.Transaction{Sql: companySqlStr, Values: companyColValues}
	adminTransaction := model.Transaction{Sql: adminSqlStr, Values: adminColValues}
	transactions = append(transactions, companyTransaction, adminTransaction)
	err := mdb.Transaction(transactions)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
