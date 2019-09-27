package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `attachment` (
//  `attachment_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件ID',
//  `module` char(32) NOT NULL DEFAULT '' COMMENT '模块',
//  `mark` char(60) NOT NULL DEFAULT '' COMMENT '标记标志',
//  `type_id` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '类别ID',
//  `name` char(50) NOT NULL DEFAULT '' COMMENT '保存的文件名称',
//  `name_original` varchar(255) NOT NULL DEFAULT '' COMMENT '原文件名',
//  `path` char(200) NOT NULL DEFAULT '' COMMENT '文件路径',
//  `size` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小',
//  `ext` char(10) NOT NULL DEFAULT '' COMMENT '文件后缀',
//  `is_image` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否图片1是0否',
//  `is_thumb` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否缩略图1是0否',
//  `downloads` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '下载次数',
//  `time_add` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间上传时间',
//  `ip` char(15) NOT NULL DEFAULT '' COMMENT '上传IP',
//  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态99正常;',
//  `md5` char(32) NOT NULL DEFAULT '' COMMENT 'md5',
//  `sha1` char(40) NOT NULL DEFAULT '' COMMENT 'sha1',
//  `id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属ID',
//  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '后台管理员ID',
//  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前台用户ID',
//  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示1是0否',
//  `http` varchar(100) NOT NULL DEFAULT '' COMMENT '图片http地址',
//  PRIMARY KEY (`attachment_id`),
//  KEY `md5` (`md5`),
//  KEY `module` (`module`),
//  KEY `mark` (`mark`),
//  KEY `id` (`id`),
//  KEY `status` (`status`),
//  KEY `aid` (`aid`),
//  KEY `uid` (`uid`),
//  KEY `is_show` (`is_show`)
//) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='附件表'
const TableAttachment = "attachment"
const AttachmentColumns = "`attachment_id`,`module`,`mark`,`type_id`,`name`,`name_original`,`path`,`size`,`ext`,`is_image`,`is_thumb`,`downloads`,`time_add`,`ip`,`status`,`md5`,`sha1`,`id`,`aid`,`uid`,`is_show`,`http`"

type AttachmentDB struct {
	AttachmentId int64  `description:"附件ID" db:"attachment_id" primary:"true"`
	Module       string `description:"模块" db:"module"`
	Mark         string `description:"标记标志" db:"mark"`
	TypeId       int64  `description:"类别ID" db:"type_id"`
	Name         string `description:"保存的文件名称" db:"name"`
	NameOriginal string `description:"原文件名" db:"name_original"`
	Path         string `description:"文件路径" db:"path"`
	Size         int64  `description:"文件大小" db:"size"`
	Ext          string `description:"文件后缀" db:"ext"`
	IsImage      int    `description:"是否图片1是0否" db:"is_image"`
	IsThumb      int    `description:"是否缩略图1是0否" db:"is_thumb"`
	Downloads    int64  `description:"下载次数" db:"downloads"`
	TimeAdd      string `description:"添加时间上传时间" db:"time_add"`
	Ip           string `description:"上传IP" db:"ip"`
	Status       int    `description:"状态99正常;" db:"status"`
	Md5          string `description:"md5" db:"md5"`
	Sha1         string `description:"sha1" db:"sha1"`
	Id           int64  `description:"所属ID" db:"id"`
	Aid          int64  `description:"后台管理员ID" db:"aid"`
	Uid          int64  `description:"前台用户ID" db:"uid"`
	IsShow       int    `description:"是否显示1是0否" db:"is_show"`
	Http         string `description:"图片http地址" db:"http"`
}
type Attachment struct {
	AttachmentId int64  `description:"附件ID" db:"attachment_id" primary:"true"`
	Module       string `description:"模块" db:"module"`
	Mark         string `description:"标记标志" db:"mark"`
	TypeId       int64  `description:"类别ID" db:"type_id"`
	Name         string `description:"保存的文件名称" db:"name"`
	NameOriginal string `description:"原文件名" db:"name_original"`
	Path         string `description:"文件路径" db:"path"`
	Size         int64  `description:"文件大小" db:"size"`
	Ext          string `description:"文件后缀" db:"ext"`
	IsImage      int    `description:"是否图片1是0否" db:"is_image"`
	IsThumb      int    `description:"是否缩略图1是0否" db:"is_thumb"`
	Downloads    int64  `description:"下载次数" db:"downloads"`
	TimeAdd      string `description:"添加时间上传时间" db:"time_add"`
	Ip           string `description:"上传IP" db:"ip"`
	Status       int    `description:"状态99正常;" db:"status"`
	Md5          string `description:"md5" db:"md5"`
	Sha1         string `description:"sha1" db:"sha1"`
	Id           int64  `description:"所属ID" db:"id"`
	Aid          int64  `description:"后台管理员ID" db:"aid"`
	Uid          int64  `description:"前台用户ID" db:"uid"`
	IsShow       int    `description:"是否显示1是0否" db:"is_show"`
	Http         string `description:"图片http地址" db:"http"`
}

//Attachment数据转换
func (a *Attachment) SetMapValue(m map[string]interface{}) {
	a.AttachmentId = db.Int64Default(m["attachment_id"])
	a.Module = db.StringDefault(m["module"])
	a.Mark = db.StringDefault(m["mark"])
	a.TypeId = db.Int64Default(m["type_id"])
	a.Name = db.StringDefault(m["name"])
	a.NameOriginal = db.StringDefault(m["name_original"])
	a.Path = db.StringDefault(m["path"])
	a.Size = db.Int64Default(m["size"])
	a.Ext = db.StringDefault(m["ext"])
	a.IsImage = db.IntDefault(m["is_image"])
	a.IsThumb = db.IntDefault(m["is_thumb"])
	a.Downloads = db.Int64Default(m["downloads"])
	a.TimeAdd = db.StringDefault(m["time_add"])
	a.Ip = db.StringDefault(m["ip"])
	a.Status = db.IntDefault(m["status"])
	a.Md5 = db.StringDefault(m["md5"])
	a.Sha1 = db.StringDefault(m["sha1"])
	a.Id = db.Int64Default(m["id"])
	a.Aid = db.Int64Default(m["aid"])
	a.Uid = db.Int64Default(m["uid"])
	a.IsShow = db.IntDefault(m["is_show"])
	a.Http = db.StringDefault(m["http"])
}

//Attachment单个查询
func AttachmentGetOne(where string, q db.Query, args ...interface{}) *Attachment {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AttachmentColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAttachment))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Attachment{}
	ret.SetMapValue(m)
	return ret
}

//Attachment列表查询
func AttachmentGetList(where string, q db.Query, args ...interface{}) []*Attachment {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AttachmentColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAttachment))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Attachment
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Attachment{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Attachment列表查询
func AttachmentGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Attachment {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(AttachmentColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableAttachment))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Attachment
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Attachment{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
