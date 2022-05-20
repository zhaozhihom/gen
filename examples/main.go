package main

import (
	"context"
	"fmt"

	"github.com/zhaozhihom/gen/examples/biz"
	"github.com/zhaozhihom/gen/examples/conf"
	"github.com/zhaozhihom/gen/examples/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	biz.Query(context.Background())
}
