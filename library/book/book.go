package book

import (
	"fmt"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) GetTitle() string {
	return b.Title
}
func (b *Book) SetTitle(title string) {
	b.Title = title
}
func CreateBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}
func parseUppercase(str string) string {
	return strings.ToUpper(str)
}

func (b *Book) PrintBook() {
	title := parseUppercase(b.Title)
	author := parseUppercase(b.Author)

	fmt.Printf("Title %s: \n Author: %s \n Pages:%d ", title, author, b.Pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}
func (b *TextBook) PrintTextBook() {
	title := parseUppercase(b.Title)
	author := parseUppercase(b.Author)

	fmt.Printf("Title %s: \n Author: %s \n Pages:%d \n Editorial : %s \n Level: %s ", title, author, b.Pages, b.editorial, b.level)
}
