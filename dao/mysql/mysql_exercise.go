/**
@description  go-sql-driver/mysql练习
			  事物会影响增、删、改操作的效率
@author pujie
@data	2018-01-15
@参考
	[Golang操作数据库](https://www.cnblogs.com/mafeng/p/6207281.html)
	[golang学习之旅：使用go语言操作mysql数据库](https://www.cnblogs.com/tsiangleo/p/4483657.html)
**/

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(10.130.130.76:3306)/sigmob_web")
	if err != nil {
		log.Fatalf("open mysql address err, message=%v \n", err)
	}
	//insert(db,"14_1", "1,2,3")
	//query(db)
	queryRow(db, "1")
	//queryAll(db)
	defer db.Close()
}

// 数据插入，同时获取插入数据的系统id。ls
func insert(db *sql.DB, indexTerm, adLocalIdList string) error {
	stmt, err := db.Prepare("insert into sig_index_inverted_file(index_term, ad_local_id_list) values (?,?) ")
	if err != nil {
		log.Printf("db.Prepare err. message=%v \n", err)
		return err
	}
	rs, err := stmt.Exec(indexTerm, adLocalIdList)
	if err != nil {
		log.Printf("stmt.Exec() err. message=%v \n", err)
		return err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Printf("rs.LastInsertId() err. message=%v \n", err)
		return err
	}
	affect, err := rs.RowsAffected()
	if err != nil {
		log.Printf("rs.RowsAffected() err. message=%v \n", err)
		return err
	}
	log.Printf("info: index_term=%s, ad_local_id_list=%s, id=%d, rowsAffected=%d \n", indexTerm, adLocalIdList, id, affect)
	return nil
}

//查询语句
func query(db *sql.DB) error {
	var indexTerm, adLocalIdList string
	rows, err := db.Query("select index_term, ad_local_id_list from sig_index_inverted_file")
	if err != nil {
		log.Printf("db.Query() err, message=%v \n", err)
		return err
	}
	defer rows.Close()
	colums, _ := rows.Columns()
	for i, n := range colums { //遍历colums
		log.Printf("index=%d, column names=%s \n", i, n)
	}
	for rows.Next() {
		err := rows.Scan(&indexTerm, &adLocalIdList)
		if err != nil {
			log.Printf("scan err, message=%v \n", err)
		}
		log.Printf("info: index_term=%s, ad_local_id_list=%s \n\t", indexTerm, adLocalIdList)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("last err: %v \n", err)
		return err
	}
	log.Printf("last: index_term=%s, ad_local_id_list=%s \n\t", indexTerm, adLocalIdList)
	return nil
}

//查询单行
func queryRow(db *sql.DB, id string) error {
	var startDate string
	err := db.QueryRow("select startdate  from sig_dsp_adver_campaign where id=?", id).Scan(&startDate)
	if err != nil { //如果没有结果，则返回error
		log.Printf("err: id=%s, start_date=%s \n\t", id, startDate)
	} else {

		s, err := time.Parse("2006-01-02 15:04:05", "2018-01-02 16:01:20")
		//s, err := time.ParseInLocation("2006-01-02 15:04:05", "2018-01-02 16:01:20", time.Local)
		if err != nil {
			log.Printf("时间格式异常：%s, err :%v \n", startDate, err)
		} else {

			log.Printf("err: id=%s, start_date=%v, utc-time=%d, again=%s \n\t", id, startDate, s.Unix(), s.Format("2006-01-02 15:04:05"))
		}

	}
	return nil
}

//select *
func queryAll(db *sql.DB) error {
	rows, err := db.Query("select * from sig_entity_channel")
	if err != nil {
		log.Printf("db.Query() err, message=%v \n", err)
		return err
	}
	defer rows.Close()
	colums, _ := rows.Columns()
	for i, n := range colums { //遍历colums
		log.Printf("colum name: index=%d, column names=%s \n", i, n)
	}
	types, _ := rows.ColumnTypes()
	for i, t := range types {
		log.Printf("colum type: index=%d, name=%s, type=%s \n", i, t.Name(), t.ScanType())
	}
	err = rows.Err()
	if err != nil {
		log.Printf("last err: %v \n", err)
		return err
	}
	return nil
}
