package license

import "go-manager/lib"

type LicenseModel struct {
	Company_id      string       `json:"companyId"`
	Content         string       `json:"content"`
	License_number  int          `json:"licenceNum"`
	Is_valid        int          `json:"isValid"`
	Start_time      lib.JsonTime `json:"validTime"`
	Expiration_time lib.JsonTime `json:"invalidTime"`
	Create_date     lib.JsonTime `json:"createDate"`
	Last_update     lib.JsonTime `json:"lastUpdate"`
}
