package admin

import "go-manager/lib"

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
