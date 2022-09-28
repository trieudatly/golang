package main

import (
	"fmt"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

var history lls.Stack = *lls.New()
var future lls.Stack = *lls.New()

func visit(url string) {
	history.Push(url)
	future.Clear()
}
func show(Stack lls.Stack) {
	if Stack.Empty() {
		fmt.Println("\nNo data")
	} else {
		fmt.Println("\n" + Stack.String())
	}
}
func back() string {
	if history.Size() > 1 {
		hispop, _ := history.Pop()
		future.Push(hispop)
		hispeek, _ := history.Peek()
		str := fmt.Sprintf("%v", hispeek)
		return str
	} else {
		return "No data"
	}
}

func forward() string {
	if future.Size() >= 1 {
		fupop, _ := future.Pop()
		history.Push(fupop)
		hispeek, _ := history.Peek()
		str := fmt.Sprintf("%v", hispeek)
		return str
	} else {
		return "No data"
	}
}

func currentPage() {
	if history.Empty() {
		fmt.Println("\nNo data")
	} else {
		fmt.Println("\n")
		fmt.Println(history.Peek())
	}
}
func main() {
	var choice int

	for ok := true; ok; ok = choice != 7 {
		fmt.Print("\n1.Visit new page\n2.Go back\n3.Go forward\n4.Show history\n5.Show future\n6.Current page\n7.Exit")
		fmt.Print("\nYour choice :")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Print("\nEnter address:")
			var name string
			fmt.Scanln(&name)
			visit(name)
			currentPage()
		case 2:
			fmt.Println("\n" + back() + "\n")
		case 3:
			fmt.Println("\n" + forward() + "\n")
		case 4:
			show(history)
		case 5:
			show(future)
		case 6:
			currentPage()
		case 7:

		}
	}
}
