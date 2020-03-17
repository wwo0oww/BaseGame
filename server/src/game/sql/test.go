package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"proto/net"
	"time"
)
type User struct {
	ID int  `db:"id"`
	Account   string          `db:"account"`
	Pwd string `db:"pwd"`
	Time  int            `db:"age"`
}

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "arrival"
)
func Test(){
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数
	//InsertData1(DB)
	//queryOne(DB)
}

func queryOne(DB *sql.DB){
	//fmt.Println("query times:",i)
	user := new(User)
	row := DB.QueryRow("SELECT account FROM user where account=? and pwd=?","YDZ",123)
	fmt.Print(row)
	if err :=row.Scan(&user.ID,&user.Account,&user.Pwd,&user.Time); err != nil{
		fmt.Printf("scan failed, err:%v",err)
		//DB.Exec("INSERT INTO user(account,pwd,time) VALUES(?,?,?)","111","111",time.Now().Unix() )
		return
	}
	fmt.Println(*user)
}

//插入数据
func InsertData1(DB *sql.DB){
	tos := net.MLoginTos{Name:"1112",Pwd:"111"}
	_,err := SqlWrapper.GetDrv().Exec("INSERT INTO user (account,pwd) VALUES(? , ?)", tos.Name, tos.Pwd)
	DB.Query("use arrival")
	_,err1 := DB.Exec("INSERT INTO user (account,pwd) VALUES(? , ?)", tos.Name, tos.Pwd)
	if err != nil{
		fmt.Printf("Insert failed,err:%v",err)

	}
	if err1 != nil{
		fmt.Printf("111 Insert failed,err1:%v",err1)
	}
	//lastInsertID,err := result.LastInsertId()
	//if err != nil {
	//	fmt.Printf("Get lastInsertID failed,err:%v",err)
	//	return
	//}
	//fmt.Println("LastInsertID:",lastInsertID)
	//rowsaffected,err := result.RowsAffected()
	//if err != nil {
	//	fmt.Printf("Get RowsAffected failed,err:%v",err)
	//	return
	//}
	//fmt.Println("RowsAffected:",rowsaffected)
}
//
////更新数据
//func updateData(DB *sql.DB){
//	result,err := DB.Exec("UPDATE users set age=? where id=?","30",3)
//	if err != nil{
//		fmt.Printf("Insert failed,err:%v",err)
//		return
//	}
//	rowsaffected,err := result.RowsAffected()
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v",err)
//		return
//	}
//	fmt.Println("RowsAffected:",rowsaffected)
//}
//
////删除数据
//func deleteData(DB *sql.DB){
//	result,err := DB.Exec("delete from users where id=?",1)
//	if err != nil{
//		fmt.Printf("Insert failed,err:%v",err)
//		return
//	}
//	rowsaffected,err := result.RowsAffected()
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v",err)
//		return
//	}
//	fmt.Println("RowsAffected:",rowsaffected)
//}