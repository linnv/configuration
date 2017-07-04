// Package main provides ...
package demo

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

var (
	user      string
	pass      string
	prot      string
	addr      string
	dbname    string
	dsn       string
	netAddr   string
	available bool
	_db       *sql.DB
)

func init() {
	var err error
	_db, err = sql.Open("mysql", GetDNS())
	if err != nil {
		log.Fatalf("Error connecting: %s", err.Error())
	}
}

func GetDNS() string {
	if netAddr != "" {
		return fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
	}

	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}
	user = env("MYSQL_TEST_USER", "root")
	pass = env("MYSQL_TEST_PASS", "root")
	prot = env("MYSQL_TEST_PROT", "tcp")
	// addr = env("MYSQL_TEST_ADDR", "localhost:3306")
	addr = env("MYSQL_TEST_ADDR", "127.0.0.1:3306")
	dbname = env("MYSQL_TEST_DBNAME", "test")
	netAddr = fmt.Sprintf("%s(%s)", prot, addr)
	r := fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
	log.Printf("r: %+v\n", r)
	return r
}

func GetDB() *sql.DB {
	if _db == nil {
		var err error
		_db, err = sql.Open("mysql", GetDNS())
		if err != nil {
			log.Fatalf("Error connecting: %s", err.Error())
		}
	}
	return _db
}

func QueryDemo() {
	println("//<<-------------------------QueryDemo start-----------")
	start := time.Now()

	// fmt.Printf("GetDNS(): %+v\n", GetDNS())
	// db, err := sql.Open("mysql", GetDNS())
	// if err != nil {
	// 	log.Fatalf("Error connecting: %s", err.Error())
	// }
	// defer db.Close()
	db := GetDB()
	query := "SELECT * FROM table_test"
	// rows, err := db.Query(query, args...)
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
		return
	}
	for rows.Next() {
		var a, b, c string
		err = rows.Scan(&a, &b, &c)
		if err != nil {
			panic(err.Error())
			return
		}
		fmt.Printf("a: %+v\n", a)
		fmt.Printf("b: %+v\n", b)
		fmt.Printf("c: %+v\n", c)
		// println("------ pause \n========================\n")
	}
	// return rows
	fmt.Printf("%v as %v millisecons\n", time.Since(start), time.Since(start).Nanoseconds()/1000000)
	println("//---------------------------QueryDemo end----------->>")
}

func checkErr(err error) {
	if err != nil {
		log.Printf("err: %+v\n", err.Error())
		panic(err.Error())
		return
	}
}

func UpdateDemo() {
	println("//<<-------------------------UpdateDemo start-----------")
	start := time.Now()
	fmt.Printf("GetDNS(): %+v\n", GetDNS())
	db, err := sql.Open("mysql", GetDNS())
	if err != nil {
		log.Fatalf("Error connecting: %s", err.Error())
	}
	defer db.Close()
	//更新数据
	stmt, err := db.Prepare("update table_test set a=? where a=?")
	checkErr(err)

	res, err := stmt.Exec("JialinWu", "astaxie")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	fmt.Printf("UpdateDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------UpdateDemo end----------->>")
}

func InsertDemo() {
	println("//<<-------------------------InsertDemo start-----------")
	start := time.Now()

	// fmt.Printf("GetDNS(): %+v\n", GetDNS())
	// db, err := sql.Open("mysql", GetDNS())
	// if err != nil {
	// 	log.Fatalf("Error connecting: %s", err.Error())
	// }
	// defer db.Close()

	db := GetDB()
	// stmt, err := db.Prepare("INSERT table_test(a,b,c) values(?,?,?)")
	stmt, err := db.Prepare(`INSERT table_test(a,b,c) values(?,?,?)`)
	checkErr(err)
	// res, err := stmt.Exec("jialinwu", "研发部门", time.Now().String())
	res, err := stmt.Exec(1, "研发部门", "x33")
	checkErr(err)

	id, err := res.LastInsertId()
	fmt.Printf("id: %+v\n", id)
	checkErr(err)
	fmt.Printf("%v as %v millisecons\n", time.Since(start), time.Since(start).Nanoseconds()/1000000)
	println("//---------------------------InsertDemo end----------->>")
}

func TransactionDemo() {
	println("//<<-------------------------TransactionDemo start-----------")
	start := time.Now()
	db, err := sql.Open("mysql", GetDNS())
	if err != nil {
		log.Fatalf("Error connecting: %s", err.Error())
	}
	defer db.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	// checkErr(err)
	// //更新数据
	const TABLE_NAME = "customer"
	go func() {
		defer wg.Done()
		dbSecond, err := sql.Open("mysql", GetDNS())
		if err != nil {
			log.Fatalf("Error connecting: %s", err.Error())
		}
		defer dbSecond.Close()

		log.Println("cat sleep 1 second to let dog go first: works")
		time.Sleep(time.Second * 1)
		tx, err := dbSecond.Begin()
		log.Println("cat transaction starts: works all the transaction has no sleeping ,querying")
		checkErr(err)
		rows, err := tx.Query("select * from " + TABLE_NAME + " where a=10 for update")
		s, err := rows.Columns()
		checkErr(err)
		rets := make([]Row, 0, len(s))
		for rows.Next() {
			row := make(Row, len(s))
			for i, _ := range row {
				row[i] = new([]byte)
			}

			rows.Scan(row...)
			rets = append(rets, row)
		}
		// you must check the value before updating
		for _, v := range rets {
			for i, _ := range v {
				fmt.Printf("cat tx  %+v: %+v\n", i, v.Str(i))
			}
		}
		log.Println("cat query for update done,holding read and write lock")

		stmt, err := tx.Prepare("update customer set b=? where a=?")
		if err != nil {
			panic(err.Error())
		}
		strt := time.Now().String()
		updateStr := strt[len(strt)-17:]
		res, err := stmt.Exec(strt[len(strt)-15:], 10)
		if err != nil {
			panic(err.Error())
		}
		_, err = res.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		// log.Println("cat sleep for 4 seconds befor commiting: works")
		// time.Sleep(4 * time.Second)
		log.Println("sleep befor rollback: works")
		time.Sleep(3 * time.Second)
		err = tx.Commit()
		if err != nil {
			panic(err.Error())
		}
		log.Println("cat commited: works", strt)
	}()

	tx, err := db.Begin()
	log.Println("dog transaction starts: works")
	checkErr(err)
	rows, err := tx.Query("select * from " + TABLE_NAME + " where a=10 for update")
	s, err := rows.Columns()
	checkErr(err)
	rets := make([]Row, 0, len(s))
	for rows.Next() {
		row := make(Row, len(s))
		for i, _ := range row {
			row[i] = new([]byte)
		}

		rows.Scan(row...)
		rets = append(rets, row)
	}
	for _, v := range rets {
		for i, _ := range v {
			fmt.Printf("dog tx  %+v: %+v\n", i, v.Str(i))
		}
	}
	log.Println("dog query for update done (holding read and writing lock);sleep 6 seconds to let cat go")
	time.Sleep(time.Second * 6)
	stmt, err := tx.Prepare("update customer set b=? where a=?")
	checkErr(err)

	// log.Println("query only dog sleep 3 seconds befor rollbacking: works")
	// time.Sleep(3 * time.Second)
	// log.Println("queyr only done:dog sleep 3 seconds befor rollbacking: works")

	strt := time.Now().String()
	updateStr := strt[len(strt)-17:]
	res, err := stmt.Exec(updateStr, 10)
	if err != nil {
		panic(err.Error())
	}
	_, err = res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	log.Println("dog TransactionDemo exec: works, update b to ", updateStr)
	log.Println("dog sleep 3 seconds befor rollbacking: works")
	time.Sleep(3 * time.Second)
	log.Println("done:dog sleep 3 seconds befor rollbacking: works")

	// rows, err = tx.Query("select * from " + TABLE_NAME + " where a=10 for update")
	// s, err = rows.Columns()
	// checkErr(err)
	// rets = make([]Row, 0, len(s))
	// for rows.Next() {
	// 	row := make(Row, len(s))
	// 	for i, _ := range row {
	// 		row[i] = new([]byte)
	// 	}
	//
	// 	rows.Scan(row...)
	// 	rets = append(rets, row)
	// }
	// for _, v := range rets {
	// 	for i, _ := range v {
	// 		fmt.Printf("dog2 tx  %+v: %+v\n", i, v.Str(i))
	// 	}
	// }

	err = tx.Commit()
	// err = tx.Rollback()
	if err != nil {
		panic(err.Error())
	}
	log.Println("dog commited or rollbacked( free reading and writing lock): works ;now cat can commit or Rollback", strt)
	log.Println("write lock is FIFF: works")

	fmt.Printf("\nTransactionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------TransactionDemo end----------->>")
	wg.Wait()
}
