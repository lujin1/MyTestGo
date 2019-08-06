package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"time"
)


var db *xorm.Engine

type User struct {
	Id int64
	Name string
	Age int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}



func init() {
	var err error
	db, err = xorm.NewEngine("postgres", "user=postgres password=123456 dbname=test host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
	if err = db.Sync2(new(User)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func AddUser(name string, age int) error {
	// 对未存在记录进行插入
	_, err := db.Insert(&User{Name: name, Age: age})
	return err
}


func in(c *gin.Context)  {
	if err := AddUser("lujin", 18); err != nil {
		log.Fatalf("AddUser error: %v\n", err)
	}

	c.JSON(200, gin.H{
		"message": "hhh",
	})
 }

func main()  {

	r := gin.Default()
	r.GET("/", in)
	r.Run()

}



