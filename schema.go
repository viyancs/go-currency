package main
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
    Rate float64
    Date time.Time
}

func init(){
    orm.RegisterModel(new(Dailyexchange))
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:@/test_go?charset=utf8", 30)
}

func main() {
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
    }
    successNums, err := o.InsertMulti(100, de)
    if err == nil {
        t := fmt.Sprintf("Insert successfully = %d ",successNums)
        fmt.Println(t)
    }
}
