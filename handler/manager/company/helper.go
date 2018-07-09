package company

import (
	"fmt"
	"go-manager/handler/manager/admin"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CompanyToEntity(company CompanyModel) CompanyModel {
	if company.ID == "" {
		uid := uuid.Must(uuid.NewV4())
		fmt.Println("uid:", fmt.Sprintf("%s", uid))
		company.ID = fmt.Sprintf("%s", uid)
	}
	company.Create_date = time.Now()
	company.Last_update = time.Now()
	company.Create_date_at_hub = time.Now()
	return company
}

func AdminToEntity(admin admin.AdminModel) admin.AdminModel {
	if admin.ID == "" {
		uid := uuid.Must(uuid.NewV4())
		fmt.Println("uid:", fmt.Sprintf("%s", uid))
		admin.ID = fmt.Sprintf("%s", uid)
	}
	return admin
}
