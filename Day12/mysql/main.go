package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

func init() {
	//driverName:"mysql",  dataSoureName:用户名:密码@协议(地址:端口)/数据库名?参数=参数值
	db, _ = sql.Open("mysql", "root:mysqlroot@tcp(localhost:3306)/hjmstudy?charset=utf8")
}

func main() {
	//insert()

	//query()

	//delete()

	//update()

	transaction()
}


//事务
func transaction() {
	tx, txErr := db.Begin()
	if txErr!= nil {
		fmt.Println("开启事务失败,error:",txErr)
		return
	}
	//将xiaoming002 的年龄 + 10,同时将xiaoming003 的年龄 - 10

	prepareStmt1, err1 := tx.Prepare("update Student set age = 32 where name = 'xiaoming002'")
	prepareStmt2, err2 := tx.Prepare("update Student set age = 12 where name = 'xiaoming003'")
	if err1 != nil || err2 != nil {
		fmt.Println("执行失败.....回滚")
		tx.Rollback()
		return
	}
	result1, execErr1 := prepareStmt1.Exec()
	result2, execErr2 := prepareStmt2.Exec()
	if execErr1 != nil || execErr2 != nil {
		fmt.Println("执行失败.....回滚")
		tx.Rollback()
		return
	}
	result1RowAffected, resultErr1 := result1.RowsAffected()
	result2RowAffected, resultErr2 := result2.RowsAffected()

	if resultErr1 != nil || resultErr2 != nil {
		fmt.Println("执行失败.....回滚")
		tx.Rollback()
		return
	}
	if result1RowAffected == 1 && result2RowAffected == 1 {
		fmt.Println("执行成功")
		tx.Commit()
	} else{
		fmt.Println("执行失败.....回滚","result1RowAffected=",result1RowAffected,"result2RowAffected=",result2RowAffected)
		tx.Rollback()

	}



	fmt.Println("transaction")


}

//插入
func insert() {
	stmt, prepareErr := db.Prepare("insert into Student(name,age, sex,intrest) values (?,?,?,?)")
	if prepareErr != nil {
		fmt.Println("prepare fail err:", prepareErr)
		return
	}
	result, execErr := stmt.Exec("xiaoming003", 12, "男", "打球")
	if execErr != nil {
		fmt.Println("执行插入失败,err:", execErr)
		return
	}
	lastInsertId, e := result.LastInsertId()
	if e != nil {
		fmt.Println("获取插入记录失败")
		return
	}
	fmt.Println("lastInsertId", lastInsertId)
	rowsAffected, affectedErr := result.RowsAffected()
	if affectedErr != nil {
		fmt.Println("affectedErr:", affectedErr)
	}
	fmt.Println("rowsAffected", rowsAffected)

	defer func() {
		stmt.Close()
	}()
}

type Student struct {
	name    string
	age     int8
	sex     string
	intrest string
}

//查询
func query() {
	rows, e := db.Query("select name,age,sex,intrest from Student")
	if e != nil {
		fmt.Println("e,", e)
		return
	}

	stuSlice := make([]Student, 0)

	for rows.Next() {
		stu := Student{}
		err := rows.Scan(&(stu.name), &(stu.age), &(stu.sex), &(stu.intrest))
		if err != nil {
			fmt.Println("err=", err)
			continue
		}
		stuSlice = append(stuSlice, stu)
	}

	fmt.Println(stuSlice)
}

func delete() {
	stmt, prepareErr := db.Prepare("delete from Student where name = 'xiaoming001'")
	if prepareErr != nil {
		fmt.Println("prepareErr: ", prepareErr)
		return
	}
	result, execErr := stmt.Exec()
	if execErr != nil {
		fmt.Println("execErr: ", execErr)
		return
	}
	rowsAffected, e := result.RowsAffected()
	if e != nil {
		fmt.Println("e", e)
		return
	}
	fmt.Println("rowsAffected: ", rowsAffected)
}

func update() {
	stmt, e := db.Prepare("update Student set age = 22 where name='xiaoming002'")
	if e != nil {
		fmt.Println("e", e)
		return
	}
	result, i := stmt.Exec()
	if i != nil {
		fmt.Println("e", i)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("affected : ", affected)

}
