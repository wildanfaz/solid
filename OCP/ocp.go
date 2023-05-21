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

type RepositoryBook interface {
	RepoGetAll() (*Books, error)
	RepoGetOne(title string) (*Book, error)
	RepoAdd(book Book) error
	RepoUpdate(title string, book Book) error
	RepoDelete(title string) error
}

// Example Implement 1 Using MySQL Database
type MockMySQLRepository struct {
	// Example Connection
	books Books
	// db    sql.DB
}

func NewRepoBookMySQL() RepositoryBook {
	return &MockMySQLRepository{}
}

func (repo *MockMySQLRepository) RepoGetAll() (*Books, error) {
	if len(repo.books) == 0 {
		return nil, errors.New("data not found")
	}

	return &repo.books, nil
}

func (repo *MockMySQLRepository) RepoGetOne(title string) (*Book, error) {
	for _, v := range repo.books {
		if v.Title == title {
			return &v, nil
		}
	}

	return nil, errors.New("data not found")
}

func (repo *MockMySQLRepository) RepoAdd(book Book) error {
	repo.books = append(repo.books, book)

	if len(repo.books) == 0 {
		return errors.New("failed add book")
	}

	return nil
}

func (repo *MockMySQLRepository) RepoUpdate(title string, book Book) error {
	for i, v := range repo.books {
		if v.Title == title {
			repo.books[i] = book
			return nil
		}

		if repo.books[i].Title == title {
			return errors.New("failed update book")
		}
	}

	return errors.New("data not found")
}

func (repo *MockMySQLRepository) RepoDelete(title string) error {
	for i, v := range repo.books {
		if v.Title == title {
			repo.books = append(repo.books[:i], repo.books[i+1:]...)
			return nil
		}

		if repo.books[i].Title == title {
			return errors.New("failed delete book")
		}
	}

	return errors.New("data not found")
}

// Example Implement 2 Using PostgreSQL
type MockPostgreSQLRepository struct {
	// Example Connection
	books Books
	// db    sql.DB
}

func NewRepoBookPostgreSQL() RepositoryBook {
	return &MockPostgreSQLRepository{}
}

func (repo *MockPostgreSQLRepository) RepoGetAll() (*Books, error) {
	if len(repo.books) == 0 {
		return nil, errors.New("data not found")
	}

	return &repo.books, nil
}

func (repo *MockPostgreSQLRepository) RepoGetOne(title string) (*Book, error) {
	for _, v := range repo.books {
		if v.Title == title {
			return &v, nil
		}
	}

	return nil, errors.New("data not found")
}

func (repo *MockPostgreSQLRepository) RepoAdd(book Book) error {
	repo.books = append(repo.books, book)

	if len(repo.books) == 0 {
		return errors.New("failed add book")
	}

	return nil
}

func (repo *MockPostgreSQLRepository) RepoUpdate(title string, book Book) error {
	for i, v := range repo.books {
		if v.Title == title {
			repo.books[i] = book
			return nil
		}

		if repo.books[i].Title == title {
			return errors.New("failed update book")
		}
	}

	return errors.New("data not found")
}

func (repo *MockPostgreSQLRepository) RepoDelete(title string) error {
	for i, v := range repo.books {
		if v.Title == title {
			repo.books = append(repo.books[:i], repo.books[i+1:]...)
			return nil
		}

		if repo.books[i].Title == title {
			return errors.New("failed delete book")
		}
	}

	return errors.New("data not found")
}

type ServiceBook interface {
	SvcGetAll() error
	SvcGetOne(title string) error
	SvcAdd(book Book) error
	SvcUpdate(title string, book Book) error
	SvcDelete(title string) error
}

type SvcBook struct {
	repo RepositoryBook
}

func NewSvcBook(repo RepositoryBook) ServiceBook {
	return &SvcBook{repo}
}

func (svc *SvcBook) SvcGetAll() error {
	books, err := svc.repo.RepoGetAll()

	if err != nil {
		return err
	}

	fmt.Println(books)
	return nil
}

func (svc *SvcBook) SvcGetOne(title string) error {
	book, err := svc.repo.RepoGetOne(title)

	if err != nil {
		return err
	}

	fmt.Println(book)
	return nil
}

func (svc *SvcBook) SvcAdd(book Book) error {
	if err := svc.repo.RepoAdd(book); err != nil {
		return err
	}

	fmt.Println("add book success")
	return nil
}

func (svc *SvcBook) SvcUpdate(title string, book Book) error {
	if err := svc.repo.RepoUpdate(title, book); err != nil {
		return err
	}

	fmt.Println("update book success")
	return nil
}

func (svc *SvcBook) SvcDelete(title string) error {
	if err := svc.repo.RepoDelete(title); err != nil {
		return err
	}

	fmt.Println("delete book success")
	return nil
}

// CRUD
func CRUD(svc ServiceBook) {
	book1 := Book{
		Title:       "New Book",
		Description: "New Release",
	}

	book2 := Book{
		Title:       "Learn Golang",
		Description: "Limited Edition",
	}

	fmt.Println("\nAdd Book 1")
	if err := svc.SvcAdd(book1); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nCheck Book 1")
	if err := svc.SvcGetOne(book1.Title); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nUpdate Book 1")
	if err := svc.SvcUpdate(book1.Title, book2); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nGet All Book")
	if err := svc.SvcGetAll(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nDelete Book 2")
	if err := svc.SvcDelete(book2.Title); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nGet All Book")
	if err := svc.SvcGetAll(); err != nil {
		fmt.Println(err)
	}
}
