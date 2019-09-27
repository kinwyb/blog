package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `tag` (
//  `tag_id` int(11) NOT NULL AUTO_INCREMENT,
//  `name` char(50) NOT NULL DEFAULT '' COMMENT '名称',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
//  PRIMARY KEY (`tag_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='标签'
const TableTag = "tag"
const TagColumns = "`tag_id`,`name`,`time_add`"

type TagDB struct {
	TagId   int64   `description:"" db:"tag_id" primary:"true"`
	Name    string  `description:"名称" db:"name"`
	TimeAdd *string `description:"添加时间" db:"time_add"`
}
type Tag struct {
	TagId   int64  `description:"" db:"tag_id" primary:"true"`
	Name    string `description:"名称" db:"name"`
	TimeAdd string `description:"添加时间" db:"time_add"`
}

//Tag数据转换
func (t *Tag) SetMapValue(m map[string]interface{}) {
	t.TagId = db.Int64Default(m["tag_id"])
	t.Name = db.StringDefault(m["name"])
	t.TimeAdd = db.StringDefault(m["time_add"])
}

//Tag单个查询
func TagGetOne(where string, q db.Query, args ...interface{}) *Tag {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TagColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableTag))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Tag{}
	ret.SetMapValue(m)
	return ret
}

//Tag列表查询
func TagGetList(where string, q db.Query, args ...interface{}) []*Tag {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TagColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableTag))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Tag
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Tag{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Tag列表查询
func TagGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Tag {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TagColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableTag))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Tag
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Tag{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
