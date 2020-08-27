package db

import "database/sql"
import _ "mysql-master"


func Connect() (*sql.DB, error){
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/showroom")
  if err != nil {
    return nil, err
  }
  return db, nil
}
