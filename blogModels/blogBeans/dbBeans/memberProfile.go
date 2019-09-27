package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `member_profile` (
//  `profile_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'UID',
//  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别1男2女3中性0保密',
//  `job` varchar(50) NOT NULL DEFAULT '' COMMENT '担任职务',
//  `qq` varchar(20) NOT NULL DEFAULT '',
//  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '电话',
//  `county` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '国家',
//  `province` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '省',
//  `city` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '市',
//  `district` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '区',
//  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
//  `wechat` varchar(20) NOT NULL DEFAULT '' COMMENT '微信',
//  `remark_admin` text COMMENT '客服备注',
//  PRIMARY KEY (`profile_id`),
//  KEY `uid` (`uid`)
//) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户其他介绍'
const TableMemberProfile = "member_profile"
const MemberProfileColumns = "`profile_id`,`uid`,`sex`,`job`,`qq`,`phone`,`county`,`province`,`city`,`district`,`address`,`wechat`,`remark_admin`"

type MemberProfileDB struct {
	ProfileId   int64   `description:"" db:"profile_id" primary:"true"`
	Uid         int64   `description:"UID" db:"uid"`
	Sex         int     `description:"性别1男2女3中性0保密" db:"sex"`
	Job         string  `description:"担任职务" db:"job"`
	Qq          string  `description:"" db:"qq"`
	Phone       string  `description:"电话" db:"phone"`
	County      int64   `description:"国家" db:"county"`
	Province    int64   `description:"省" db:"province"`
	City        int64   `description:"市" db:"city"`
	District    int64   `description:"区" db:"district"`
	Address     string  `description:"地址" db:"address"`
	Wechat      string  `description:"微信" db:"wechat"`
	RemarkAdmin *string `description:"客服备注" db:"remark_admin"`
}
type MemberProfile struct {
	ProfileId   int64  `description:"" db:"profile_id" primary:"true"`
	Uid         int64  `description:"UID" db:"uid"`
	Sex         int    `description:"性别1男2女3中性0保密" db:"sex"`
	Job         string `description:"担任职务" db:"job"`
	Qq          string `description:"" db:"qq"`
	Phone       string `description:"电话" db:"phone"`
	County      int64  `description:"国家" db:"county"`
	Province    int64  `description:"省" db:"province"`
	City        int64  `description:"市" db:"city"`
	District    int64  `description:"区" db:"district"`
	Address     string `description:"地址" db:"address"`
	Wechat      string `description:"微信" db:"wechat"`
	RemarkAdmin string `description:"客服备注" db:"remark_admin"`
}

//MemberProfile数据转换
func (mm *MemberProfile) SetMapValue(m map[string]interface{}) {
	mm.ProfileId = db.Int64Default(m["profile_id"])
	mm.Uid = db.Int64Default(m["uid"])
	mm.Sex = db.IntDefault(m["sex"])
	mm.Job = db.StringDefault(m["job"])
	mm.Qq = db.StringDefault(m["qq"])
	mm.Phone = db.StringDefault(m["phone"])
	mm.County = db.Int64Default(m["county"])
	mm.Province = db.Int64Default(m["province"])
	mm.City = db.Int64Default(m["city"])
	mm.District = db.Int64Default(m["district"])
	mm.Address = db.StringDefault(m["address"])
	mm.Wechat = db.StringDefault(m["wechat"])
	mm.RemarkAdmin = db.StringDefault(m["remark_admin"])
}

//MemberProfile单个查询
func MemberProfileGetOne(where string, q db.Query, args ...interface{}) *MemberProfile {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberProfileColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberProfile))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &MemberProfile{}
	ret.SetMapValue(m)
	return ret
}

//MemberProfile列表查询
func MemberProfileGetList(where string, q db.Query, args ...interface{}) []*MemberProfile {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberProfileColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberProfile))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*MemberProfile
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &MemberProfile{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//MemberProfile列表查询
func MemberProfileGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*MemberProfile {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberProfileColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberProfile))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*MemberProfile
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &MemberProfile{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
