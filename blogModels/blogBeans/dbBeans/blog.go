package dbBeans

import (
	"strings"

	"github.com/kinwyb/go/db"
)

//CREATE TABLE `blog` (
//  `blog_id` int(11) NOT NULL AUTO_INCREMENT,
//  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员AID',
//  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
//  `is_open` tinyint(1) NOT NULL DEFAULT '1' COMMENT '启用1是0否',
//  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
//  `time_system` timestamp NULL DEFAULT NULL COMMENT '创建时间,系统时间不可修改',
//  `time_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//  `time_add` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间,可修改',
//  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
//  `author` varchar(255) NOT NULL DEFAULT '' COMMENT '作者',
//  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '网址',
//  `url_source` varchar(255) NOT NULL DEFAULT '' COMMENT '来源地址(转载)',
//  `url_rewrite` char(255) NOT NULL DEFAULT '' COMMENT '自定义伪静态Url',
//  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '摘要',
//  `content` text COMMENT '内容',
//  `type` int(11) NOT NULL DEFAULT '0' COMMENT '类型0文章10001博客栏目',
//  `module_id` int(10) NOT NULL DEFAULT '0' COMMENT '模块10019技术10018生活',
//  `source_id` int(11) NOT NULL DEFAULT '0' COMMENT '来源:后台，接口，其他',
//  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别ID，原创，转载，翻译',
//  `cat_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类ID，栏目',
//  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
//  `thumb` varchar(255) NOT NULL DEFAULT '' COMMENT '缩略图',
//  `is_relevant` tinyint(1) NOT NULL DEFAULT '0' COMMENT '相关文章1是0否',
//  `is_jump` tinyint(1) NOT NULL DEFAULT '0' COMMENT '跳转1是0否',
//  `is_comment` tinyint(1) NOT NULL DEFAULT '1' COMMENT '允许评论1是0否',
//  `is_read` int(11) NOT NULL DEFAULT '10014' COMMENT '是否阅读10014未看10015在看10016已看',
//  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
//  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
//  PRIMARY KEY (`blog_id`),
//  KEY `is_del` (`is_del`,`is_open`,`status`,`type_id`,`cat_id`,`sort`),
//  KEY `url_rewrite` (`url_rewrite`),
//  KEY `type` (`type`),
//  KEY `module_id` (`module_id`) USING BTREE,
//  KEY `source_id` (`source_id`)
//) ENGINE=InnoDB AUTO_INCREMENT=132 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='博客内容'
const TableBlog = "blog"
const BlogColumns = "`blog_id`,`aid`,`is_del`,`is_open`,`status`,`time_system`,`time_update`,`time_add`,`title`,`author`,`url`,`url_source`,`url_rewrite`,`description`,`content`,`type`,`module_id`,`source_id`,`type_id`,`cat_id`,`tag`,`thumb`,`is_relevant`,`is_jump`,`is_comment`,`is_read`,`sort`,`remark`"

type BlogDB struct {
	BlogId      int64   `description:"" db:"blog_id" primary:"true"`
	Aid         int64   `description:"管理员AID" db:"aid"`
	IsDel       int     `description:"是否删除1是0否" db:"is_del"`
	IsOpen      int     `description:"启用1是0否" db:"is_open"`
	Status      int64   `description:"状态" db:"status"`
	TimeSystem  *string `description:"创建时间,系统时间不可修改" db:"time_system"`
	TimeUpdate  *string `description:"更新时间" db:"time_update"`
	TimeAdd     *string `description:"添加时间,可修改" db:"time_add"`
	Title       string  `description:"标题" db:"title"`
	Author      string  `description:"作者" db:"author"`
	Url         string  `description:"网址" db:"url"`
	UrlSource   string  `description:"来源地址(转载)" db:"url_source"`
	UrlRewrite  string  `description:"自定义伪静态Url" db:"url_rewrite"`
	Description string  `description:"摘要" db:"description"`
	Content     *string `description:"内容" db:"content"`
	Type        int64   `description:"类型0文章10001博客栏目" db:"type"`
	ModuleId    int64   `description:"模块10019技术10018生活" db:"module_id"`
	SourceId    int64   `description:"来源:后台，接口，其他" db:"source_id"`
	TypeId      int64   `description:"类别ID，原创，转载，翻译" db:"type_id"`
	CatId       int64   `description:"分类ID，栏目" db:"cat_id"`
	Tag         string  `description:"标签" db:"tag"`
	Thumb       string  `description:"缩略图" db:"thumb"`
	IsRelevant  int     `description:"相关文章1是0否" db:"is_relevant"`
	IsJump      int     `description:"跳转1是0否" db:"is_jump"`
	IsComment   int     `description:"允许评论1是0否" db:"is_comment"`
	IsRead      int64   `description:"是否阅读10014未看10015在看10016已看" db:"is_read"`
	Sort        int64   `description:"排序" db:"sort"`
	Remark      string  `description:"备注" db:"remark"`
}
type Blog struct {
	BlogId      int64  `description:"" db:"blog_id" primary:"true"`
	Aid         int64  `description:"管理员AID" db:"aid"`
	IsDel       int    `description:"是否删除1是0否" db:"is_del"`
	IsOpen      int    `description:"启用1是0否" db:"is_open"`
	Status      int64  `description:"状态" db:"status"`
	TimeSystem  string `description:"创建时间,系统时间不可修改" db:"time_system"`
	TimeUpdate  string `description:"更新时间" db:"time_update"`
	TimeAdd     string `description:"添加时间,可修改" db:"time_add"`
	Title       string `description:"标题" db:"title"`
	Author      string `description:"作者" db:"author"`
	Url         string `description:"网址" db:"url"`
	UrlSource   string `description:"来源地址(转载)" db:"url_source"`
	UrlRewrite  string `description:"自定义伪静态Url" db:"url_rewrite"`
	Description string `description:"摘要" db:"description"`
	Content     string `description:"内容" db:"content"`
	Type        int64  `description:"类型0文章10001博客栏目" db:"type"`
	ModuleId    int64  `description:"模块10019技术10018生活" db:"module_id"`
	SourceId    int64  `description:"来源:后台，接口，其他" db:"source_id"`
	TypeId      int64  `description:"类别ID，原创，转载，翻译" db:"type_id"`
	CatId       int64  `description:"分类ID，栏目" db:"cat_id"`
	Tag         string `description:"标签" db:"tag"`
	Thumb       string `description:"缩略图" db:"thumb"`
	IsRelevant  int    `description:"相关文章1是0否" db:"is_relevant"`
	IsJump      int    `description:"跳转1是0否" db:"is_jump"`
	IsComment   int    `description:"允许评论1是0否" db:"is_comment"`
	IsRead      int64  `description:"是否阅读10014未看10015在看10016已看" db:"is_read"`
	Sort        int64  `description:"排序" db:"sort"`
	Remark      string `description:"备注" db:"remark"`
}

//Blog数据转换
func (b *Blog) SetMapValue(m map[string]interface{}) {
	b.BlogId = db.Int64Default(m["blog_id"])
	b.Aid = db.Int64Default(m["aid"])
	b.IsDel = db.IntDefault(m["is_del"])
	b.IsOpen = db.IntDefault(m["is_open"])
	b.Status = db.Int64Default(m["status"])
	b.TimeSystem = db.StringDefault(m["time_system"])
	b.TimeUpdate = db.StringDefault(m["time_update"])
	b.TimeAdd = db.StringDefault(m["time_add"])
	b.Title = db.StringDefault(m["title"])
	b.Author = db.StringDefault(m["author"])
	b.Url = db.StringDefault(m["url"])
	b.UrlSource = db.StringDefault(m["url_source"])
	b.UrlRewrite = db.StringDefault(m["url_rewrite"])
	b.Description = db.StringDefault(m["description"])
	b.Content = db.StringDefault(m["content"])
	b.Type = db.Int64Default(m["type"])
	b.ModuleId = db.Int64Default(m["module_id"])
	b.SourceId = db.Int64Default(m["source_id"])
	b.TypeId = db.Int64Default(m["type_id"])
	b.CatId = db.Int64Default(m["cat_id"])
	b.Tag = db.StringDefault(m["tag"])
	b.Thumb = db.StringDefault(m["thumb"])
	b.IsRelevant = db.IntDefault(m["is_relevant"])
	b.IsJump = db.IntDefault(m["is_jump"])
	b.IsComment = db.IntDefault(m["is_comment"])
	b.IsRead = db.Int64Default(m["is_read"])
	b.Sort = db.Int64Default(m["sort"])
	b.Remark = db.StringDefault(m["remark"])
}

//Blog单个查询
func BlogGetOne(where string, q db.Query, args ...interface{}) *Blog {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlog))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	m := q.QueryRow(sqlBuilder.String(), args...).GetMap()
	if m == nil {
		return nil
	}
	ret := &Blog{}
	ret.SetMapValue(m)
	return ret
}

//Blog列表查询
func BlogGetList(where string, q db.Query, args ...interface{}) []*Blog {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlog))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Blog
	q.QueryRows(sqlBuilder.String(), args...).ForEach(func(m map[string]interface{}) bool {
		r := &Blog{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//Blog列表查询
func BlogGetPageList(where string, q db.Query, pg *db.PageObj, args ...interface{}) []*Blog {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(BlogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(q.Table(TableBlog))
	if where != "" {
		sqlBuilder.WriteString(" WHERE ")
		sqlBuilder.WriteString(where)
	}
	var ret []*Blog
	q.QueryWithPage(sqlBuilder.String(), pg, args...).ForEach(func(m map[string]interface{}) bool {
		r := &Blog{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
