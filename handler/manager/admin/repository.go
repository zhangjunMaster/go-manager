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
			SELECT t1.*, t2.is_private_development, t2.edition, t2.domain_name FROM admin AS t1
			INNER JOIN company AS t2
			WHERE t1.is_locked = 0 AND t1.is_deleted = 0
			AND t1.login_name = ? 
			AND t1.password = ?
		   `
	rows, err := mdb.Query(sqlStr, params)
	return rows, err
}
