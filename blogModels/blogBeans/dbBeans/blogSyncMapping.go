package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `blog_sync_mapping` (
//  `map_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `blog_id` int(11) NOT NULL DEFAULT '0' COMMENT '本站blog的id',
//  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id',
//  `id` varchar(64) NOT NULL DEFAULT '' COMMENT 'csdn的id',
//  `time_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
//  `mark` char(32) NOT NULL DEFAULT '' COMMENT '标志',
//  `is_sync` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否同步过',
//  `extend` varchar(5000) DEFAULT NULL COMMENT '扩展参数',
//  PRIMARY KEY (`map_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='本站blog_id 与其他同步站点的id关系'
const TableBlogSyncMapping = "blog_sync_mapping"
const BlogSyncMappingColumns = "`map_id`,`blog_id`,`type_id`,`id`,`time_update`,`time_add`,`mark`,`is_sync`,`extend`"

type BlogSyncMappingDB struct {
	MapId      int64   `description:"" db:"map_id" primary:"true"`
	BlogId     int64   `description:"本站blog的id" db:"blog_id"`
	TypeId     int64   `description:"类别id" db:"type_id"`
	Id         string  `description:"csdn的id" db:"id"`
	TimeUpdate *string `description:"最后一次更新时间" db:"time_update"`
	TimeAdd    *string `description:"插入时间" db:"time_add"`
	Mark       string  `description:"标志" db:"mark"`
	IsSync     int     `description:"是否同步过" db:"is_sync"`
	Extend     *string `description:"扩展参数" db:"extend"`
}
type BlogSyncMapping struct {
	MapId      int64  `description:"" db:"map_id" primary:"true"`
	BlogId     int64  `description:"本站blog的id" db:"blog_id"`
	TypeId     int64  `description:"类别id" db:"type_id"`
	Id         string `description:"csdn的id" db:"id"`
	TimeUpdate string `description:"最后一次更新时间" db:"time_update"`
	TimeAdd    string `description:"插入时间" db:"time_add"`
	Mark       string `description:"标志" db:"mark"`
	IsSync     int    `description:"是否同步过" db:"is_sync"`
	Extend     string `description:"扩展参数" db:"extend"`
}

//BlogSyncMapping数据转换
func (b *BlogSyncMapping) SetMapValue(m map[string]interface{}) {
	b.MapId = db.Int64Default(m["map_id"])
	b.BlogId = db.Int64Default(m["blog_id"])
	b.TypeId = db.Int64Default(m["type_id"])
	b.Id = db.StringDefault(m["id"])
	b.TimeUpdate = db.StringDefault(m["time_update"])
	b.TimeAdd = db.StringDefault(m["time_add"])
	b.Mark = db.StringDefault(m["mark"])
	b.IsSync = db.IntDefault(m["is_sync"])
	b.Extend = db.StringDefault(m["extend"])
}

//BlogSyncMapping单个查询
func BlogSyncMappingGetOne(where string, q db.Query, args ...interface{}) *BlogSyncMapping {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogSyncMappingColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogSyncMapping))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &BlogSyncMapping{}
	ret.SetMapValue(m)
	return ret
}

//BlogSyncMapping列表查询
func BlogSyncMappingGetList(where string, q db.Query, args ...interface{}) []*BlogSyncMapping {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogSyncMappingColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogSyncMapping))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogSyncMapping
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogSyncMapping{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//BlogSyncMapping列表查询
func BlogSyncMappingGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*BlogSyncMapping {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogSyncMappingColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogSyncMapping))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogSyncMapping
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogSyncMapping{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
