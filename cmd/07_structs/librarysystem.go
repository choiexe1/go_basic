package librarysystem

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
}

type Member struct {
	ID   string
	Name string
}

type Loan struct {
	Book
	Member
	DueDate string
}

func NewLoan(member Member, book Book, due string) Loan {
	return Loan{
		Book:    book,
		Member:  member,
		DueDate: due,
	}
}

// toString override와 동일..
func (l Loan) String() string {
	return fmt.Sprintf("%s by %s (%s) due %s", l.Title, l.Author, l.ISBN, l.DueDate)
}
