package main

import (
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
	"database/sql"
	"fmt"
)

type Userinfo struct {
	Uid        int `PK`
	Username   string
	Departname string
	Created    time.Time
}

func main() {
	//初始化
	db, err := sql.Open("mymysql", "test/root/topsec") // dataSourceName简洁
	//db, err := sql.Open("mymysql", "root:topsec@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	//两个参数， first标准接口的db， second参数是使用的数据库引擎，
	orm := beedb.New(db)

	// 插入数据 , 两种方式，第一种点的形式， 第二种是map的形式
	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	//orm.Save(saveone) 可以吗， 如果不可以， 为什么
	orm.Save(&saveone)

	// 更新数据
	saveone.Username = "Update Username"
	saveone.Departname = "Update Department"
	saveone.Created = time.Now()
	orm.Save(&saveone) // 为什么要用指针

	// 查询数据  TODO  未调试成功
	//主键
	var user Userinfo
	a := orm.Where("uid=?", 42).Find(&user)
	fmt.Println(a)
	//非主键
	var ser3 Userinfo
	b := orm.Where("username = ?", "zhangzhiyuan").Find(&ser3)
	fmt.Println(b)
	// 复杂的条件
	var user4 Userinfo
	c := orm.Where("username = ? and uid < ?", "zhangzhiyuan", 10).Find(&user4)
	fmt.Println(c)

	//删除
	orm.SetTable("userinfo").Where("uid>?", 3).DeleteRow()
}
