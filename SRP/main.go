package main

import (
	"fmt"
)

// cd SRP
// go run .
// cd ..

func main() {
	fmt.Println("\nDalam konteks CRUD, implementasi yang tepat dari SRP dapat dilakukan dengan memisahkan tanggung jawab antara lapisan repository dan lapisan service")

	repo := NewRepoBook()
	svc := NewSvcBook(repo)

	CRUD(svc)
}
