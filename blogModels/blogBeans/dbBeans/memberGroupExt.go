package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `member_group_ext` (
//  `group_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
//  PRIMARY KEY (`group_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='会员用户组扩展'
const TableMemberGroupExt = "member_group_ext"
const MemberGroupExtColumns = "`group_id`"

type MemberGroupExtDB struct {
	GroupId int64 `description:"用户ID" db:"group_id" primary:"true"`
}
type MemberGroupExt struct {
	GroupId int64 `description:"用户ID" db:"group_id" primary:"true"`
}

//MemberGroupExt数据转换
func (mm *MemberGroupExt) SetMapValue(m map[string]interface{}) {
	mm.GroupId = db.Int64Default(m["group_id"])
}

//MemberGroupExt单个查询
func MemberGroupExtGetOne(where string, q db.Query, args ...interface{}) *MemberGroupExt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberGroupExtColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberGroupExt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &MemberGroupExt{}
	ret.SetMapValue(m)
	return ret
}

//MemberGroupExt列表查询
func MemberGroupExtGetList(where string, q db.Query, args ...interface{}) []*MemberGroupExt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberGroupExtColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberGroupExt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*MemberGroupExt
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &MemberGroupExt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//MemberGroupExt列表查询
func MemberGroupExtGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*MemberGroupExt {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(MemberGroupExtColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableMemberGroupExt))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*MemberGroupExt
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &MemberGroupExt{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
