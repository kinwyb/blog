package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `area` (
//  `id` int(11) NOT NULL AUTO_INCREMENT,
//  `name` char(50) DEFAULT '' COMMENT '名称',
//  `name_en` varchar(100) DEFAULT '' COMMENT '英文名称',
//  `parent_id` int(11) DEFAULT '0' COMMENT '上级栏目ID',
//  `type` tinyint(4) DEFAULT '0' COMMENT '类别;0默认;',
//  `name_traditional` varchar(50) DEFAULT '' COMMENT '繁体名称',
//  `sort` int(11) DEFAULT '0' COMMENT '排序',
//  PRIMARY KEY (`id`),
//  KEY `parent_id` (`parent_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=659004503 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='地区表'
const TableArea = "area"
const AreaColumns = "`id`,`name`,`name_en`,`parent_id`,`type`,`name_traditional`,`sort`"

type AreaDB struct {
	Id              int64   `description:"" db:"id" primary:"true"`
	Name            *string `description:"名称" db:"name"`
	NameEn          *string `description:"英文名称" db:"name_en"`
	ParentId        *int64  `description:"上级栏目ID" db:"parent_id"`
	Type            *int    `description:"类别;0默认;" db:"type"`
	NameTraditional *string `description:"繁体名称" db:"name_traditional"`
	Sort            *int64  `description:"排序" db:"sort"`
}
type Area struct {
	Id              int64  `description:"" db:"id" primary:"true"`
	Name            string `description:"名称" db:"name"`
	NameEn          string `description:"英文名称" db:"name_en"`
	ParentId        int64  `description:"上级栏目ID" db:"parent_id"`
	Type            int    `description:"类别;0默认;" db:"type"`
	NameTraditional string `description:"繁体名称" db:"name_traditional"`
	Sort            int64  `description:"排序" db:"sort"`
}

//Area数据转换
func (a *Area) SetMapValue(m map[string]interface{}) {
	a.Id = db.Int64Default(m["id"])
	a.Name = db.StringDefault(m["name"])
	a.NameEn = db.StringDefault(m["name_en"])
	a.ParentId = db.Int64Default(m["parent_id"])
	a.Type = db.IntDefault(m["type"])
	a.NameTraditional = db.StringDefault(m["name_traditional"])
	a.Sort = db.Int64Default(m["sort"])
}

//Area单个查询
func AreaGetOne(where string, q db.Query, args ...interface{}) *Area {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AreaColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableArea))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Area{}
	ret.SetMapValue(m)
	return ret
}

//Area列表查询
func AreaGetList(where string, q db.Query, args ...interface{}) []*Area {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AreaColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableArea))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Area
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Area{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Area列表查询
func AreaGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Area {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AreaColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableArea))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Area
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Area{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
