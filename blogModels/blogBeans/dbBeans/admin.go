package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `admin` (
//  `aid` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `username` char(30) NOT NULL DEFAULT '' COMMENT '用户名',
//  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
//  `mail` varchar(80) NOT NULL DEFAULT '' COMMENT '邮箱',
//  `salt` varchar(10) NOT NULL DEFAULT '' COMMENT '干扰码',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
//  `time_update` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//  `ip` char(15) NOT NULL DEFAULT '' COMMENT '添加IP',
//  `job_no` varchar(15) NOT NULL DEFAULT '' COMMENT '工号',
//  `nick_name` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
//  `true_name` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
//  `qq` varchar(50) NOT NULL DEFAULT '' COMMENT 'qq',
//  `phone` varchar(50) NOT NULL DEFAULT '' COMMENT '电话',
//  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
//  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '删除0否1是',
//  PRIMARY KEY (`aid`),
//  KEY `username` (`username`),
//  KEY `is_del` (`is_del`)
//) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员表'
const TableAdmin = "admin"
const AdminColumns = "`aid`,`username`,`password`,`mail`,`salt`,`time_add`,`time_update`,`ip`,`job_no`,`nick_name`,`true_name`,`qq`,`phone`,`mobile`,`is_del`"

type AdminDB struct {
	Aid        int64   `description:"" db:"aid" primary:"true"`
	Username   string  `description:"用户名" db:"username"`
	Password   string  `description:"密码" db:"password"`
	Mail       string  `description:"邮箱" db:"mail"`
	Salt       string  `description:"干扰码" db:"salt"`
	TimeAdd    *string `description:"注册时间" db:"time_add"`
	TimeUpdate *string `description:"更新时间" db:"time_update"`
	Ip         string  `description:"添加IP" db:"ip"`
	JobNo      string  `description:"工号" db:"job_no"`
	NickName   string  `description:"昵称" db:"nick_name"`
	TrueName   string  `description:"真实姓名" db:"true_name"`
	Qq         string  `description:"qq" db:"qq"`
	Phone      string  `description:"电话" db:"phone"`
	Mobile     string  `description:"手机" db:"mobile"`
	IsDel      int     `description:"删除0否1是" db:"is_del"`
}
type Admin struct {
	Aid        int64  `description:"" db:"aid" primary:"true"`
	Username   string `description:"用户名" db:"username"`
	Password   string `description:"密码" db:"password"`
	Mail       string `description:"邮箱" db:"mail"`
	Salt       string `description:"干扰码" db:"salt"`
	TimeAdd    string `description:"注册时间" db:"time_add"`
	TimeUpdate string `description:"更新时间" db:"time_update"`
	Ip         string `description:"添加IP" db:"ip"`
	JobNo      string `description:"工号" db:"job_no"`
	NickName   string `description:"昵称" db:"nick_name"`
	TrueName   string `description:"真实姓名" db:"true_name"`
	Qq         string `description:"qq" db:"qq"`
	Phone      string `description:"电话" db:"phone"`
	Mobile     string `description:"手机" db:"mobile"`
	IsDel      int    `description:"删除0否1是" db:"is_del"`
}

//Admin数据转换
func (a *Admin) SetMapValue(m map[string]interface{}) {
	a.Aid = db.Int64Default(m["aid"])
	a.Username = db.StringDefault(m["username"])
	a.Password = db.StringDefault(m["password"])
	a.Mail = db.StringDefault(m["mail"])
	a.Salt = db.StringDefault(m["salt"])
	a.TimeAdd = db.StringDefault(m["time_add"])
	a.TimeUpdate = db.StringDefault(m["time_update"])
	a.Ip = db.StringDefault(m["ip"])
	a.JobNo = db.StringDefault(m["job_no"])
	a.NickName = db.StringDefault(m["nick_name"])
	a.TrueName = db.StringDefault(m["true_name"])
	a.Qq = db.StringDefault(m["qq"])
	a.Phone = db.StringDefault(m["phone"])
	a.Mobile = db.StringDefault(m["mobile"])
	a.IsDel = db.IntDefault(m["is_del"])
}

//Admin单个查询
func AdminGetOne(where string, q db.Query, args ...interface{}) *Admin {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdmin))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Admin{}
	ret.SetMapValue(m)
	return ret
}

//Admin列表查询
func AdminGetList(where string, q db.Query, args ...interface{}) []*Admin {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdmin))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Admin
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Admin{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Admin列表查询
func AdminGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Admin {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AdminColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAdmin))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Admin
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Admin{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
