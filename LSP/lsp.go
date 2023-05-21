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

// Repository Book as Super Class
type SuperClassRepositoryBook interface {
	RepoGetAll() (*Books, error)
	RepoGetOne(title string) (*Book, error)
	RepoAdd(book Book) error
	RepoUpdate(title string, book Book) error
	RepoDelete(title string) error
}

// the implement used in subclass
type ImplementSuperClass struct {
	books Books
}

func NewSuperClass() SuperClassRepositoryBook {
	return &ImplementSuperClass{}
}

func (repo *ImplementSuperClass) RepoGetAll() (*Books, error) {
	if len(repo.books) == 0 {
		return nil, errors.New("data not found")
	}

	return &repo.books, nil
}

func (repo *ImplementSuperClass) RepoGetOne(title string) (*Book, error) {
	for _, v := range repo.books {
		if v.Title == title {
			return &v, nil
		}
	}

	return nil, errors.New("data not found")
}

func (repo *ImplementSuperClass) RepoAdd(book Book) error {
	repo.books = append(repo.books, book)

	if len(repo.books) == 0 {
		return errors.New("failed add book")
	}

	return nil
}

func (repo *ImplementSuperClass) RepoUpdate(title string, book Book) error {
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

func (repo *ImplementSuperClass) RepoDelete(title string) error {
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

// Owner as Sub Class
type Owner struct {
	book SuperClassRepositoryBook
}

func (sub *Owner) OwnerGetAll() error {
	books, err := sub.book.RepoGetAll()

	if err != nil {
		return err
	}

	fmt.Println(books)
	return nil
}

func (sub *Owner) OwnerGetOne(title string) error {
	book, err := sub.book.RepoGetOne(title)

	if err != nil {
		return err
	}

	fmt.Println(book)
	return nil
}

func (sub *Owner) OwnerAdd(book Book) error {
	if err := sub.book.RepoAdd(book); err != nil {
		return err
	}

	fmt.Println("add book success")
	return nil
}

func (sub *Owner) OwnerUpdate(title string, book Book) error {
	if err := sub.book.RepoUpdate(title, book); err != nil {
		return err
	}

	fmt.Println("update book success")
	return nil
}

func (sub *Owner) OwnerDelete(title string) error {
	if err := sub.book.RepoDelete(title); err != nil {
		return err
	}

	fmt.Println("delete book success")
	return nil
}

// CRUD
func CRUD(sub Owner) {
	book1 := Book{
		Title:       "New Book",
		Description: "New Release",
	}

	book2 := Book{
		Title:       "Learn Golang",
		Description: "Limited Edition",
	}

	fmt.Println("\nAdd Book 1")
	if err := sub.OwnerAdd(book1); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nCheck Book 1")
	if err := sub.OwnerGetOne(book1.Title); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nUpdate Book 1")
	if err := sub.OwnerUpdate(book1.Title, book2); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nGet All Book")
	if err := sub.OwnerGetAll(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nDelete Book 2")
	if err := sub.OwnerDelete(book2.Title); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nGet All Book")
	if err := sub.OwnerGetAll(); err != nil {
		fmt.Println(err)
	}
}
