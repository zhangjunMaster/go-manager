package admin

import (
	"go-manager/handler/manager/license"
	"go-manager/lib"
	"strconv"
	"time"
)

func LicenseToEntity(license license.LicenseModel) license.LicenseModel {
	license.Create_date = lib.JsonTime(time.Now())
	license.Last_update = lib.JsonTime(time.Now())
	return license
}

func AdminToModel(source map[string]string) map[string]interface{} {
	target := make(map[string]interface{})
	license := make(map[string]interface{})
	target["adminUserId"] = source["id"]
	target["companyId"] = source["company_id"]
	target["domainName"] = source["domain_name"]
	target["edition"] = source["edition"]
	target["email"] = source["email"]
	target["id"] = source["id"]
	target["loginName"] = source["email"]
	target["mobile"] = source["mobile"]
	target["isPrivateDevelopment"] = source["is_private_development"]
	target["name"] = source["name"]
	target["isFirstLogin"] = false
	target["isLocked"] = 0
	target["is_deleted"] = 0
	license["isValid"] = source["is_valid"]
	license["licenseNumber"] = source["license_number"]
	license["startTime"] = source["start_time"]
	license["expirationTime"] = source["expiration_time"]
	license["edition"], _ = strconv.Atoi(source["edition"])
	license["adminEmail"] = source["email"]
	license["loginName"] = source["email"]
	license["status"] = "true"
	target["license"] = license
	return target
}
