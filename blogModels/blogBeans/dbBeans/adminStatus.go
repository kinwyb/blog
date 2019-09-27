package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `admin_status` (
//  `status_id` int(11) NOT NULL AUTO_INCREMENT,
//  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员iD',
//  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
//  `login_ip` char(20) NOT NULL DEFAULT '' COMMENT 'IP',
//  `login` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
//  `aid_add` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
//  `aid_update` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
//  `time_update` timestamp NULL DEFAULT NULL COMMENT '更新时间',
//  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
//  PRIMARY KEY (`status_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='状态'
const TableAdminStatus = "admin_status"
const AdminStatusColumns = "`status_id`,`aid`,`login_time`,`login_ip`,`login`,`aid_add`,`aid_update`,`time_update`,`remark`"

type AdminStatusDB struct {
	StatusId   int64   `description:"" db:"status_id" primary:"true"`
	Aid        int64   `description:"管理员iD" db:"aid"`
	LoginTime  *string `description:"登录时间" db:"login_time"`
	LoginIp    string  `description:"IP" db:"login_ip"`
	Login      int64   `description:"登录次数" db:"login"`
	AidAdd     int64   `description:"添加人" db:"aid_add"`
	AidUpdate  int64   `description:"更新人" db:"aid_update"`
	TimeUpdate *string `description:"更新时间" db:"time_update"`
	Remark     string  `description:"备注" db:"remark"`
}
type AdminStatus struct {
	StatusId   int64  `description:"" db:"status_id" primary:"true"`
	Aid        int64  `description:"管理员iD" db:"aid"`
	LoginTime  string `description:"登录时间" db:"login_time"`
	LoginIp    string `description:"IP" db:"login_ip"`
	Login      int64  `description:"登录次数" db:"login"`
	AidAdd     int64  `description:"添加人" db:"aid_add"`
	AidUpdate  int64  `description:"更新人" db:"aid_update"`
	TimeUpdate string `description:"更新时间" db:"time_update"`
	Remark     string `description:"备注" db:"remark"`
}

//AdminStatus数据转换
func (a *AdminStatus) SetMapValue(m map[string]interface{}) {
	a.StatusId = db.Int64Default(m["status_id"])
	a.Aid = db.Int64Default(m["aid"])
	a.LoginTime = db.StringDefault(m["login_time"])
	a.LoginIp = db.StringDefault(m["login_ip"])
	a.Login = db.Int64Default(m["login"])
	a.AidAdd = db.Int64Default(m["aid_add"])
	a.AidUpdate = db.Int64Default(m["aid_update"])
	a.TimeUpdate = db.StringDefault(m["time_update"])
	a.Remark = db.StringDefault(m["remark"])
}

//AdminStatus单个查询
func AdminStatusGetOne(where string, q db.Query, args ...interface{}) *AdminStatus {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminStatusColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminStatus))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &AdminStatus{}
	ret.SetMapValue(m)
	return ret
}

//AdminStatus列表查询
func AdminStatusGetList(where string, q db.Query, args ...interface{}) []*AdminStatus {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminStatusColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminStatus))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminStatus
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminStatus{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//AdminStatus列表查询
func AdminStatusGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*AdminStatus {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminStatusColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminStatus))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminStatus
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminStatus{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
