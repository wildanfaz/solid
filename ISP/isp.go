package main

import (
	"errors"
	"fmt"
)

type Book struct {
	Title       string
	Description string
}

type Books []Book

// Separate Interface Into More Specific
type RepositoryGetBook interface {
	RepoGetAll() (*Books, error)
	RepoGetOne(title string) (*Book, error)
}

type ImplementRepositoryGetBook struct {
	books *Books
}

func (repo *ImplementRepositoryGetBook) RepoGetAll() (*Books, error) {
	if len(*repo.books) == 0 {
		return nil, errors.New("data not found")
	}

	return repo.books, nil
}

func (repo *ImplementRepositoryGetBook) RepoGetOne(title string) (*Book, error) {
	for _, v := range *repo.books {
		if v.Title == title {
			return &v, nil
		}
	}

	return nil, errors.New("data not found")
}

type RepositoryAddBook interface {
	RepoAdd(book Book) error
}

type ImplementRepositoryAddBook struct {
	books *Books
}

func (repo *ImplementRepositoryAddBook) RepoAdd(book Book) error {
	*repo.books = append(*repo.books, book)

	if len(*repo.books) == 0 {
		return errors.New("failed add book")
	}

	return nil
}

type RepositoryUpdateBook interface {
	RepoUpdate(title string, book Book) error
}

type ImplementRepositoryUpdateBook struct {
	books *Books
}

func (repo *ImplementRepositoryUpdateBook) RepoUpdate(title string, book Book) error {
	for i, v := range *repo.books {
		if v.Title == title {
			(*repo.books)[i] = book
			return nil
		}

		if (*repo.books)[i].Title == title {
			return errors.New("failed update book")
		}
	}

	return errors.New("data not found")
}

type RepositoryDeleteBook interface {
	RepoDelete(title string) error
}

type ImplementRepositoryDeleteBook struct {
	books *Books
}

func (repo *ImplementRepositoryDeleteBook) RepoDelete(title string) error {
	for i, v := range *repo.books {
		if v.Title == title {
			*repo.books = append((*repo.books)[:i], (*repo.books)[i+1:]...)
			return nil
		}

		if (*repo.books)[i].Title == title {
			return errors.New("failed delete book")
		}
	}

	return errors.New("data not found")
}

type RepositoryBook struct {
	getBook    RepositoryGetBook
	addBook    RepositoryAddBook
	updateBook RepositoryUpdateBook
	deleteBook RepositoryDeleteBook
}

func (repo *RepositoryBook) GetAll() error {
	books, err := repo.getBook.RepoGetAll()

	if err != nil {
		return err
	}

	fmt.Println(books)
	return nil
}

func (repo *RepositoryBook) GetOne(title string) error {
	book, err := repo.getBook.RepoGetOne(title)

	if err != nil {
		return err
	}

	fmt.Println(book)
	return nil
}

func (repo *RepositoryBook) Add(book Book) error {
	if err := repo.addBook.RepoAdd(book); err != nil {
		return err
	}

	fmt.Println("add book success")
	return nil
}

func (repo *RepositoryBook) Update(title string, book Book) error {
	if err := repo.updateBook.RepoUpdate(title, book); err != nil {
		return err
	}

	fmt.Println("update book success")
	return nil
}

func (repo *RepositoryBook) Delete(title string) error {
	if err := repo.deleteBook.RepoDelete(title); err != nil {
		return err
	}

	fmt.Println("delete book success")
	return nil
}

// CRUD
func CRUD(repo RepositoryBook) {
	book1 := Book{
		Title:       "New Book",
		Description: "New Release",
	}

	book2 := Book{
		Title:       "Learn Golang",
		Description: "Limited Edition",
	}

	fmt.Println("\nAdd Book 1")
	if err := repo.Add(book1); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nCheck Book 1")
	if err := repo.GetOne(book1.Title); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nUpdate Book 1")
	if err := repo.Update(book1.Title, book2); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nGet All Book")
	if err := repo.GetAll(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nDelete Book 2")
	if err := repo.Delete(book2.Title); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nGet All Book")
	if err := repo.GetAll(); err != nil {
		fmt.Println(err)
	}
}
