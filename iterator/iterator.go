package iterator

import (
	"fmt"
)

type Aggregator[E comparable] interface {
	Iterator() Iterator[E]
}

type Iterator[E comparable] interface {
	HasNext() bool
	Next() (*E, error)
}

type Book struct {
	name string
}

func (b Book) Name() string {
	return b.name
}

type BookShelf struct {
	Books []Book
}
var _ Aggregator[Book] = &BookShelf{}

func (s BookShelf) BookAt(i int) *Book {
	return &s.Books[i]
}

func (s *BookShelf) Append(b Book) {
	s.Books = append(s.Books, b)
}

func (s BookShelf) Length() int {
	return len(s.Books)
}

func (s BookShelf) Iterator() Iterator[Book] {
	return &BookShelfIterator{BookShelf: s, Index: 0}
}

type BookShelfIterator struct {
	BookShelf BookShelf
	Index     int
}
var _ Iterator[Book] = &BookShelfIterator{}

func (i BookShelfIterator) HasNext() bool {
	return i.Index < i.BookShelf.Length()
}

func (i *BookShelfIterator) Next() (*Book, error) {
	if !i.HasNext() {
		return nil, fmt.Errorf("no next element")
	}

	b := i.BookShelf.BookAt(i.Index)
	i.Index += 1
	return b, nil
}
