package dashbord

import (
	"fmt"
	"go-manager/model"
)

const (
	tableName   = "company"
	timeFormart = "2006-01-02 15:04:05"
)

var companyModel = model.Model{TableName: "company"}

func GetActivatedUserCountData(id string) (map[int]map[string]string, error) {
	var mdb = companyModel.Open()
	//设置成interface{}格式是因为slice可以展开，但又不确定这个slice每个数据的类型
	params := []interface{}{id, id, id}
	sqlStr := `
			select total_user.total, activated_user.activated, today_user.today from
			(select count(*) as total, 1 as id from user where company_id = ?) total_user
			inner join
			(select count(*) as activated, 1 as id from user
				where company_id = ?
				and is_activated = true) activated_user
			on total_user.id = activated_user.id
			inner join
			(select count(*) as today, 1 as id from user
				where company_id = ?
				and is_activated = true
				and to_days(activated_time) = to_days(now())) today_user
			on total_user.id = today_user.id;
			`
	fmt.Println(sqlStr, params)
	rows, err := mdb.Query(sqlStr, params)
	return rows, err
}
