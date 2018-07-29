package admin

import (
	"go-manager/handler/manager/license"
	"go-manager/lib"
	"time"
)

func LicenseToEntity(license license.LicenseModel) license.LicenseModel {
	license.Create_date = lib.JsonTime(time.Now())
	license.Last_update = lib.JsonTime(time.Now())
	return license
}

func AdminToModel(source map[string]string) map[string]string {
	target := make(map[string]string)
	target["companyId"] = source["company_id"]
	target["domainName"] = source["domain_name"]
	target["edition"] = source["edition"]
	target["email"] = source["email"]
	target["failedLoginTimes"] = source["failed_login_times"]
	target["id"] = source["id"]
	target["mobile"] = source["mobile"]
	target["isFirstLogin"] = source["is_first_login"]
	target["isPrivateDevelopment"] = source["is_private_development"]
	target["name"] = source["name"]

	return target
}
