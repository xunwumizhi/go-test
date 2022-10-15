package main

import "fmt"

func main() {
	pwd, err1 := GetUserCredential("gxyu")
	if err1 != nil {
		fmt.Printf("error of GetUserCredential: %v \n", err1)
	}
	fmt.Println("get gxyu pwd: " + pwd)

	err2 := AddUser("gxyu", "123")
	if err2 != nil {
		fmt.Printf("error of AddUser: %v \n", err2)
	}

}
