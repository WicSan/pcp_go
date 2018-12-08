package main

import (
    "fmt"
)

type User struct {
    name string
}

func (user User) String() string {
    return fmt.Sprintf("User: name = %s", user.name)
}

func main() {
    user := User{name: "foo"}
    fmt.Println(user)         // User: name = foo
}