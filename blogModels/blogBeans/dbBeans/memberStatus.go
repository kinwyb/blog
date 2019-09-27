package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `member_status` (
//  `status_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'UID',
//  `reg_ip` char(15) NOT NULL DEFAULT '' COMMENT '注册IP',
//  `reg_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
//  `reg_type` int(11) NOT NULL DEFAULT '0' COMMENT '注册方式',
//  `reg_app_id` int(11) NOT NULL DEFAULT '1' COMMENT '注册来源',
//  `last_login_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
//  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
//  `last_login_app_id` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录app_id',
//  `login` smallint(5) NOT NULL DEFAULT '0' COMMENT '登录次数',
//  `is_mobile` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '手机号是否已验证1已验证0未验证',
//  `is_email` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '邮箱是否已验证1已验证0未验证',
//  `aid_add` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '客服AID',
//  PRIMARY KEY (`status_id`),
//  KEY `uid` (`uid`)
//) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户状态'
const TableMemberStatus = "member_status"
const MemberStatusColumns = "`status_id`,`uid`,`reg_ip`,`reg_time`,`reg_type`,`reg_app_id`,`last_login_ip`,`last_login_time`,`last_login_app_id`,`login`,`is_mobile`,`is_email`,`aid_add`"

type MemberStatusDB struct {
	StatusId       int64   `description:"" db:"status_id" primary:"true"`
	Uid            int64   `description:"UID" db:"uid"`
	RegIp          string  `description:"注册IP" db:"reg_ip"`
	RegTime        string  `description:"注册时间" db:"reg_time"`
	RegType        int64   `description:"注册方式" db:"reg_type"`
	RegAppId       int64   `description:"注册来源" db:"reg_app_id"`
	LastLoginIp    string  `description:"最后登录IP" db:"last_login_ip"`
	LastLoginTime  *string `description:"最后登录时间" db:"last_login_time"`
	LastLoginAppId int64   `description:"最后登录app_id" db:"last_login_app_id"`
	Login          int     `description:"登录次数" db:"login"`
	IsMobile       int     `description:"手机号是否已验证1已验证0未验证" db:"is_mobile"`
	IsEmail        int     `description:"邮箱是否已验证1已验证0未验证" db:"is_email"`
	AidAdd         int64   `description:"客服AID" db:"aid_add"`
}
type MemberStatus struct {
	StatusId       int64  `description:"" db:"status_id" primary:"true"`
	Uid            int64  `description:"UID" db:"uid"`
	RegIp          string `description:"注册IP" db:"reg_ip"`
	RegTime        string `description:"注册时间" db:"reg_time"`
	RegType        int64  `description:"注册方式" db:"reg_type"`
	RegAppId       int64  `description:"注册来源" db:"reg_app_id"`
	LastLoginIp    string `description:"最后登录IP" db:"last_login_ip"`
	LastLoginTime  string `description:"最后登录时间" db:"last_login_time"`
	LastLoginAppId int64  `description:"最后登录app_id" db:"last_login_app_id"`
	Login          int    `description:"登录次数" db:"login"`
	IsMobile       int    `description:"手机号是否已验证1已验证0未验证" db:"is_mobile"`
	IsEmail        int    `description:"邮箱是否已验证1已验证0未验证" db:"is_email"`
	AidAdd         int64  `description:"客服AID" db:"aid_add"`
}

//MemberStatus数据转换
func (mm *MemberStatus) SetMapValue(m map[string]interface{}) {
	mm.StatusId = db.Int64Default(m["status_id"])
	mm.Uid = db.Int64Default(m["uid"])
	mm.RegIp = db.StringDefault(m["reg_ip"])
	mm.RegTime = db.StringDefault(m["reg_time"])
	mm.RegType = db.Int64Default(m["reg_type"])
	mm.RegAppId = db.Int64Default(m["reg_app_id"])
	mm.LastLoginIp = db.StringDefault(m["last_login_ip"])
	mm.LastLoginTime = db.StringDefault(m["last_login_time"])
	mm.LastLoginAppId = db.Int64Default(m["last_login_app_id"])
	mm.Login = db.IntDefault(m["login"])
	mm.IsMobile = db.IntDefault(m["is_mobile"])
	mm.IsEmail = db.IntDefault(m["is_email"])
	mm.AidAdd = db.Int64Default(m["aid_add"])
}

//MemberStatus单个查询
func MemberStatusGetOne(where string, q db.Query, args ...interface{}) *MemberStatus {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberStatusColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberStatus))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &MemberStatus{}
	ret.SetMapValue(m)
	return ret
}

//MemberStatus列表查询
func MemberStatusGetList(where string, q db.Query, args ...interface{}) []*MemberStatus {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberStatusColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberStatus))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*MemberStatus
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &MemberStatus{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//MemberStatus列表查询
func MemberStatusGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*MemberStatus {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberStatusColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberStatus))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*MemberStatus
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &MemberStatus{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
