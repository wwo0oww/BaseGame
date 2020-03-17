package db

import "game/sql"

func GetMaxPlayerNum() (MaxNum int32) {
	MaxNum = 0
	sql.QueryOne("select max(id) from user").OneScan(&MaxNum)
	return
}

func GetAccountInfo(Name string, Pwd string) (Account, UserID string) {
	sql.QueryOne("SELECT account,user_id FROM user where account=? and pwd=?", Name, Pwd).OneScan(&Account, &UserID)
	return
}

func GetAccountInfoByName(Name string) (Account string) {
	sql.QueryOne("SELECT account FROM user where account=?", Name).OneScan(&Account)
	return
}

func InsertPlayer(Name, Pwd, UserID string) error {
	_, err := sql.InsertData("INSERT INTO user (account,pwd,user_id) VALUES(?, ?, ?)", Name, Pwd, UserID)
	return err
}
