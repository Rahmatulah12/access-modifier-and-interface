package library

import (
	"fmt"
)

var Student = struct {
	Name   string
	Age    int
	Gender string
}{}

/*
	init()
	Fungsi ini akan dipanggil pertama kali ketika package dimana fungsi berada di import.
	fungsi init() dijalankan sebelum fungsi main()
*/
func init() {
	Student.Name = "Anna"
	Student.Age = 22
	Student.Gender = "Perempuan"
	fmt.Println("--> access-modifier/library.go imported")
}

// ini adalah access modifier public, karena huruf awalnya kapital
func SayHello(name string) {
	fmt.Println("Hello Everyone")
	introduce(name)
}

// ini adalah private, karena diawali huruf kecil
func introduce(nama string) {
	fmt.Println("Hello nama saya :", nama)
}
