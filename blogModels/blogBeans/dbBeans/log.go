package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `log` (
//  `log_id` int(11) NOT NULL AUTO_INCREMENT,
//  `id` int(11) NOT NULL DEFAULT '0' COMMENT 'id',
//  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
//  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//  `mark` char(32) NOT NULL DEFAULT '' COMMENT '标志自定义标志',
//  `data` text COMMENT '其他内容',
//  `no` char(50) NOT NULL DEFAULT '' COMMENT '单号',
//  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录方式;302前台还是后台301',
//  `type_client` int(11) NOT NULL DEFAULT '0' COMMENT '登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他',
//  `ip` char(20) NOT NULL DEFAULT '' COMMENT 'IP',
//  `msg` varchar(255) DEFAULT NULL COMMENT '自定义说明',
//  PRIMARY KEY (`log_id`),
//  KEY `type_login` (`type_login`),
//  KEY `type_client` (`type_client`),
//  KEY `uid` (`uid`),
//  KEY `aid` (`aid`),
//  KEY `id` (`id`),
//  KEY `no` (`no`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='日志表'
const TableLog = "log"
const LogColumns = "`log_id`,`id`,`aid`,`uid`,`time_add`,`mark`,`data`,`no`,`type_login`,`type_client`,`ip`,`msg`"

type LogDB struct {
	LogId      int64   `description:"" db:"log_id" primary:"true"`
	Id         int64   `description:"id" db:"id"`
	Aid        int64   `description:"管理员ID" db:"aid"`
	Uid        int64   `description:"用户id" db:"uid"`
	TimeAdd    *string `description:"创建时间" db:"time_add"`
	Mark       string  `description:"标志自定义标志" db:"mark"`
	Data       *string `description:"其他内容" db:"data"`
	No         string  `description:"单号" db:"no"`
	TypeLogin  int64   `description:"登录方式;302前台还是后台301" db:"type_login"`
	TypeClient int64   `description:"登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他" db:"type_client"`
	Ip         string  `description:"IP" db:"ip"`
	Msg        *string `description:"自定义说明" db:"msg"`
}
type Log struct {
	LogId      int64  `description:"" db:"log_id" primary:"true"`
	Id         int64  `description:"id" db:"id"`
	Aid        int64  `description:"管理员ID" db:"aid"`
	Uid        int64  `description:"用户id" db:"uid"`
	TimeAdd    string `description:"创建时间" db:"time_add"`
	Mark       string `description:"标志自定义标志" db:"mark"`
	Data       string `description:"其他内容" db:"data"`
	No         string `description:"单号" db:"no"`
	TypeLogin  int64  `description:"登录方式;302前台还是后台301" db:"type_login"`
	TypeClient int64  `description:"登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他" db:"type_client"`
	Ip         string `description:"IP" db:"ip"`
	Msg        string `description:"自定义说明" db:"msg"`
}

//Log数据转换
func (l *Log) SetMapValue(m map[string]interface{}) {
	l.LogId = db.Int64Default(m["log_id"])
	l.Id = db.Int64Default(m["id"])
	l.Aid = db.Int64Default(m["aid"])
	l.Uid = db.Int64Default(m["uid"])
	l.TimeAdd = db.StringDefault(m["time_add"])
	l.Mark = db.StringDefault(m["mark"])
	l.Data = db.StringDefault(m["data"])
	l.No = db.StringDefault(m["no"])
	l.TypeLogin = db.Int64Default(m["type_login"])
	l.TypeClient = db.Int64Default(m["type_client"])
	l.Ip = db.StringDefault(m["ip"])
	l.Msg = db.StringDefault(m["msg"])
}

//Log单个查询
func LogGetOne(where string, q db.Query, args ...interface{}) *Log {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(LogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableLog))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Log{}
	ret.SetMapValue(m)
	return ret
}

//Log列表查询
func LogGetList(where string, q db.Query, args ...interface{}) []*Log {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(LogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableLog))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Log
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Log{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Log列表查询
func LogGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Log {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(LogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableLog))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Log
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Log{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
