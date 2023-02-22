package main

import (
	"dous/biz/dal/mysql"
	"dous/biz/model/query"
)

func init() {
	print("init函数已经被执行过了")
	mysql.Init()
	query.SetDefault(mysql.DB)
}
