package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `blog_tag` (
//  `tag_id` int(11) NOT NULL AUTO_INCREMENT,
//  `name` char(100) NOT NULL DEFAULT '' COMMENT '名称',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
//  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
//  `blog_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
//  PRIMARY KEY (`tag_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=167 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='博客标签'
const TableBlogTag = "blog_tag"
const BlogTagColumns = "`tag_id`,`name`,`time_add`,`aid`,`blog_id`"

type BlogTagDB struct {
	TagId   int64   `description:"" db:"tag_id" primary:"true"`
	Name    string  `description:"名称" db:"name"`
	TimeAdd *string `description:"添加时间" db:"time_add"`
	Aid     int64   `description:"管理员ID" db:"aid"`
	BlogId  int64   `description:"文章ID" db:"blog_id"`
}
type BlogTag struct {
	TagId   int64  `description:"" db:"tag_id" primary:"true"`
	Name    string `description:"名称" db:"name"`
	TimeAdd string `description:"添加时间" db:"time_add"`
	Aid     int64  `description:"管理员ID" db:"aid"`
	BlogId  int64  `description:"文章ID" db:"blog_id"`
}

//BlogTag数据转换
func (b *BlogTag) SetMapValue(m map[string]interface{}) {
	b.TagId = db.Int64Default(m["tag_id"])
	b.Name = db.StringDefault(m["name"])
	b.TimeAdd = db.StringDefault(m["time_add"])
	b.Aid = db.Int64Default(m["aid"])
	b.BlogId = db.Int64Default(m["blog_id"])
}

//BlogTag单个查询
func BlogTagGetOne(where string, q db.Query, args ...interface{}) *BlogTag {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogTagColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogTag))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &BlogTag{}
	ret.SetMapValue(m)
	return ret
}

//BlogTag列表查询
func BlogTagGetList(where string, q db.Query, args ...interface{}) []*BlogTag {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogTagColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogTag))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogTag
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogTag{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//BlogTag列表查询
func BlogTagGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*BlogTag {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogTagColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogTag))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogTag
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogTag{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
