// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	"github.com/AliyevH/mysqlToExcel/configs"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {

// 	config := configs.Config

// 	db, sql_err := sql.Open(config.DB, config.DSN)
// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)

// 	defer db.Close()

// 	if sql_err != nil {
// 		panic(sql_err.Error())
// 	}

// 	// var result map[string]interface{}

// 	rows, _ := db.Query("select * from city_city")

// 	columns, _ := rows.Columns()
// 	count := len(columns)

// 	values := make([]interface{}, count)
// 	valuePtrs := make([]interface{}, count)

// 	finalResult := map[int]map[string]string{}
// 	// map[0:map[created_at:2018-07-10 17:20:35.000000 description: id:88 name:Gəncə region_id:11 status:1 updated_at:2020-10-08 08:45:30.075937]
// 	result_id := 0

// 	for rows.Next() {
// 		for i, _ := range columns {
// 			valuePtrs[i] = &values[i]
// 		}
// 		rows.Scan(valuePtrs...)
// 		tmpStruct := map[string]string{}
// 		for i, col := range columns {
// 			var v interface{}
// 			val := values[i]
// 			b, ok := val.([]byte)
// 			if ok {
// 				v = string(b)
// 			} else {
// 				v = val
// 			}
// 			tmpStruct[col] = fmt.Sprintf("%s", v)
// 		}

// 		finalResult[result_id] = tmpStruct
// 		result_id++

// 	}
// 	// fmt.Println(finalResult)
// }
