package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `area_ext` (
//  `ext_id` int(11) NOT NULL AUTO_INCREMENT,
//  `id` int(11) DEFAULT '0' COMMENT 'ID',
//  `name` char(50) DEFAULT '' COMMENT '名称',
//  `name_en` varchar(100) DEFAULT '' COMMENT '英文名称',
//  `parent_id` int(11) DEFAULT '0' COMMENT '上级栏目ID',
//  `type` tinyint(4) DEFAULT '0' COMMENT '类别;0默认;1又名;2;3属于;11已合并到;12已更名为',
//  `name_traditional` varchar(50) DEFAULT '' COMMENT '繁体名称',
//  `sort` int(11) DEFAULT '0' COMMENT '排序',
//  `type_name` varchar(50) DEFAULT '' COMMENT '类别名称',
//  `other_name` varchar(50) DEFAULT '' COMMENT '根据类别名称填写',
//  PRIMARY KEY (`ext_id`),
//  KEY `id` (`id`,`parent_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=820345 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='地区扩展表'
const TableAreaExt = "area_ext"
const AreaExtColumns = "`ext_id`,`id`,`name`,`name_en`,`parent_id`,`type`,`name_traditional`,`sort`,`type_name`,`other_name`"

type AreaExtDB struct {
	ExtId           int64   `description:"" db:"ext_id" primary:"true"`
	Id              *int64  `description:"ID" db:"id"`
	Name            *string `description:"名称" db:"name"`
	NameEn          *string `description:"英文名称" db:"name_en"`
	ParentId        *int64  `description:"上级栏目ID" db:"parent_id"`
	Type            *int    `description:"类别;0默认;1又名;2;3属于;11已合并到;12已更名为" db:"type"`
	NameTraditional *string `description:"繁体名称" db:"name_traditional"`
	Sort            *int64  `description:"排序" db:"sort"`
	TypeName        *string `description:"类别名称" db:"type_name"`
	OtherName       *string `description:"根据类别名称填写" db:"other_name"`
}
type AreaExt struct {
	ExtId           int64  `description:"" db:"ext_id" primary:"true"`
	Id              int64  `description:"ID" db:"id"`
	Name            string `description:"名称" db:"name"`
	NameEn          string `description:"英文名称" db:"name_en"`
	ParentId        int64  `description:"上级栏目ID" db:"parent_id"`
	Type            int    `description:"类别;0默认;1又名;2;3属于;11已合并到;12已更名为" db:"type"`
	NameTraditional string `description:"繁体名称" db:"name_traditional"`
	Sort            int64  `description:"排序" db:"sort"`
	TypeName        string `description:"类别名称" db:"type_name"`
	OtherName       string `description:"根据类别名称填写" db:"other_name"`
}

//AreaExt数据转换
func (a *AreaExt) SetMapValue(m map[string]interface{}) {
	a.ExtId = db.Int64Default(m["ext_id"])
	a.Id = db.Int64Default(m["id"])
	a.Name = db.StringDefault(m["name"])
	a.NameEn = db.StringDefault(m["name_en"])
	a.ParentId = db.Int64Default(m["parent_id"])
	a.Type = db.IntDefault(m["type"])
	a.NameTraditional = db.StringDefault(m["name_traditional"])
	a.Sort = db.Int64Default(m["sort"])
	a.TypeName = db.StringDefault(m["type_name"])
	a.OtherName = db.StringDefault(m["other_name"])
}

//AreaExt单个查询
func AreaExtGetOne(where string, q db.Query, args ...interface{}) *AreaExt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AreaExtColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAreaExt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &AreaExt{}
	ret.SetMapValue(m)
	return ret
}

//AreaExt列表查询
func AreaExtGetList(where string, q db.Query, args ...interface{}) []*AreaExt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AreaExtColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAreaExt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AreaExt
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &AreaExt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//AreaExt列表查询
func AreaExtGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*AreaExt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AreaExtColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAreaExt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*AreaExt
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &AreaExt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
