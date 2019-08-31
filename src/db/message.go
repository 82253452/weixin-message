package db

import "log"

type Message struct {
	Robotid       string `form:"robotid" db:"robotid"`
	Msgid         string `form:"msgid" db:"msgid"`
	Gid           string `form:"gid" db:"gid"`
	Gusername     string `form:"gusername" db:"gusername"`
	Gname         string `form:"gname" db:"gname"`
	Mid           string `form:"mid" db:"mid"`
	Nickname      string `form:"nickname" db:"nickname"`
	Displayname   string `form:"displayname" db:"displayname"`
	Gadmin        string `form:"gadmin" db:"gadmin"`
	Skw           string `form:"skw" db:"skw"`
	Content       string `form:"content" db:"content"`
	Atlist        string `form:"atlist" db:"atlist"`
	Robotnickname string `form:"robotnickname" db:"robotnickname"`
	Atmod         string `form:"atmod" db:"atmod"`
}

func (row *Message) Save() {
	tx := DB.MustBegin()
	stmt, _ := tx.Prepare("insert into message(robotid,msgid,gid,gusername,gname,mid,nickname,displayname,gadmin,skw,content,atlist,robotnickname,atmod) value (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	_, err := stmt.Exec(
		row.Robotid, row.Msgid, row.Gid, row.Gusername, row.Gname, row.Mid,
		row.Nickname, row.Displayname, row.Gadmin, row.Skw, row.Content,
		row.Atlist, row.Robotnickname, row.Atmod)
	if err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	stmt.Close()
}
