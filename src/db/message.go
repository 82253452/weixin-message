package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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

type MessageGroupDto struct {
	WordNumm int    `db:"wordNumm"`
	TotalNum int    `db:"totalNum"`
	ImgNum   int    `db:"imgNum"`
	TextNum  int    `db:"textNum"`
	VideoNum int    `db:"videoNum"`
	Gid      string `db:"gid"`
	Gname    string `db:"gname"`
	LinkNum  string `db:"linkNum"`
	Nickname string `db:"nickname"`
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
										order by count(nickname) desc limit 10`)
	return messages
}

func SelectAllGroups(names []string) []MessageGroupDto {
	messages := []MessageGroupDto{}
	//	querys := `select count(*)                              as wordNumm,
	//       gid,
	//       gname,
	//       s.nickname,
	//       (select count(*)
	//        from message a
	//        where a.gid = s.gid
	//          and a.mid = s.mid)                 as totalNum,
	//       (select count(*)
	//        from message a
	//        where a.gid = s.gid
	//          and a.content like '%<img%'
	//          and a.mid = s.mid)                 as imgNum,
	//       (select count(*)
	//        from message a
	//        where a.gid = s.gid
	//          and a.content like '%<video%'
	//          and a.mid = s.mid)                 as videoNum,
	//       (select count(*)
	//        from message a
	//        where a.gid = s.gid
	//          and a.mid = s.mid
	//          and a.content not like '%<img%'
	//          and a.content not like '%<video%'
	//          and a.content not like '%http%') as textNum,
	//(select count(*)
	//        from message a
	//        where a.gid = s.gid
	//          and a.mid = s.mid
	//          and a.content like '%http%') as linkNum
	//from message s
	//where nickname in (?)
	//  and content like '%<img%'
	//and TO_DAYS(ctime) = TO_DAYS(NOW())
	//group by gid, gname, mid,nickname`
	querys := `
		select count(*)                            as wordNumm,
       gid,
       gname,
       (select count(*)
        from message a
        where a.gid = s.gid
          and TO_DAYS(ctime) = TO_DAYS(NOW())
          and nickname in (?)
       )                                   as totalNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and a.content like '%<img%'
          and TO_DAYS(ctime) = TO_DAYS(NOW())
            and nickname in (?)
       )                                   as imgNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and a.content like '%<video%'
          and TO_DAYS(ctime) = TO_DAYS(NOW())
            and nickname in (?)
       )                                   as videoNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and a.content not like '%<img%'
          and a.content not like '%<video%'
          and TO_DAYS(ctime) = TO_DAYS(NOW())
          and a.content not like '%http%') as textNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and TO_DAYS(ctime) = TO_DAYS(NOW())
           and nickname in (?)
          and a.content like '%http%')     as linkNum
from message s
where 1=1 
  and 	nickname in (?)
  and TO_DAYS(ctime) = TO_DAYS(NOW())
group by gid, gname
`
	//nickname in (?)
	query, args, err := sqlx.In(querys, names, names, names, names, names)
	if err != nil {
		fmt.Println(err)
	}
	error := DB.Select(&messages, query, args...)
	if error != nil {
		fmt.Println(error)
	}
	return messages
}

func SelectAllGroupsNew(names []string) []MessageGroupDto {
	querys := `
		select count(*)                            as wordNumm,
       gid,
       gname,
       (select count(*)
        from message a
        where a.gid = s.gid
          and TO_DAYS(ctime) = TO_DAYS(NOW())
       )                                   as totalNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and a.content like '%<img%'
          and TO_DAYS(ctime) = TO_DAYS(NOW())
       )                                   as imgNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and a.content like '%<video%'
          and TO_DAYS(ctime) = TO_DAYS(NOW())
       )                                   as videoNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and a.content not like '%<img%'
          and a.content not like '%<video%'
          and TO_DAYS(ctime) = TO_DAYS(NOW())
          and a.content not like '%http%') as textNum,
       (select count(*)
        from message a
        where a.gid = s.gid
          and TO_DAYS(ctime) = TO_DAYS(NOW())
          and a.content like '%http%')     as linkNum
from message s
where 1=1 
  and TO_DAYS(ctime) = TO_DAYS(NOW())
group by gid, gname
`
	messages := []MessageGroupDto{}
	error := DB.Select(&messages, querys)
	if error != nil {
		fmt.Println(error)
	}
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
