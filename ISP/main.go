package main

import "fmt"

func main() {
	fmt.Println("\nTidak boleh memaksa untuk mengimplementasikan metode yang tidak digunakan dan interface yang spesifik harus dipilih daripada satu interface umum")

	books := Books{}
	getBook := ImplementRepositoryGetBook{&books}
	addBook := ImplementRepositoryAddBook{&books}
	updateBook := ImplementRepositoryUpdateBook{&books}
	deleteBook := ImplementRepositoryDeleteBook{&books}

	repo := RepositoryBook{
		getBook:    &getBook,
		addBook:    &addBook,
		updateBook: &updateBook,
		deleteBook: &deleteBook,
	}

	CRUD(repo)
}
