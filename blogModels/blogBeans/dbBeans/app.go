package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `app` (
//  `app_id` int(11) NOT NULL AUTO_INCREMENT,
//  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT 'app_id,来源type表',
//  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
//  `mark` char(32) NOT NULL DEFAULT '' COMMENT '标志',
//  `setting` varchar(5000) DEFAULT NULL COMMENT '扩展参数',
//  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
//  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除0否1是',
//  PRIMARY KEY (`app_id`),
//  UNIQUE KEY `type_id` (`type_id`) USING BTREE
//) ENGINE=InnoDB AUTO_INCREMENT=10012 DEFAULT CHARSET=utf8 COMMENT='应用'
const TableApp = "app"
const AppColumns = "`app_id`,`type_id`,`name`,`mark`,`setting`,`remark`,`time_add`,`is_del`"

type AppDB struct {
	AppId   int64   `description:"" db:"app_id" primary:"true"`
	TypeId  int64   `description:"app_id,来源type表" db:"type_id"`
	Name    string  `description:"名称" db:"name"`
	Mark    string  `description:"标志" db:"mark"`
	Setting *string `description:"扩展参数" db:"setting"`
	Remark  *string `description:"备注" db:"remark"`
	TimeAdd *string `description:"添加时间" db:"time_add"`
	IsDel   int64   `description:"是否删除0否1是" db:"is_del"`
}
type App struct {
	AppId   int64  `description:"" db:"app_id" primary:"true"`
	TypeId  int64  `description:"app_id,来源type表" db:"type_id"`
	Name    string `description:"名称" db:"name"`
	Mark    string `description:"标志" db:"mark"`
	Setting string `description:"扩展参数" db:"setting"`
	Remark  string `description:"备注" db:"remark"`
	TimeAdd string `description:"添加时间" db:"time_add"`
	IsDel   int64  `description:"是否删除0否1是" db:"is_del"`
}

//App数据转换
func (a *App) SetMapValue(m map[string]interface{}) {
	a.AppId = db.Int64Default(m["app_id"])
	a.TypeId = db.Int64Default(m["type_id"])
	a.Name = db.StringDefault(m["name"])
	a.Mark = db.StringDefault(m["mark"])
	a.Setting = db.StringDefault(m["setting"])
	a.Remark = db.StringDefault(m["remark"])
	a.TimeAdd = db.StringDefault(m["time_add"])
	a.IsDel = db.Int64Default(m["is_del"])
}

//App单个查询
func AppGetOne(where string, q db.Query, args ...interface{}) *App {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AppColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableApp))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &App{}
	ret.SetMapValue(m)
	return ret
}

//App列表查询
func AppGetList(where string, q db.Query, args ...interface{}) []*App {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AppColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableApp))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*App
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &App{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//App列表查询
func AppGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*App {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AppColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableApp))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*App
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &App{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
