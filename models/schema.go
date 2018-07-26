/**
 *  main file for schema comand line this file using beego orm and can be use for managing table
 *
 *  created by @viyancs, 25/07/2018
 *  make sure always create simple code and clean code,
 *  always give mark comment for easy maintenancen
 */
package models
import (
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql" // import your required driver
)

type Dailyexchange struct {
    Id   int    `orm:"column(id)"`
    C_from string   `orm:"size(3)"`
    C_to string   `orm:"size(3)"`
    Rate float64    `orm:"null"`
    Date time.Time  `orm:"null"`
}


func init(){
    orm.RegisterModel(new(Dailyexchange))
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:@/test_go?charset=utf8", 30)
}

func GetDailyexchange(params ...string) []orm.Params{
    o := orm.NewOrm()
    var maps []orm.Params

    date := "";
    if len(params) > 0 {
        date = params[0] //set date srting in 0 index
    }

    if (len(date) > 0) {
        query := "select id,c_from,c_to,date,IF(rate = '0', 'insufficient data', rate) as rate,IFNULL(ROUND(AVG(CASE WHEN date between (date - INTERVAL 2 DAY) and date then rate else 0 end), 2),'insufficient data') AS 'last 2 days' from dailyexchange where date = ? group by c_from,c_to"
        o.Raw(query,date).Values(&maps)
        return maps
    }

    o.Raw("select * from dailyexchange").Values(&maps)
    return maps

}

func SaveDailyexchange(d Dailyexchange) int64{
    o := orm.NewOrm()
    r,err := o.Insert(&d)
    if err == nil {
        return r
    }
    return 0
}

func DelDailyexchange(from string, to string) int64{
    o := orm.NewOrm()
    _, err := o.Raw("delete FROM dailyexchange where c_from = ? and c_to = ?").SetArgs(from, to).Exec()
    if err == nil {
        return 1
    }
    return 0

}

//this cmd for cli comand
func Cmd() {
    if (len(os.Args) < 2) {
        content := `schema command usage:

        sample             - importing sample data
        orm syncdb -force  - syncdb model to table by droping table first
        orm syncdb         - syncdb model to table
        [options] -v       - verbose output
    `
        fmt.Println(content)
        os.Exit(2)

    }
    if os.Args[1] == "sample" {
		sample()
	}
    orm.RunCommand()
}

func sample() {
    layout := "2006-01-02"
    twofour, err := time.Parse(layout, "2018-07-24")
    twofive, err := time.Parse(layout, "2018-07-25")
    o := orm.NewOrm()
    de := []Dailyexchange{
        {C_from: "USD",C_to:"IDR",Rate:15000,Date:twofour},
        {C_from: "USD",C_to:"EUR",Rate:0.854579,Date:twofour},
        {C_from: "GBP",C_to:"EUR",Rate:1.12142,Date:twofour},
        {C_from: "GBP",C_to:"INR",Rate:90.4866,Date:twofive},
        {C_from: "USD",C_to:"IDR",Rate:14750,Date:twofive},
        {C_from: "USD",C_to:"GBP",Rate:0.761913,Date:twofive},
        {C_from: "USD",C_to:"JPN",Date:twofive},
        {C_from: "JPN",C_to:"GBP",Date:twofive},
    }
    successNums, err := o.InsertMulti(100, de)
    if err == nil {
        t := fmt.Sprintf("Insert successfully = %d ",successNums)
        fmt.Println(t)
    }
}
