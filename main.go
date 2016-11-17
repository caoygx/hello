package main

import (
    _ "hello/routers"
//    "github.com/astaxie/beego"
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:@/orm_test?charset=utf8")
}

func main() {
    o := orm.NewOrm()
    o.Using("default") // 默认使用 default，你可以指定为其他数据库

    user := User{Id:1}
    o.Read(&user)
    fmt.Println(user.Id, user.Name)

//    profile := new(Profile)
//    profile.Age = 30

    //user := new(User)

//    user.Profile = profile
//    user.Name = "slene"

//    fmt.Println(o.Insert(profile))
  //  fmt.Println(o.Insert(user))
}
