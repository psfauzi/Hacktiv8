package main

// MINI QUIZ
// budi adalah user bertipe admin
// dia memiliki foto sebanyak 10 foto
// buatlah program yang akan menghasilkan output
// 1. jika user adalah admin dan foto > 10 => akan print "ok"
// 2. jika user adalah admin dan foto >= 10 => akan print "tidak ok"
// 3. selain diatas => akan print "oke gak ya?"
// kriteria:
// type user disimpan di package user
// batas foto disimpan di package photo
// program akan dijalankan di main.go
import (
	challange1 "Hactiv8/challange_1"
	"Hactiv8/paket"
	"Hactiv8/photo"
	"Hactiv8/user"
	"fmt"
)

func main() {
	paket.MyFunction()

	photo1 := 0
	if user.ADMIN&photo1 > photo.Photo {
		fmt.Println("ok")
	} else if user.ADMIN&photo1 >= photo.Photo {
		fmt.Println("Tidak Ok")
	} else {
		fmt.Println("OKe gak ya!")
	}

	challange1.Tantangan()
}
