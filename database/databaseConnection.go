package database

import "database/sql"

func DBinstance(){
db,err := sql.Open("mysql","root:password1@tcp(127.0.0.1:3306)/test")
}