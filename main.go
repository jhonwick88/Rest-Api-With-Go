package main

import (
	"fmt"
)

type Student struct {
	nama  string
	kelas int
}
type person struct {
	name string
	age  int
}

func main() {
	var siswa Student
	siswa.nama = "Imam Thoib"
	siswa.kelas = 1
	fmt.Println("Nama Siswa", siswa.nama)
	fmt.Println("kelas siswa", siswa.kelas)
	var siswa2 = Student{kelas: 10}
	fmt.Println("Nama Siswa 2", siswa2.nama)
	fmt.Println("kelas2 ", siswa2.kelas)
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
	// _ = "paijo gone"
	// firstname, lastname := "paidi", "panduwinata"
	// var message = `Nama saya "John Wick". Salam kenal. Mari belajar "Golang".`

	// fmt.Println("Hello world " + "my name " + firstname + " and the last name is " + lastname)
	// fmt.Println(message)

	// const lastName = "wick"
	// fmt.Print("nice to meet you ", lastName, "!\n")
	// fmt.Println("john wick time", time.Now())
	// fmt.Println("john", "wick")

	// fmt.Print("john wick")
	// fmt.Print("john ", "wick\n")
	// fmt.Print("john", " ", "wick\n")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka dd", i)
	// 	i++
	// }
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)
	// var minMax = func(min, max int) (int, int) {
	// 	return min, max
	// }

	// var min, max = minMax(1, 12000)
	// fmt.Println("min is", min)
	// fmt.Println("max is", max)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Angka is", i)
	// }
}
