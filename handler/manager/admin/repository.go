package admin

import (
	"go-manager/lib"
	"go-manager/model"
)

type AdminModel struct {
	ID                 string       `json:"id"`
	Company_id         string       `json:"companyId"`
	Name               string       `json:"adminName"`
	Login_name         string       `json:"loginName"`
	Password           string       `json:"password"`
	Email              string       `json:"adminEmail"`
	Mobile             string       `json:"adminTel"`
	Failed_login_times int          `json:"failedLoginTimes"`
	Is_locked          int          `json:"isLocked"`
	Unlock_time        lib.JsonTime `json:"unlockTime"`
	Is_deleted         int          `json:"isDeleted"`
	Create_date        lib.JsonTime `json:"createDate"`
	Last_update        lib.JsonTime `json:"lastUpdate"`
	Is_first_login     int          `json:"isFirstLogin"`
}

var adminModel = model.Model{TableName: "admin"}

func Get(username string, password string) (map[int]map[string]string, error) {
	var mdb = adminModel.Open()
	params := []interface{}{username, password}
	sqlStr := `
			SELECT a.*, c.is_private_development, c.edition, c.domain_name, l.*
			FROM admin AS a INNER JOIN company AS c ON a.company_id = c.id 
			INNER JOIN license AS l ON a.company_id = l.company_id
			WHERE a.is_locked = 0 AND a.is_deleted = 0
			AND a.login_name = ? 
			AND a.password = ?
		   `
	rows, err := mdb.Query(sqlStr, params)
	return rows, err
}
