package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `type` (
//  `id` int(11) NOT NULL AUTO_INCREMENT,
//  `name` char(100) NOT NULL DEFAULT '' COMMENT '名称',
//  `code` char(32) NOT NULL DEFAULT '' COMMENT '代码',
//  `mark` char(32) NOT NULL DEFAULT '' COMMENT '标志',
//  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属类别ID',
//  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级ID、属于/上级ID',
//  `value` int(10) NOT NULL DEFAULT '0' COMMENT '值',
//  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '内容',
//  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除0否1是',
//  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
//  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
//  `aid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加人',
//  `module` char(50) NOT NULL DEFAULT '' COMMENT '模块',
//  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
//  `setting` varchar(255) DEFAULT NULL COMMENT '扩展参数',
//  `is_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子类1是0否',
//  `is_system` tinyint(1) NOT NULL DEFAULT '0' COMMENT '系统参数禁止修改',
//  `is_show` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示在配置页面上',
//  PRIMARY KEY (`id`),
//  KEY `is_del` (`is_del`),
//  KEY `parent_id` (`parent_id`),
//  KEY `sort` (`sort`),
//  KEY `mark` (`mark`),
//  KEY `type_id` (`type_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='类别'
const TableType = "type"
const TypeColumns = "`id`,`name`,`code`,`mark`,`type_id`,`parent_id`,`value`,`content`,`is_del`,`sort`,`remark`,`time_add`,`aid`,`module`,`is_default`,`setting`,`is_child`,`is_system`,`is_show`"

type TypeDB struct {
	Id        int64   `description:"" db:"id" primary:"true"`
	Name      string  `description:"名称" db:"name"`
	Code      string  `description:"代码" db:"code"`
	Mark      string  `description:"标志" db:"mark"`
	TypeId    int64   `description:"所属类别ID" db:"type_id"`
	ParentId  int64   `description:"上级ID、属于/上级ID" db:"parent_id"`
	Value     int64   `description:"值" db:"value"`
	Content   string  `description:"内容" db:"content"`
	IsDel     int64   `description:"是否删除0否1是" db:"is_del"`
	Sort      int64   `description:"排序" db:"sort"`
	Remark    *string `description:"备注" db:"remark"`
	TimeAdd   *string `description:"添加时间" db:"time_add"`
	Aid       int64   `description:"添加人" db:"aid"`
	Module    string  `description:"模块" db:"module"`
	IsDefault int     `description:"是否默认" db:"is_default"`
	Setting   *string `description:"扩展参数" db:"setting"`
	IsChild   int     `description:"是否有子类1是0否" db:"is_child"`
	IsSystem  int     `description:"系统参数禁止修改" db:"is_system"`
	IsShow    int     `description:"是否显示在配置页面上" db:"is_show"`
}
type Type struct {
	Id        int64  `description:"" db:"id" primary:"true"`
	Name      string `description:"名称" db:"name"`
	Code      string `description:"代码" db:"code"`
	Mark      string `description:"标志" db:"mark"`
	TypeId    int64  `description:"所属类别ID" db:"type_id"`
	ParentId  int64  `description:"上级ID、属于/上级ID" db:"parent_id"`
	Value     int64  `description:"值" db:"value"`
	Content   string `description:"内容" db:"content"`
	IsDel     int64  `description:"是否删除0否1是" db:"is_del"`
	Sort      int64  `description:"排序" db:"sort"`
	Remark    string `description:"备注" db:"remark"`
	TimeAdd   string `description:"添加时间" db:"time_add"`
	Aid       int64  `description:"添加人" db:"aid"`
	Module    string `description:"模块" db:"module"`
	IsDefault int    `description:"是否默认" db:"is_default"`
	Setting   string `description:"扩展参数" db:"setting"`
	IsChild   int    `description:"是否有子类1是0否" db:"is_child"`
	IsSystem  int    `description:"系统参数禁止修改" db:"is_system"`
	IsShow    int    `description:"是否显示在配置页面上" db:"is_show"`
}

//Type数据转换
func (t *Type) SetMapValue(m map[string]interface{}) {
	t.Id = db.Int64Default(m["id"])
	t.Name = db.StringDefault(m["name"])
	t.Code = db.StringDefault(m["code"])
	t.Mark = db.StringDefault(m["mark"])
	t.TypeId = db.Int64Default(m["type_id"])
	t.ParentId = db.Int64Default(m["parent_id"])
	t.Value = db.Int64Default(m["value"])
	t.Content = db.StringDefault(m["content"])
	t.IsDel = db.Int64Default(m["is_del"])
	t.Sort = db.Int64Default(m["sort"])
	t.Remark = db.StringDefault(m["remark"])
	t.TimeAdd = db.StringDefault(m["time_add"])
	t.Aid = db.Int64Default(m["aid"])
	t.Module = db.StringDefault(m["module"])
	t.IsDefault = db.IntDefault(m["is_default"])
	t.Setting = db.StringDefault(m["setting"])
	t.IsChild = db.IntDefault(m["is_child"])
	t.IsSystem = db.IntDefault(m["is_system"])
	t.IsShow = db.IntDefault(m["is_show"])
}

//Type单个查询
func TypeGetOne(where string, q db.Query, args ...interface{}) *Type {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TypeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableType))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Type{}
	ret.SetMapValue(m)
	return ret
}

//Type列表查询
func TypeGetList(where string, q db.Query, args ...interface{}) []*Type {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TypeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableType))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Type
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Type{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Type列表查询
func TypeGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Type {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TypeColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableType))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Type
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Type{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
