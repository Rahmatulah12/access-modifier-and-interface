package main

/*
	Alias Ketika Import Package
	alias (s) => package yg diimport pribadi/access-nodifier/student
*/
import (
	"fmt"
	"math"
	"pribadi/access-modifier/library"
	s "pribadi/access-modifier/student"
	"strings"
)

func main() {
	library.SayHello("Rahmatulah Sidik")

	var student = s.Student{}
	student.Name = "Siska"
	fmt.Println("Nama :", student.Name)

	fmt.Printf("Name : %s\n", library.Student.Name)
	fmt.Printf("Gender : %s\n", library.Student.Gender)
	fmt.Printf("Age : %d\n", library.Student.Age)

	var bangunDatar hitung = persegi{10.0}
	fmt.Println("===== Persegi =====")
	fmt.Printf("Luas Persegi: %.2f\n", bangunDatar.luas())
	fmt.Printf("Keliling Persegi: %.2f\n", bangunDatar.keliling())
	fmt.Println()

	var lingkaran hitung = lingkaran{12.25}
	fmt.Println("===== Lingkaran =====")
	fmt.Printf("Luas Lingkaran: %.2f\n", lingkaran.luas())
	fmt.Printf("Keliling Lingkaran: %.2f\n", lingkaran.keliling())
	fmt.Println()

	var bangunRuang hitung2 = &kubus{20.5}
	fmt.Println("===== Kubus =====")
	fmt.Printf("Luas Kubus: %.2f\n", bangunRuang.luas())
	fmt.Printf("Keliling Kubus: %.2f\n", bangunRuang.keliling())
	fmt.Printf("Volume Kubus: %.2f\n", bangunRuang.volume())
	fmt.Println()

	/*
		interface kosong interface{}, adalah tipe data spesial.
		jika setelah interface ditambahkan kurung kurawal dibelakangnya,
		maka interface{} menjadi tipe data
	*/
	var secret interface{}
	secret = "Kalimat"
	fmt.Println("String :", secret)

	secret = 1550.21
	fmt.Println("Number :", secret)

	secret = 100
	fmt.Println("Nilai awal secret :", secret) // ditampilkan sebagai string
	// casting interface{}
	var number = secret.(int) * 2
	// tanpa casting
	// var number2 = secret + 3 // akan terjadi error jika dilakukan operasi matematika/perbandingan
	fmt.Printf("Number : %d\n", number)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println("Slice/Array :", secret)

	// akses slice/array secret
	// fmt.Println("index secret ke 0 :", secret[0]) // akan error, karena harus dicasting terlebih dahulu
	var secret0 = secret.([]string)
	fmt.Println("secret ke 0 :", secret0[0])

	// casting slice/array dari secret, menjadi string
	var hobbies = strings.Join(secret.([]string), ",")
	fmt.Println("hoby saya : ", hobbies)

	var data map[string]interface{} // tipe data collection

	data = map[string]interface{}{
		"name":    "Rahmatulah Sidik",
		"age":     27,
		"hobbies": []string{"Makan", "Pencak Silat", "Ngoding"},
	}

	fmt.Println()
	fmt.Println("===== Collection/map =====")
	fmt.Println("Name :", data["name"])
	fmt.Println("Age :", data["age"])
	fmt.Println("Hobby :", data["hobbies"])
	fmt.Println("All Data :", data)

	/*
		Casting variable interface kosong ke objek pointer
	*/
	var castingInterfaceKeObjekKosong interface{} = &person{name: "John", age: 50}
	var pointerName = castingInterfaceKeObjekKosong.(*person).name
	fmt.Println(pointerName)

}

type hitung interface {
	luas() float64
	keliling() float64
}
type hitung2 interface {
	hitung2d
	hitung3d
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type person struct {
	name string
	age  int
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2)
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}
