package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `admin_role_priv` (
//  `id` int(10) NOT NULL AUTO_INCREMENT,
//  `role_id` smallint(3) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
//  `s` char(100) NOT NULL DEFAULT '' COMMENT '模块/控制器/动作',
//  `data` char(50) NOT NULL DEFAULT '' COMMENT '其他参数',
//  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
//  `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
//  `type` char(32) NOT NULL DEFAULT 'url' COMMENT '类别url菜单function独立功能user用户独有',
//  PRIMARY KEY (`id`),
//  KEY `role_id` (`role_id`),
//  KEY `role_id_2` (`role_id`,`s`)
//) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='角色权限表'
const TableAdminRolePriv = "admin_role_priv"
const AdminRolePrivColumns = "`id`,`role_id`,`s`,`data`,`aid`,`menu_id`,`type`"

type AdminRolePrivDB struct {
	Id     int64  `description:"" db:"id" primary:"true"`
	RoleId int    `description:"角色ID" db:"role_id"`
	S      string `description:"模块/控制器/动作" db:"s"`
	Data   string `description:"其他参数" db:"data"`
	Aid    int64  `description:"管理员ID" db:"aid"`
	MenuId int64  `description:"菜单ID" db:"menu_id"`
	Type   string `description:"类别url菜单function独立功能user用户独有" db:"type"`
}
type AdminRolePriv struct {
	Id     int64  `description:"" db:"id" primary:"true"`
	RoleId int    `description:"角色ID" db:"role_id"`
	S      string `description:"模块/控制器/动作" db:"s"`
	Data   string `description:"其他参数" db:"data"`
	Aid    int64  `description:"管理员ID" db:"aid"`
	MenuId int64  `description:"菜单ID" db:"menu_id"`
	Type   string `description:"类别url菜单function独立功能user用户独有" db:"type"`
}

//AdminRolePriv数据转换
func (a *AdminRolePriv) SetMapValue(m map[string]interface{}) {
	a.Id = db.Int64Default(m["id"])
	a.RoleId = db.IntDefault(m["role_id"])
	a.S = db.StringDefault(m["s"])
	a.Data = db.StringDefault(m["data"])
	a.Aid = db.Int64Default(m["aid"])
	a.MenuId = db.Int64Default(m["menu_id"])
	a.Type = db.StringDefault(m["type"])
}

//AdminRolePriv单个查询
func AdminRolePrivGetOne(where string, q db.Query, args ...interface{}) *AdminRolePriv {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminRolePrivColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminRolePriv))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &AdminRolePriv{}
	ret.SetMapValue(m)
	return ret
}

//AdminRolePriv列表查询
func AdminRolePrivGetList(where string, q db.Query, args ...interface{}) []*AdminRolePriv {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminRolePrivColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminRolePriv))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminRolePriv
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminRolePriv{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//AdminRolePriv列表查询
func AdminRolePrivGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*AdminRolePriv {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminRolePrivColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminRolePriv))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminRolePriv
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminRolePriv{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
