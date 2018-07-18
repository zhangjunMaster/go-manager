package company

import (
	"database/sql"
	"fmt"
	"go-manager/handler/manager/admin"
	"go-manager/handler/manager/license"
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
	ID                      string       `json:"id"`
	Name                    string       `json:"name"`
	Domain_name             string       `json:"domainName"`
	Manager_server          string       `json:"managerServer"`
	Is_private_development  int          `json:"isPrivateDevelopment"`
	Edition                 int          `json:"edition"`
	Type                    int          `json:"type"`
	Original_manager_server string       `json:"originalManagerServer"`
	Is_deleted              int          `json:"isDeleted"`
	Create_date             lib.JsonTime `json:"createDate"`
	Last_update             lib.JsonTime `json:"lastUpdate"`
	Create_date_at_hub      lib.JsonTime `json:"createDateAtHub"`
}

var companyModel = model.Model{TableName: "company"}

func CreateCompany(companyData CompanyModel, adminData admin.AdminModel, licenseData license.LicenseModel) (sql.Result, error) {
	var mdb = companyModel.Open()
	var transactions []model.Transaction
	companyColNames, companyPlaceholders, companyColValues := lib.Quote(companyData)
	adminColNames, adminPlaceholders, adminColValues := lib.Quote(adminData)
	licenseColNames, licensePlaceholders, licenseColValues := lib.Quote(licenseData)
	companySqlStr := fmt.Sprintf("INSERT INTO `company` (%v) VALUES (%v)",
		strings.Join(companyColNames, ", "),
		strings.Join(companyPlaceholders, ", "))
	adminSqlStr := fmt.Sprintf("INSERT INTO `admin` (%v) VALUES (%v)",
		strings.Join(adminColNames, ", "),
		strings.Join(adminPlaceholders, ", "))
	licenseSqlStr := fmt.Sprintf("INSERT INTO `license` (%v) VALUES (%v)",
		strings.Join(licenseColNames, ", "),
		strings.Join(licensePlaceholders, ", "))
	companyTransaction := model.Transaction{Sql: companySqlStr, Values: companyColValues}
	adminTransaction := model.Transaction{Sql: adminSqlStr, Values: adminColValues}
	licenseTransaction := model.Transaction{Sql: licenseSqlStr, Values: licenseColValues}
	transactions = append(transactions, companyTransaction, adminTransaction, licenseTransaction)
	err := mdb.Transaction(transactions)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}

func UpdateCompany(companyData CompanyModel, adminData admin.AdminModel, licenseData license.LicenseModel) (sql.Result, error) {
	var mdb = companyModel.Open()
	var transactions []model.Transaction
	companyColNames, companyColValues := mdb.QuoteUpdateFields(companyData)
	adminColNames, adminColValues := mdb.QuoteUpdateFields(adminData)
	licenseColNames, licenseColValues := mdb.QuoteUpdateFields(licenseData)
	companySqlStr := fmt.Sprintf("UPDATE `company` SET %v", companyColNames)
	adminSqlStr := fmt.Sprintf("UPDATE `admin` SET %v", adminColNames)
	licenseSqlStr := fmt.Sprintf("UPDATE `license` SET %v", licenseColNames)
	if len(companyColValues) > 0 {
		companyTransaction := model.Transaction{Sql: companySqlStr, Values: companyColValues}
		transactions = append(transactions, companyTransaction)
	}
	if len(adminColValues) > 0 {
		adminTransaction := model.Transaction{Sql: adminSqlStr, Values: adminColValues}
		transactions = append(transactions, adminTransaction)
	}
	if len(licenseColValues) > 0 {
		licenseTransaction := model.Transaction{Sql: licenseSqlStr, Values: licenseColValues}
		transactions = append(transactions, licenseTransaction)
	}
	err := mdb.Transaction(transactions)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}

func GetDeploymentOfcomapny(domainName string) (map[int]map[string]string, error) {
	var mdb = companyModel.Open()
	params := []interface{}{domainName}
	sqlStr := `
		  SELECT id, name, domain_name AS domainName, manager_server AS managerServer,
		  is_private_development AS isPrivateDevelopment,edition,type,original_manager_server AS originalManagerServer
		  FROM company 
		  	   WHERE domain_name = ?
		  `
	rows, err := mdb.Query(sqlStr, params)
	return rows, err
}
