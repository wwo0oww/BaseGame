package sql

import (
	"database/sql"
	"fmt"
	"lib/cellnet"
	"lib/cellnet/peer"
	"lib/cellnet/peer/mysql"
	"sync/atomic"
	"time"
)

var (
	SqlWrapper *mysql.Wrapper
)

var Running int32 = 0

func Startsql() {
	// 从地址中选择mysql数据库，这里选择mysql系统库
	p := peer.NewGenericPeer("mysql.Connector", "mysqldemo", "root:123456@(localhost:3306)/arrival", nil)
	p.(cellnet.MySQLConnector).SetConnectionCount(3)

	// 阻塞
	p.Start()

	op := p.(cellnet.MySQLOperator)

	op.Operate(func(rawClient interface{}) interface{} {

		client := rawClient.(*sql.DB)
		// 查找默认用户
		SqlWrapper = mysql.NewWrapper(client)
		atomic.StoreInt32(&Running, 1)
		return nil
	})
}

func WaitForRunning() {
	for {
		if Running == 1 {
			break
		}
		time.Sleep(100000)
	}
}

//查询单行
func QueryOne(query string, args ...interface{}) *mysql.Wrapper {
	Wrap := SqlWrapper.QueryOne(query, args...)
	return Wrap
}

//插入数据
func InsertData(query string, args ...interface{}) (succ bool, err error) {

	succ = true
	Wrap := SqlWrapper.Execute(query, args...)
	err = Wrap.Err
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		succ = false
		return
	}
	return
}

func Instance() *mysql.Wrapper {
	return SqlWrapper
}
