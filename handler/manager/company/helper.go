package company

import (
	"fmt"
	"go-manager/handler/manager/admin"
	"go-manager/handler/manager/department"
	"go-manager/handler/manager/license"
	"go-manager/lib"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CompanyToEntity(company CompanyModel) CompanyModel {
	if company.ID == "" {
		uid := uuid.Must(uuid.NewV4())
		//fmt.Println("uid:", fmt.Sprintf("%s", uid))
		company.ID = fmt.Sprintf("%s", uid)
	}
	company.Create_date = lib.JsonTime(time.Now())
	company.Last_update = lib.JsonTime(time.Now())
	company.Create_date_at_hub = lib.JsonTime(time.Now())
	return company
}

func AdminToEntity(admin admin.AdminModel) admin.AdminModel {
	if admin.ID == "" {
		uid := uuid.Must(uuid.NewV4())
		//fmt.Println("uid:", fmt.Sprintf("%s", uid))
		admin.ID = fmt.Sprintf("%s", uid)
	}
	admin.Unlock_time = lib.JsonTime(time.Now())
	admin.Create_date = lib.JsonTime(time.Now())
	admin.Last_update = lib.JsonTime(time.Now())
	admin.Login_name = admin.Email
	if admin.Password == "" {
		admin.Password = lib.Md5Salt("12345678")
	} else {
		admin.Password = lib.Md5Salt(admin.Password)
	}
	//fmt.Println("-----admin-----")
	//fmt.Printf("%+v", admin)
	return admin
}

func LicenseToEntity(license license.LicenseModel) license.LicenseModel {
	license.Create_date = lib.JsonTime(time.Now())
	license.Last_update = lib.JsonTime(time.Now())
	return license
}

func TopdepartmentToEntity(company CompanyModel) department.DepartmentModel {
	uid := uuid.Must(uuid.NewV4())
	d := department.DepartmentModel{
		ID:                       fmt.Sprintf("%s", uid),
		Company_id:               company.ID,
		Name:                     company.Name,
		Parent_department_id:     "0",
		Source:                   1,
		Is_im_group_created:      1,
		Full_department_id:       "0/" + fmt.Sprintf("%s", uid),
		User_device_num:          3,
		Is_limit_user_device_num: 1,
		Create_date:              lib.JsonTime(time.Now()),
		Last_update:              lib.JsonTime(time.Now()),
	}
	return d
}
func CompanyToModel(map[string]string) {

}
