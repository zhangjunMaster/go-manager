package user

import (
	"fmt"
	"go-manager/lib"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func UserToEntity(user *UserModel) *UserModel {
	if (*user).ID == "" {
		uid := uuid.Must(uuid.NewV4())
		//fmt.Println("uid:", fmt.Sprintf("%s", uid))
		(*user).ID = fmt.Sprintf("%s", uid)
	}
	(*user).Create_date = lib.JsonTime(time.Now())
	(*user).Last_update = lib.JsonTime(time.Now())
	return user
}
func UserToModel(source map[string]string) map[string]interface{} {
	target := make(map[string]interface{})
	target["companyId"] = source["company_id"]
	target["departmentId"] = source["department_id"]
	target["id"] = source["id"]
	target["email"] = source["email"]
	target["mobile"] = source["mobile"]
	target["failedLoginTimes"] = source["failed_login_times"]
	target["fullDepartmentId"] = source["full_department_id"]
	target["fullDepartmentName"] = source["full_department_name"]
	target["lastLoginTime"] = source["last_login_time"]
	target["lastUpdate"] = source["last_update"]
	target["loginName"] = source["login_name"]
	target["name"] = source["name"]
	target["source"] = source["source"]
	target["isActivated"], _ = strconv.Atoi(source["is_activated"])
	target["isDeleted"], _ = strconv.Atoi(source["is_deleted"])
	target["isDisabled"], _ = strconv.Atoi(source["is_disabled"])
	return target
}
