package main

import "fmt"

type User struct {
	name     string
	age      int8
	password string
}

func change(u User) User {
	u.age = 23
	return u

}

func main() {
	user := User{name: "Ivan", age: 14, password: "151242"}
	user = change(user)
	fmt.Println(user)
}
