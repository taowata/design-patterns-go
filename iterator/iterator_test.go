package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := &BookShelf{}

	titles := []string{
		"Around the World in 80 Days",
		"Bible",
		"Cinderella",
		"Daddy-Long-Legs",
	}

	for _, v := range titles {
		bookShelf.Append(Book{name: v})
	}
	i := 0
	var it Iterator[Book] = bookShelf.Iterator()
	for it.HasNext() {
		want := titles[i]
		book, err := it.Next()
		if err != nil {
			t.Errorf(err.Error())
		}
		if book.Name() != want {
			t.Errorf("expected %v, got %v", want, book.Name())
		}
		i += 1
	}
}
