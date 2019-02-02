package models

import (
	"fmt"
	"crud/db"
)

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
}

type Users struct {
  Users    []User  `json:"user"`
}

func checkErr(err error){
  if err != nil {
    panic(err)
  }
}

func GetUser(id int) User{
  con     :=  db.Connect()
  query   :=  "select * from users where id=?"

  result  :=  User{}
  err     :=  con.QueryRow(query,id).Scan(&result.Id, &result.Name, &result.Email)

  checkErr(err)

  // fmt.Printf("id: %d\nname: %s\n", result.id, result.name)

  defer con.Close()

  return result
}

func GetUsers() Users{
  con     :=  db.Connect()
  query   :=  "select * from users"

  rows, err := con.Query(query)

  checkErr(err)

  defer rows.Close()

  result := Users{}
  user := User{}

  for rows.Next() {
    err := rows.Scan(&user.Id, &user.Name, &user.Email)
    checkErr(err)
    result.Users = append(result.Users, user)
    // fmt.Println(result.Name)
    // fmt.Println(result.Email)
  }
  defer con.Close()

  return result
}

func Insert(name, email string) {
  con     :=  db.Connect()
  _, err  :=  con.Exec("insert into users (name, email) values (?, ?)", name, email)
  checkErr(err)
  fmt.Println("Inserted!")
  defer con.Close()
}

func Update(id int, name string, email string){
	con     :=  db.Connect()
	_, err 	:=  con.Exec("update users set name = ?, email= ? where id = ?", name, email, id)
	checkErr(err)
  fmt.Println("update success!")
	defer con.Close()
}

func Delete(id int){
	con     :=  db.Connect()
	_, err 	:=  con.Exec("delete from users where id = ?", id)
	checkErr(err)
  fmt.Println("delete success!")
	defer con.Close()
}
