package main

import "fmt"

// cd OCP
// go run .
// cd ..

func main() {
	fmt.Println("\nMenambahkan implementasi baru tanpa mengubah code yang sudah ada")

	fmt.Println("\n---Implementasi CRUD 1 Mock MySQL---")
	ExampleMockMySQL()

	fmt.Println("\n---Implementasi CRUD 2 Mock PostgreSQL---")
	ExampleMockPostgreSQL()
}

// Implement 1
func ExampleMockMySQL() {
	repo := NewRepoBookMySQL()
	svc := NewSvcBook(repo)
	CRUD(svc)
}

// Implement 2
func ExampleMockPostgreSQL() {
	repo := NewRepoBookPostgreSQL()
	svc := NewSvcBook(repo)
	CRUD(svc)
}
