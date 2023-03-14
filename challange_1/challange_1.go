package challange1

import (
	"fmt"
	"reflect"
)

/*
Buatlah sebuah program go dengan :menampilkan nilai i : 21 fmt.Printf("%T \n", i) //
contoh : fmt.Printf("%v \n", i)
menampilkan tipe data dari variabel i
menampilkan tanda %
menampilkan nilai boolean j : true
menampilkan nilai boolean j : true menampilkan unicode russia : Я (ya)
menampilkan nilai base 10 : 21
menampilkan nilai base 8 :25
menampilkan nilai base 16 : f
menampilkan nilai base 16 : F 13
menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
menampilkan float : 123.456000
menampilkan float scientific : 1.234560E+02 expected output
*/
func Tantangan() {
	i := 21
	j := true
	fmt.Println("Nilai dari variable i =", i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println("%")
	fmt.Println(j)
	fmt.Println(j)

	unicodeed := "\u042F"
	fmt.Println(unicodeed)

	base10 := 21
	fmt.Println(base10)

	base8 := 025
	fmt.Printf("%o\n", base8)

	base16 := 0xf
	fmt.Printf("%x\n", base16)

	base16_1 := 0xf13
	fmt.Printf("%x\n", base16_1)

	r := 'Я'
	fmt.Printf("karakter russian Я adalah %U\n", r)

	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%E", k)
}
