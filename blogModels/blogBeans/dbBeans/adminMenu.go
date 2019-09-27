package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `admin_menu` (
//  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `name` char(100) NOT NULL DEFAULT '' COMMENT '名称',
//  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级菜单',
//  `s` char(60) NOT NULL DEFAULT '' COMMENT '模块/控制器/动作',
//  `data` char(100) NOT NULL DEFAULT '' COMMENT '其他参数',
//  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
//  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
//  `type` char(32) NOT NULL DEFAULT 'url' COMMENT '类别url菜单function独立功能user用户独有',
//  `level` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '级别',
//  `level1_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '1级栏目ID',
//  `md5` char(32) NOT NULL DEFAULT '' COMMENT 's的md5值',
//  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '显示隐藏;1显示;0隐藏',
//  `is_unique` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户独有此功能1是0否',
//  PRIMARY KEY (`id`),
//  KEY `sort` (`sort`),
//  KEY `parent_id` (`parent_id`),
//  KEY `s` (`s`)
//) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='菜单'
const TableAdminMenu = "admin_menu"
const AdminMenuColumns = "`id`,`name`,`parent_id`,`s`,`data`,`sort`,`remark`,`type`,`level`,`level1_id`,`md5`,`is_show`,`is_unique`"

type AdminMenuDB struct {
	Id       int64  `description:"" db:"id" primary:"true"`
	Name     string `description:"名称" db:"name"`
	ParentId int64  `description:"上级菜单" db:"parent_id"`
	S        string `description:"模块/控制器/动作" db:"s"`
	Data     string `description:"其他参数" db:"data"`
	Sort     int64  `description:"排序" db:"sort"`
	Remark   string `description:"备注" db:"remark"`
	Type     string `description:"类别url菜单function独立功能user用户独有" db:"type"`
	Level    int    `description:"级别" db:"level"`
	Level1Id int64  `description:"1级栏目ID" db:"level1_id"`
	Md5      string `description:"s的md5值" db:"md5"`
	IsShow   int    `description:"显示隐藏;1显示;0隐藏" db:"is_show"`
	IsUnique int    `description:"用户独有此功能1是0否" db:"is_unique"`
}
type AdminMenu struct {
	Id       int64  `description:"" db:"id" primary:"true"`
	Name     string `description:"名称" db:"name"`
	ParentId int64  `description:"上级菜单" db:"parent_id"`
	S        string `description:"模块/控制器/动作" db:"s"`
	Data     string `description:"其他参数" db:"data"`
	Sort     int64  `description:"排序" db:"sort"`
	Remark   string `description:"备注" db:"remark"`
	Type     string `description:"类别url菜单function独立功能user用户独有" db:"type"`
	Level    int    `description:"级别" db:"level"`
	Level1Id int64  `description:"1级栏目ID" db:"level1_id"`
	Md5      string `description:"s的md5值" db:"md5"`
	IsShow   int    `description:"显示隐藏;1显示;0隐藏" db:"is_show"`
	IsUnique int    `description:"用户独有此功能1是0否" db:"is_unique"`
}

//AdminMenu数据转换
func (a *AdminMenu) SetMapValue(m map[string]interface{}) {
	a.Id = db.Int64Default(m["id"])
	a.Name = db.StringDefault(m["name"])
	a.ParentId = db.Int64Default(m["parent_id"])
	a.S = db.StringDefault(m["s"])
	a.Data = db.StringDefault(m["data"])
	a.Sort = db.Int64Default(m["sort"])
	a.Remark = db.StringDefault(m["remark"])
	a.Type = db.StringDefault(m["type"])
	a.Level = db.IntDefault(m["level"])
	a.Level1Id = db.Int64Default(m["level1_id"])
	a.Md5 = db.StringDefault(m["md5"])
	a.IsShow = db.IntDefault(m["is_show"])
	a.IsUnique = db.IntDefault(m["is_unique"])
}

//AdminMenu单个查询
func AdminMenuGetOne(where string, q db.Query, args ...interface{}) *AdminMenu {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminMenuColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminMenu))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &AdminMenu{}
	ret.SetMapValue(m)
	return ret
}

//AdminMenu列表查询
func AdminMenuGetList(where string, q db.Query, args ...interface{}) []*AdminMenu {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminMenuColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminMenu))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminMenu
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminMenu{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//AdminMenu列表查询
func AdminMenuGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*AdminMenu {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminMenuColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdminMenu))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AdminMenu
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &AdminMenu{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
