package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"bufio"
	"os"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Ayush@1507"
	dbname   = "go"
)

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

// ------------------------------------------------------------------------------

func insertInput()(int64,string,int64,string){
	fmt.Println("# Inserting Row Values")

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Printf("Type student_id : ")
	scanner.Scan()
	id,_:=strconv.ParseInt(scanner.Text(),10,64)

	fmt.Printf("Type student_name : ")
	scanner.Scan()
	name:= scanner.Text()

	fmt.Printf("Type student_age : ")
	scanner.Scan()
	age,_:=strconv.ParseInt(scanner.Text(),10,64)

	fmt.Printf("Type student_address : ")
	scanner.Scan()
	address:= scanner.Text()

	return id,name,age,address

}

func insertTable(db *sql.DB,id int64,name string,age int64,address string){
	sql1:=`INSERT INTO student values($1,$2,$3,$4)`
	_,err:=db.Exec(sql1,id,name,age,address)
	checkErr(err)
}

// ------------------------------------------------------------------------------

func updateInput()(int64,string,string){
	fmt.Println("# Updating Row Values")

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Printf("Type student_id : ")
	scanner.Scan()
	id,_:=strconv.ParseInt(scanner.Text(),10,64)

	fmt.Printf("Type student_name : ")
	scanner.Scan()
	name:= scanner.Text()

	fmt.Printf("Type student_address : ")
	scanner.Scan()
	address:= scanner.Text()

	return id,name,address

}

func updateTable(db *sql.DB,id int64,name string,address string){
	fmt.Println("# Updating Row")
	sql2 := `UPDATE student SET name = $2, address = $3 WHERE id = $1;`
	_, err := db.Exec(sql2, id, name, address)
	checkErr(err)
}

// ------------------------------------------------------------------------------

func deleteRowFromTable(db *sql.DB){
	fmt.Println("# Deleting Row")

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Printf("Type student_id : ")
	scanner.Scan()
	id,_:=strconv.ParseInt(scanner.Text(),10,64)

	sql3 := `DELETE FROM student WHERE id = $1;`
	_, err := db.Exec(sql3, id)
	checkErr(err)

}

//  ------------------------------------------------------------------------------ 

func selectSingleRowFromTable(db *sql.DB){
	fmt.Println("# Query Single Row")

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Printf("Type student_id : ")
	scanner.Scan()
	idd,_:=strconv.ParseInt(scanner.Text(),10,64)

	sql4 := `SELECT id, address FROM student WHERE id=$1;`
	var address string
	var id int
	row := db.QueryRow(sql4, idd)
	switch err := row.Scan(&id, &address); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, address)
	default:
		checkErr(err)
	}
}

// ------------------------------------------------------------------------------ 

func selectMultiplerRowFromTable(db *sql.DB){
	fmt.Println("# Query Multiple Row")

	rows, err := db.Query("SELECT * FROM student")
	checkErr(err)
  	defer rows.Close()
  	for rows.Next(){
		var id int
		var name string
		var age int
		var address string
		err = rows.Scan(&id, &name ,&age ,&address)
		checkErr(err)
		fmt.Println(id, name,age,address)
	  }
  	err = rows.Err()
	checkErr(err)

	
}


func main(){
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)
	checkErr(err)
	defer db.Close()

	fmt.Println("Type 1 for Insert\nType 2 for Update\nType 3 for Delete\nType 4 for Select One Row\nType 5 for Select Multiple Row")

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter : ")
	scanner.Scan()
	flag,_:=strconv.ParseInt(scanner.Text(),10,64)

	if flag == 1{
		// Inserting into table
		id,name,age,address:=insertInput()
		insertTable(db,id,name,age,address)
	}else if flag == 2{
		// Updating table values
		id,name,address:=updateInput()
		updateTable(db,id,name,address)
	}else if flag == 3{
		// Deleting table row
		deleteRowFromTable(db)
	}else if flag == 4{
		// Selecting Single Row
		selectSingleRowFromTable(db)
	}else if flag == 5{
		// Select Multiple Row
		selectMultiplerRowFromTable(db)
	}else{
		fmt.Println("You entered wrong value")
	}

}

