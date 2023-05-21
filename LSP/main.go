package main

import "fmt"

func main() {
	fmt.Println("\nKelas turunan harus dapat digunakan sebagai pengganti kelas induknya")

	impl := NewSuperClass()
	sub := Owner{book: impl}

	CRUD(sub)
}
