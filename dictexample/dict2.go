package dictexample

import "fmt"

func Bookdic() Dictionary {
	book := Dictionary{}
	book.Add("Go Programming", "A book about Go programming language")
	book.Add("Python Programming", "A book about Python programming language")
	book.Add("Java Programming", "A book about Java programming language")
	return book
}

func Addkey() {
	m := make(map[string]int)
	m["a"] = 1
	m["a"]++
	fmt.Println("Map value for key 'a':", m["a"])
}
