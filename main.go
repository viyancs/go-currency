/**
 *  main file for start running rest api
 *
 *  created by @viyancs, 25/07/2018
 *  make sure always create simple code and clean code,
 *  always give mark comment for easy maintenancen
 */
package main

import (
    "log"
    "github.com/go-gem/gem"
    "github.com/astaxie/beego/orm"
    "./models"
    "github.com/davecgh/go-spew/spew"
    "time"
    "strconv"

)

const BAD_REQUEST = "Bad Request"
const DATA_NOT_FOUND = "Data Not Found"
const DATE_INVALID = "Date Invalid"
const RATE_INVALID = "Rate Invalid"
const ERROR_WHEN_SAVE_TO_DB = "Error When Save to DB"
const ERROR_WHEN_DELETE_TO_DB = "Error When Delete to DB"
const SUCCESS_INSERT = "success insert"
const SUCCESS_DELETE = "success delete"
const C_ERR = 400;
const C_OK = 200;

type Jsonformat struct {
    Code int    `json:"code"`
    Results []orm.Params    `json:"results"`
    Msg string    `json:"msg"`
    Total int    `json:"total"`
}

func index(ctx *gem.Context) {
    ctx.HTML(C_OK, "Hi Lets Play ")
}

// get all exchange
func exchange(ctx *gem.Context) {
    rows := models.GetDailyexchange();
    datas := Jsonformat{Code: C_OK, Results: rows, Total:len(rows)}
    spew.Dump(datas) //trace dumb variable
    ctx.JSON(C_OK, datas)
    return
}

// get exchange by date
func exchangeByDate(ctx *gem.Context) {
    date, err := gem.String(ctx.UserValue("date"))
    if err != nil {
        datas := Jsonformat{Code: C_ERR, Msg: BAD_REQUEST}
        ctx.JSON(C_ERR, datas)
        return
    }
    rows := models.GetDailyexchange(date);
    if (rows == nil) {
        datas := Jsonformat{Code: C_OK, Msg: DATA_NOT_FOUND}
        ctx.JSON(C_OK, datas)
        return
    }
    datas := Jsonformat{Code: C_OK, Results: rows, Total:len(rows)}
    ctx.JSON(C_OK, datas)
}

//new exchange
func newExchange(ctx *gem.Context) {
    ctx.Request.ParseForm()

    //parse date
    layout := "2006-01-02"
    date, errdate := time.Parse(layout, ctx.Request.FormValue("date"))
    rate, errrate := strconv.ParseFloat(ctx.Request.FormValue("rate"), 64)

    if errdate != nil {
        datas := Jsonformat{Code: C_ERR, Msg: DATE_INVALID}
        ctx.JSON(C_ERR, datas)
        return
    }

    if errrate != nil {
        datas := Jsonformat{Code: C_ERR, Msg: RATE_INVALID}
        ctx.JSON(C_ERR, datas)
        return
    }

    var exc models.Dailyexchange
    exc.C_from = ctx.Request.FormValue("from")
    exc.C_to = ctx.Request.FormValue("to")
    exc.Rate = rate
    exc.Date = date

    //this should be add more validation security first before go to database
    res := models.SaveDailyexchange(exc)
    if res == 0 {
        datas := Jsonformat{Code: C_ERR, Msg: ERROR_WHEN_SAVE_TO_DB}
        ctx.JSON(C_ERR, datas)
        return
    }
    spew.Dump(res) //trace dumb variable
    id := strconv.Itoa(int(res))
    datas := Jsonformat{Code: C_OK, Msg: id }
    ctx.JSON(C_OK, datas)
}

//new track
func newTrack(ctx *gem.Context) {
    ctx.Request.ParseForm()
    var exc models.Dailyexchange
    exc.C_from = ctx.Request.FormValue("from")
    exc.C_to = ctx.Request.FormValue("to")

    //this should be add more validation security first before go to database
    res := models.SaveDailyexchange(exc)
    if res == 0 {
        datas := Jsonformat{Code: C_ERR, Msg: ERROR_WHEN_SAVE_TO_DB}
        ctx.JSON(C_ERR, datas)
        return
    }
    spew.Dump(res) //trace dumb variable
    id := strconv.Itoa(int(res))
    datas := Jsonformat{Code: C_OK, Msg: id }
    ctx.JSON(C_OK, datas)
}

//delete exchange
func delExchange(ctx *gem.Context) {
    ctx.Request.ParseForm()

    form := ctx.Request.FormValue("from")
    to := ctx.Request.FormValue("to")

    //this should be add more validation security first before go to database
    res := models.DelDailyexchange(form,to)
    if res == 0 {
        datas := Jsonformat{Code: C_ERR, Msg: ERROR_WHEN_DELETE_TO_DB}
        ctx.JSON(C_ERR, datas)
        return
    }

    datas := Jsonformat{Code: C_OK, Msg: SUCCESS_DELETE }
    ctx.JSON(C_OK, datas)
}

func main() {
    // Create server.
    srv := gem.New(":8282")

    // Create router.
    router := gem.NewRouter()
    // Register handler
    router.GET("/", index)
    router.GET("/exchange", exchange)
    router.GET("/exchange/:date", exchangeByDate)
    router.POST("/exchange", newExchange)
    router.POST("/track", newTrack)
    router.PUT("/exchange", delExchange)

    // Start server.
    log.Println(srv.ListenAndServe(router.Handler()))
}
