package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `blog_sync_queue` (
//  `queue_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `blog_id` int(11) NOT NULL DEFAULT '0' COMMENT '本站博客id',
//  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类型',
//  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0:待运行 10:失败 99:成功',
//  `time_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
//  `msg` varchar(255) NOT NULL DEFAULT '' COMMENT '内容',
//  `map_id` int(11) NOT NULL DEFAULT '0' COMMENT '同步ID',
//  PRIMARY KEY (`queue_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='博客同步队列'
const TableBlogSyncQueue = "blog_sync_queue"
const BlogSyncQueueColumns = "`queue_id`,`blog_id`,`type_id`,`status`,`time_update`,`time_add`,`msg`,`map_id`"

type BlogSyncQueueDB struct {
	QueueId    int64   `description:"" db:"queue_id" primary:"true"`
	BlogId     int64   `description:"本站博客id" db:"blog_id"`
	TypeId     int64   `description:"类型" db:"type_id"`
	Status     int     `description:"状态：0:待运行 10:失败 99:成功" db:"status"`
	TimeUpdate *string `description:"最后一次更新时间" db:"time_update"`
	TimeAdd    *string `description:"插入时间" db:"time_add"`
	Msg        string  `description:"内容" db:"msg"`
	MapId      int64   `description:"同步ID" db:"map_id"`
}
type BlogSyncQueue struct {
	QueueId    int64  `description:"" db:"queue_id" primary:"true"`
	BlogId     int64  `description:"本站博客id" db:"blog_id"`
	TypeId     int64  `description:"类型" db:"type_id"`
	Status     int    `description:"状态：0:待运行 10:失败 99:成功" db:"status"`
	TimeUpdate string `description:"最后一次更新时间" db:"time_update"`
	TimeAdd    string `description:"插入时间" db:"time_add"`
	Msg        string `description:"内容" db:"msg"`
	MapId      int64  `description:"同步ID" db:"map_id"`
}

//BlogSyncQueue数据转换
func (b *BlogSyncQueue) SetMapValue(m map[string]interface{}) {
	b.QueueId = db.Int64Default(m["queue_id"])
	b.BlogId = db.Int64Default(m["blog_id"])
	b.TypeId = db.Int64Default(m["type_id"])
	b.Status = db.IntDefault(m["status"])
	b.TimeUpdate = db.StringDefault(m["time_update"])
	b.TimeAdd = db.StringDefault(m["time_add"])
	b.Msg = db.StringDefault(m["msg"])
	b.MapId = db.Int64Default(m["map_id"])
}

//BlogSyncQueue单个查询
func BlogSyncQueueGetOne(where string, q db.Query, args ...interface{}) *BlogSyncQueue {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogSyncQueueColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogSyncQueue))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &BlogSyncQueue{}
	ret.SetMapValue(m)
	return ret
}

//BlogSyncQueue列表查询
func BlogSyncQueueGetList(where string, q db.Query, args ...interface{}) []*BlogSyncQueue {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogSyncQueueColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogSyncQueue))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogSyncQueue
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogSyncQueue{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//BlogSyncQueue列表查询
func BlogSyncQueueGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*BlogSyncQueue {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogSyncQueueColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlogSyncQueue))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*BlogSyncQueue
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &BlogSyncQueue{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
