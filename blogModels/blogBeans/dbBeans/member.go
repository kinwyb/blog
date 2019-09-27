package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `member` (
//  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `mobile` char(11) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
//  `username` char(30) NOT NULL DEFAULT '' COMMENT '用户名',
//  `mail` char(32) NOT NULL DEFAULT '' COMMENT '邮箱',
//  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
//  `salt` char(6) NOT NULL DEFAULT '' COMMENT '干扰码',
//  `reg_ip` char(15) NOT NULL DEFAULT '' COMMENT '注册IP',
//  `reg_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
//  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态0正常1删除',
//  `group_id` int(11) unsigned NOT NULL DEFAULT '410' COMMENT '用户组ID',
//  `true_name` varchar(32) NOT NULL DEFAULT '' COMMENT '真实姓名',
//  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '店铺名称',
//  PRIMARY KEY (`uid`),
//  KEY `mobile` (`mobile`),
//  KEY `is_del` (`is_del`),
//  KEY `username` (`username`),
//  KEY `email` (`mail`),
//  KEY `group_id` (`group_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户表'
const TableMember = "member"
const MemberColumns = "`uid`,`mobile`,`username`,`mail`,`password`,`salt`,`reg_ip`,`reg_time`,`is_del`,`group_id`,`true_name`,`name`"

type MemberDB struct {
	Uid      int64  `description:"" db:"uid" primary:"true"`
	Mobile   string `description:"" db:"mobile"`
	Username string `description:"用户名" db:"username"`
	Mail     string `description:"邮箱" db:"mail"`
	Password string `description:"密码" db:"password"`
	Salt     string `description:"干扰码" db:"salt"`
	RegIp    string `description:"注册IP" db:"reg_ip"`
	RegTime  string `description:"注册时间" db:"reg_time"`
	IsDel    int    `description:"状态0正常1删除" db:"is_del"`
	GroupId  int64  `description:"用户组ID" db:"group_id"`
	TrueName string `description:"真实姓名" db:"true_name"`
	Name     string `description:"店铺名称" db:"name"`
}
type Member struct {
	Uid      int64  `description:"" db:"uid" primary:"true"`
	Mobile   string `description:"" db:"mobile"`
	Username string `description:"用户名" db:"username"`
	Mail     string `description:"邮箱" db:"mail"`
	Password string `description:"密码" db:"password"`
	Salt     string `description:"干扰码" db:"salt"`
	RegIp    string `description:"注册IP" db:"reg_ip"`
	RegTime  string `description:"注册时间" db:"reg_time"`
	IsDel    int    `description:"状态0正常1删除" db:"is_del"`
	GroupId  int64  `description:"用户组ID" db:"group_id"`
	TrueName string `description:"真实姓名" db:"true_name"`
	Name     string `description:"店铺名称" db:"name"`
}

//Member数据转换
func (mm *Member) SetMapValue(m map[string]interface{}) {
	mm.Uid = db.Int64Default(m["uid"])
	mm.Mobile = db.StringDefault(m["mobile"])
	mm.Username = db.StringDefault(m["username"])
	mm.Mail = db.StringDefault(m["mail"])
	mm.Password = db.StringDefault(m["password"])
	mm.Salt = db.StringDefault(m["salt"])
	mm.RegIp = db.StringDefault(m["reg_ip"])
	mm.RegTime = db.StringDefault(m["reg_time"])
	mm.IsDel = db.IntDefault(m["is_del"])
	mm.GroupId = db.Int64Default(m["group_id"])
	mm.TrueName = db.StringDefault(m["true_name"])
	mm.Name = db.StringDefault(m["name"])
}

//Member单个查询
func MemberGetOne(where string, q db.Query, args ...interface{}) *Member {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMember))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Member{}
	ret.SetMapValue(m)
	return ret
}

//Member列表查询
func MemberGetList(where string, q db.Query, args ...interface{}) []*Member {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMember))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Member
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Member{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Member列表查询
func MemberGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Member {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMember))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Member
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Member{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
