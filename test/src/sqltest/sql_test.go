/**
 *  this file usage for testing reliable databse when big data is come
 *
 *  created by @viyancs, 25/07/2018
 *  make sure always create simple code and clean code,
 *  always give mark comment for easy maintenancen
 */
package sqltest

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
)

type Tester interface {
	RunTest(*testing.T, func(params))
}

var (
	myMysql Tester = &myMysqlDB{}
)

const TablePrefix = ""

type myMysqlDB struct {
	once    sync.Once // guards init of running
	running bool      // whether port 3306 is listening
}

func (m *myMysqlDB) Running() bool {
	m.once.Do(func() {
		c, err := net.Dial("tcp", "localhost:3306")
		if err == nil {
			m.running = true
			c.Close()
		}
	})
	return m.running
}


type params struct {
	dbType Tester
	*testing.T
	*sql.DB
}

func (t params) mustExec(sql string, args ...interface{}) sql.Result {
	res, err := t.DB.Exec(sql, args...)
	if err != nil {
		t.Fatalf("Error running %q: %v", sql, err)
	}
	return res
}

var qrx = regexp.MustCompile(`\?`)

func (t params) q(sql string) string {
	var pref string
	switch t.dbType {
	default:
		return sql
	}
	n := 0
	return qrx.ReplaceAllStringFunc(sql, func(string) string {
		n++
		return pref + strconv.Itoa(n)
	})
}

func (m *myMysqlDB) RunTest(t *testing.T, fn func(params)) {
	if !m.Running() {
		t.Logf("skipping test; no MySQL running on localhost:3306")
		return
	}
	user := os.Getenv("GOSQLTEST_MYSQL_USER")
	if user == "" {
		user = "root"
	}
	pass, ok := getenvOk("GOSQLTEST_MYSQL_PASS")
	if !ok {
		pass = ""
	}
	dbName := "test_go"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user,pass,dbName))
	if err != nil {
		t.Fatalf("error connecting: %v", err)
	}

	params := params{myMysql, t, db}

	// Drop all tables in the test database.
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		t.Fatalf("failed to enumerate tables: %v", err)
	}
	for rows.Next() {
		var table string
		if rows.Scan(&table) == nil &&
			strings.HasPrefix(strings.ToLower(table), strings.ToLower(TablePrefix)) {
			params.mustExec("DROP TABLE " + table)
		}
	}

	fn(params)
}

func TestManyQueryRow_MyMySQL(t *testing.T) { myMysql.RunTest(t, testManyQueryRow) }

func testManyQueryRow(t params) {
	if testing.Short() {
		t.Logf("skipping in short mode")
		return
	}
	t.mustExec("create table " + TablePrefix + "foo (id integer primary key, name varchar(50))")
	t.mustExec(t.q("insert into "+TablePrefix+"foo (id, name) values(?,?)"), 1, "bob")
	var name string
	for i := 0; i < 10000; i++ {
		err := t.QueryRow(t.q("select name from "+TablePrefix+"foo where id = ?"), 1).Scan(&name)
		if err != nil || name != "bob" {
			t.Fatalf("on query %d: err=%v, name=%q", i, err, name)
		}
	}
}


func getenvOk(k string) (v string, ok bool) {
	v = os.Getenv(k)
	if v != "" {
		return v, true
	}
	keq := k + "="
	for _, kv := range os.Environ() {
		if kv == keq {
			return "", true
		}
	}
	return "", false
}
