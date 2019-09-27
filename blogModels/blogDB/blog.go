package blogDB

import (
	"bytes"
	"strings"

	"github.com/kinwyb/go/db/mysql"

	"github.com/kinwyb/go/err1"

	"github.com/kinwyb/blog/blogModels/blogBeans"
	"github.com/kinwyb/blog/blogModels/blogBeans/customer"
	"github.com/kinwyb/blog/blogModels/blogBeans/dbBeans"
	"github.com/kinwyb/go/db"
)

//博客列表
func BlogList(req *customer.BlogListReq, pg *db.PageObj, ctx *blogBeans.Context) []*dbBeans.Blog {
	defer ctx.Start("db.BlogList").Finish()
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(dbBeans.BlogColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(ctx.Query.Table(dbBeans.TableBlog))
	sqlBuilder.WriteString(" WHERE 1=1 ")
	var args []interface{}
	if req.Status != 0 {
		sqlBuilder.WriteString(" AND status = ? ")
		args = append(args, req.Status)
	}
	if req.IsOpen != "" {
		sqlBuilder.WriteString(" AND is_open = ? ")
		args = append(args, req.IsOpen)
	}
	if req.Tp != 0 {
		sqlBuilder.WriteString(" AND type = ? ")
		args = append(args, req.Tp)
	}
	if req.Search != "" {
		sqlBuilder.WriteString(" AND title LIKE ? ")
		args = append(args, "%"+req.Search+"%")
	}
	sqlBuilder.WriteString(" ORDER BY sort,blog_id DESC")
	var ret []*dbBeans.Blog
	ctx.Query.QueryWithPage(sqlBuilder.String(), pg, args...).Error(func(err err1.Error) {
		log.Error("博客列表查询失败:%s", err.Error())
	}).ForEach(func(m map[string]interface{}) bool {
		r := &dbBeans.Blog{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}

//博客状态
func BlogStatistics(blogIDs []string, ctx *blogBeans.Context) []*dbBeans.BlogStatistics {
	defer ctx.Start("db.BlogStatistics").Finish()
	sqlBuilder := bytes.NewBufferString("SELECT ")
	sqlBuilder.WriteString(dbBeans.BlogStatisticsColumns)
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(ctx.Query.Table(dbBeans.TableBlogStatistics))
	sqlBuilder, args := mysql.WhereIN(" WHERE blog_id", strings.Join(blogIDs, ","), nil, sqlBuilder)
	var ret []*dbBeans.BlogStatistics
	ctx.Query.QueryRows(sqlBuilder.String(), args...).Error(func(err err1.Error) {
		log.Error("博客状态查询失败:%s", err.Error())
	}).ForEach(func(m map[string]interface{}) bool {
		r := &dbBeans.BlogStatistics{}
		r.SetMapValue(m)
		ret = append(ret, r)
		return true
	})
	return ret
}
