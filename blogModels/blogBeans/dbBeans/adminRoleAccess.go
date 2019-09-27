package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `admin_role_access` (
//  `aid` int(11) DEFAULT '0' COMMENT '管理员ID',
//  `role_id` int(11) DEFAULT '0' COMMENT '角色ID',
//  UNIQUE KEY `aid_role_id` (`aid`,`role_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员与角色的关系、一个管理员可以有多个角色'
const TableAdminRoleAccess = "admin_role_access"
const AdminRoleAccessColumns = "`aid`,`role_id`"

type AdminRoleAccessDB struct {
	Aid    *int64 `description:"管理员ID" db:"aid"`
	RoleId *int64 `description:"角色ID" db:"role_id"`
}
type AdminRoleAccess struct {
	Aid    int64 `description:"管理员ID" db:"aid"`
	RoleId int64 `description:"角色ID" db:"role_id"`
}

//AdminRoleAccess数据转换
func (a *AdminRoleAccess) SetMapValue(m map[string]interface{}) {
	a.Aid = db.Int64Default(m["aid"])
	a.RoleId = db.Int64Default(m["role_id"])
}

//AdminRoleAccess单个查询
func AdminRoleAccessGetOne(where string, q db.Query, args ...interface{}) *AdminRoleAccess {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminRoleAccessColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminRoleAccess))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &AdminRoleAccess{}
	ret.SetMapValue(m)
	return ret
}

//AdminRoleAccess列表查询
func AdminRoleAccessGetList(where string, q db.Query, args ...interface{}) []*AdminRoleAccess {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminRoleAccessColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminRoleAccess))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminRoleAccess
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminRoleAccess{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//AdminRoleAccess列表查询
func AdminRoleAccessGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*AdminRoleAccess {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminRoleAccessColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminRoleAccess))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminRoleAccess
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminRoleAccess{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
