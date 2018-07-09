package admin

import "time"

type AdminModel struct {
	ID                 string    `json:"id"`
	Company_id         string    `json:"companyId"`
	Name               string    `json:"name"`
	Login_name         string    `json:"loginName"`
	Password           string    `json:"password"`
	Email              string    `json:"email"`
	Mobile             string    `json:"mobile"`
	Failed_login_times int       `json:"failedLoginTimes"`
	Is_locked          int       `json:"isLocked"`
	Unlock_time        time.Time `json:"unlockTime"`
	Is_deleted         int       `json:"isDeleted"`
	Create_date        time.Time `json:"createDate"`
	Last_update        time.Time `json:"lastUpdate"`
	Is_first_login     int       `json:"isFirstLogin"`
}
