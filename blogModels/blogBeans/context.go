package blogBeans

import (
	"github.com/kinwyb/go"
	"github.com/kinwyb/go/db"
)

//上下文
type Context struct {
	*heldiamgo.Context
	Token string   //token
	Data map[interface{}]interface{}
	Page *db.PageObj
	Template string
	Query db.Query //数据库连接
}

//ChildOf
func (t *Context) Child() *Context {
	heldiamgo.ContextChild(t)
	return t
}

//FollowsFrom
func (t *Context) Follows() *Context {
	heldiamgo.ContextFollows(t)
	return t
}

func (t *Context) QueryTransaction(tx db.Query) *Context {
	ret := &Context{
		Token:   t.Token,
		Query:   t.Query,
		Context: t.Context,
	}
	ret.Query = tx
	return ret
}

//初始化上下文
func NewContext(companyID int, operationName string) *Context {
	ctx := &Context{}
	if operationName != "" && heldiamgo.Tracing {
		ctx.Context = heldiamgo.NewContext(operationName)
	}
	return ctx
}

//初始化上下文
func NewContextWithTracing(tracingSpan heldiamgo.TracingSpan) *Context {
	if !heldiamgo.Tracing {
		tracingSpan = nil
	}
	return &Context{
		Context: heldiamgo.NewContextWithTracing(tracingSpan),
	}
}
