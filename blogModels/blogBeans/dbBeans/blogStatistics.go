package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `blog_statistics` (
//  `statistics_id` int(11) NOT NULL AUTO_INCREMENT,
//  `blog_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
//  `comment` int(11) NOT NULL DEFAULT '0' COMMENT '评论人数',
//  `read` int(11) NOT NULL DEFAULT '0' COMMENT '阅读人数',
//  `seo_title` varchar(255) NOT NULL DEFAULT '' COMMENT 'SEO标题',
//  `seo_description` varchar(255) NOT NULL DEFAULT '' COMMENT 'SEO摘要',
//  `seo_keyword` varchar(255) NOT NULL DEFAULT '' COMMENT 'SEO关键词',
//  PRIMARY KEY (`statistics_id`),
//  KEY `blog_id` (`blog_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=132 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='博客统计'
const TableBlogStatistics = "blog_statistics"
const BlogStatisticsColumns = "`statistics_id`,`blog_id`,`comment`,`read`,`seo_title`,`seo_description`,`seo_keyword`"

type BlogStatisticsDB struct {
	StatisticsId   int64  `description:"" db:"statistics_id" primary:"true"`
	BlogId         int64  `description:"文章ID" db:"blog_id"`
	Comment        int64  `description:"评论人数" db:"comment"`
	Read           int64  `description:"阅读人数" db:"read"`
	SeoTitle       string `description:"SEO标题" db:"seo_title"`
	SeoDescription string `description:"SEO摘要" db:"seo_description"`
	SeoKeyword     string `description:"SEO关键词" db:"seo_keyword"`
}
type BlogStatistics struct {
	StatisticsId   int64  `description:"" db:"statistics_id" primary:"true"`
	BlogId         int64  `description:"文章ID" db:"blog_id"`
	Comment        int64  `description:"评论人数" db:"comment"`
	Read           int64  `description:"阅读人数" db:"read"`
	SeoTitle       string `description:"SEO标题" db:"seo_title"`
	SeoDescription string `description:"SEO摘要" db:"seo_description"`
	SeoKeyword     string `description:"SEO关键词" db:"seo_keyword"`
}

//BlogStatistics数据转换
func (b *BlogStatistics) SetMapValue(m map[string]interface{}) {
	b.StatisticsId = db.Int64Default(m["statistics_id"])
	b.BlogId = db.Int64Default(m["blog_id"])
	b.Comment = db.Int64Default(m["comment"])
	b.Read = db.Int64Default(m["read"])
	b.SeoTitle = db.StringDefault(m["seo_title"])
	b.SeoDescription = db.StringDefault(m["seo_description"])
	b.SeoKeyword = db.StringDefault(m["seo_keyword"])
}

//BlogStatistics单个查询
func BlogStatisticsGetOne(where string, q db.Query, args ...interface{}) *BlogStatistics {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogStatisticsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogStatistics))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &BlogStatistics{}
	ret.SetMapValue(m)
	return ret
}

//BlogStatistics列表查询
func BlogStatisticsGetList(where string, q db.Query, args ...interface{}) []*BlogStatistics {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogStatisticsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogStatistics))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogStatistics
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogStatistics{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//BlogStatistics列表查询
func BlogStatisticsGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*BlogStatistics {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogStatisticsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogStatistics))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogStatistics
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogStatistics{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
