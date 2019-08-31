package db

import (
	"fmt"
	"log"
)

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

type MessageDto struct {
	Num      int    `db:"num"`
	Nickname string `db:"nickname"`
	Ctime    string `db:"ctime"`
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
func SelectNameLine(name string) []MessageDto {
	messages := []MessageDto{}
	err := DB.Select(&messages, `select ctime,nickname from message where nickname=? order by ctime desc`, name)
	if err != nil {
		fmt.Println(err)
	}
	return messages
}
func SelectAllNames() []MessageDto {
	messages := []MessageDto{}
	_ = DB.Select(&messages, `select count(nickname) as num, nickname 
										from message
										group by nickname
										order by count(nickname) desc`)
	return messages
}

func Test() {
	place := MessageDto{}
	rows, err := DB.Queryx("SELECT * FROM message")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}
}
