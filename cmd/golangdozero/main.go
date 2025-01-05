package main

import "github.com/fabianoflorentino/golangdozero/pkg/sendmail"

func main() {
	println("Aprenda Golang do Zero!")
	println(sendmail.CheckEmail("fabianoflorentino@outlook.com"))
}
