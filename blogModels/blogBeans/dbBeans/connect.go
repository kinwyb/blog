package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `connect` (
//  `connect_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id',
//  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
//  `open_id` char(80) NOT NULL DEFAULT '' COMMENT '对应唯一开放id',
//  `token` varchar(80) NOT NULL DEFAULT '' COMMENT '开放密钥',
//  `type` int(11) NOT NULL DEFAULT '1' COMMENT '登录类型1腾讯QQ2新浪微博',
//  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录模块;302前台还是后台301',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//  `extend` varchar(5000) DEFAULT '' COMMENT '扩展参数',
//  PRIMARY KEY (`connect_id`),
//  KEY `openid` (`open_id`),
//  KEY `uid` (`uid`) USING BTREE,
//  KEY `type_id` (`type_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=10012 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='快捷登陆/qq'
const TableConnect = "connect"
const ConnectColumns = "`connect_id`,`type_id`,`uid`,`open_id`,`token`,`type`,`type_login`,`time_add`,`extend`"

type ConnectDB struct {
	ConnectId int64   `description:"" db:"connect_id" primary:"true"`
	TypeId    int64   `description:"类别id" db:"type_id"`
	Uid       int64   `description:"用户id" db:"uid"`
	OpenId    string  `description:"对应唯一开放id" db:"open_id"`
	Token     string  `description:"开放密钥" db:"token"`
	Type      int64   `description:"登录类型1腾讯QQ2新浪微博" db:"type"`
	TypeLogin int64   `description:"登录模块;302前台还是后台301" db:"type_login"`
	TimeAdd   *string `description:"创建时间" db:"time_add"`
	Extend    *string `description:"扩展参数" db:"extend"`
}
type Connect struct {
	ConnectId int64  `description:"" db:"connect_id" primary:"true"`
	TypeId    int64  `description:"类别id" db:"type_id"`
	Uid       int64  `description:"用户id" db:"uid"`
	OpenId    string `description:"对应唯一开放id" db:"open_id"`
	Token     string `description:"开放密钥" db:"token"`
	Type      int64  `description:"登录类型1腾讯QQ2新浪微博" db:"type"`
	TypeLogin int64  `description:"登录模块;302前台还是后台301" db:"type_login"`
	TimeAdd   string `description:"创建时间" db:"time_add"`
	Extend    string `description:"扩展参数" db:"extend"`
}

//Connect数据转换
func (c *Connect) SetMapValue(m map[string]interface{}) {
	c.ConnectId = db.Int64Default(m["connect_id"])
	c.TypeId = db.Int64Default(m["type_id"])
	c.Uid = db.Int64Default(m["uid"])
	c.OpenId = db.StringDefault(m["open_id"])
	c.Token = db.StringDefault(m["token"])
	c.Type = db.Int64Default(m["type"])
	c.TypeLogin = db.Int64Default(m["type_login"])
	c.TimeAdd = db.StringDefault(m["time_add"])
	c.Extend = db.StringDefault(m["extend"])
}

//Connect单个查询
func ConnectGetOne(where string, q db.Query, args ...interface{}) *Connect {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ConnectColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableConnect))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Connect{}
	ret.SetMapValue(m)
	return ret
}

//Connect列表查询
func ConnectGetList(where string, q db.Query, args ...interface{}) []*Connect {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ConnectColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableConnect))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Connect
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Connect{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Connect列表查询
func ConnectGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Connect {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(ConnectColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableConnect))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Connect
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Connect{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
