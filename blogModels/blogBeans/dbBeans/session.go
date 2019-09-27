package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `session` (
//  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
//  `ip` char(15) NOT NULL DEFAULT '' COMMENT 'IP',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
//  `error_count` tinyint(1) NOT NULL DEFAULT '0' COMMENT '密码输入错误次数',
//  `app_id` int(11) NOT NULL DEFAULT '0' COMMENT '登录应用',
//  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录方式;302前台还是后台301',
//  `md5` char(32) NOT NULL DEFAULT '' COMMENT 'md5',
//  `type_client` int(11) NOT NULL DEFAULT '0' COMMENT '登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他',
//  PRIMARY KEY (`id`),
//  KEY `uid` (`uid`,`type_login`,`type_client`)
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='SESSION'
const TableSession = "session"
const SessionColumns = "`id`,`uid`,`ip`,`time_add`,`error_count`,`app_id`,`type_login`,`md5`,`type_client`"

type SessionDB struct {
	Id         int64   `description:"" db:"id" primary:"true"`
	Uid        int64   `description:"用户UID" db:"uid"`
	Ip         string  `description:"IP" db:"ip"`
	TimeAdd    *string `description:"登录时间" db:"time_add"`
	ErrorCount int     `description:"密码输入错误次数" db:"error_count"`
	AppId      int64   `description:"登录应用" db:"app_id"`
	TypeLogin  int64   `description:"登录方式;302前台还是后台301" db:"type_login"`
	Md5        string  `description:"md5" db:"md5"`
	TypeClient int64   `description:"登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他" db:"type_client"`
}
type Session struct {
	Id         int64  `description:"" db:"id" primary:"true"`
	Uid        int64  `description:"用户UID" db:"uid"`
	Ip         string `description:"IP" db:"ip"`
	TimeAdd    string `description:"登录时间" db:"time_add"`
	ErrorCount int    `description:"密码输入错误次数" db:"error_count"`
	AppId      int64  `description:"登录应用" db:"app_id"`
	TypeLogin  int64  `description:"登录方式;302前台还是后台301" db:"type_login"`
	Md5        string `description:"md5" db:"md5"`
	TypeClient int64  `description:"登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他" db:"type_client"`
}

//Session数据转换
func (s *Session) SetMapValue(m map[string]interface{}) {
	s.Id = db.Int64Default(m["id"])
	s.Uid = db.Int64Default(m["uid"])
	s.Ip = db.StringDefault(m["ip"])
	s.TimeAdd = db.StringDefault(m["time_add"])
	s.ErrorCount = db.IntDefault(m["error_count"])
	s.AppId = db.Int64Default(m["app_id"])
	s.TypeLogin = db.Int64Default(m["type_login"])
	s.Md5 = db.StringDefault(m["md5"])
	s.TypeClient = db.Int64Default(m["type_client"])
}

//Session单个查询
func SessionGetOne(where string, q db.Query, args ...interface{}) *Session {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SessionColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSession))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Session{}
	ret.SetMapValue(m)
	return ret
}

//Session列表查询
func SessionGetList(where string, q db.Query, args ...interface{}) []*Session {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SessionColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSession))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Session
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Session{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Session列表查询
func SessionGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Session {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(SessionColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableSession))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Session
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Session{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
