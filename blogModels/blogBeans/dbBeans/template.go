package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `template` (
//  `template_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '模板ID',
//  `name` varchar(80) NOT NULL DEFAULT '' COMMENT '模板名称(中文)',
//  `mark` varchar(80) NOT NULL DEFAULT '' COMMENT '模板名称标志(英文)（调用时使用）',
//  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '邮件标题',
//  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '模板类型1短信模板2邮箱模板',
//  `use` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '用途',
//  `content` text,
//  `remark` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//  `time_update` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//  `code_num` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '验证码位数',
//  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
//  PRIMARY KEY (`template_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='会员表'
const TableTemplate = "template"
const TemplateColumns = "`template_id`,`name`,`mark`,`title`,`type`,`use`,`content`,`remark`,`time_add`,`time_update`,`code_num`,`aid`"

type TemplateDB struct {
	TemplateId int64   `description:"模板ID" db:"template_id" primary:"true"`
	Name       string  `description:"模板名称(中文)" db:"name"`
	Mark       string  `description:"模板名称标志(英文)（调用时使用）" db:"mark"`
	Title      string  `description:"邮件标题" db:"title"`
	Type       int     `description:"模板类型1短信模板2邮箱模板" db:"type"`
	Use        int     `description:"用途" db:"use"`
	Content    *string `description:"" db:"content"`
	Remark     string  `description:"备注" db:"remark"`
	TimeAdd    *string `description:"创建时间" db:"time_add"`
	TimeUpdate *string `description:"更新时间" db:"time_update"`
	CodeNum    int     `description:"验证码位数" db:"code_num"`
	Aid        int64   `description:"添加人" db:"aid"`
}
type Template struct {
	TemplateId int64  `description:"模板ID" db:"template_id" primary:"true"`
	Name       string `description:"模板名称(中文)" db:"name"`
	Mark       string `description:"模板名称标志(英文)（调用时使用）" db:"mark"`
	Title      string `description:"邮件标题" db:"title"`
	Type       int    `description:"模板类型1短信模板2邮箱模板" db:"type"`
	Use        int    `description:"用途" db:"use"`
	Content    string `description:"" db:"content"`
	Remark     string `description:"备注" db:"remark"`
	TimeAdd    string `description:"创建时间" db:"time_add"`
	TimeUpdate string `description:"更新时间" db:"time_update"`
	CodeNum    int    `description:"验证码位数" db:"code_num"`
	Aid        int64  `description:"添加人" db:"aid"`
}

//Template数据转换
func (t *Template) SetMapValue(m map[string]interface{}) {
	t.TemplateId = db.Int64Default(m["template_id"])
	t.Name = db.StringDefault(m["name"])
	t.Mark = db.StringDefault(m["mark"])
	t.Title = db.StringDefault(m["title"])
	t.Type = db.IntDefault(m["type"])
	t.Use = db.IntDefault(m["use"])
	t.Content = db.StringDefault(m["content"])
	t.Remark = db.StringDefault(m["remark"])
	t.TimeAdd = db.StringDefault(m["time_add"])
	t.TimeUpdate = db.StringDefault(m["time_update"])
	t.CodeNum = db.IntDefault(m["code_num"])
	t.Aid = db.Int64Default(m["aid"])
}

//Template单个查询
func TemplateGetOne(where string, q db.Query, args ...interface{}) *Template {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TemplateColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableTemplate))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Template{}
	ret.SetMapValue(m)
	return ret
}

//Template列表查询
func TemplateGetList(where string, q db.Query, args ...interface{}) []*Template {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TemplateColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableTemplate))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Template
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Template{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Template列表查询
func TemplateGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Template {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(TemplateColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableTemplate))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Template
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Template{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
