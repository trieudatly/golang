package main

import "fmt"

type Stack struct {
	data []interface{}
	top  int
}

func (s *Stack) Push(element interface{}) {
	s.top++
	s.data = append(s.data, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		s.top--
		last := s.data[s.top]
		s.data = s.data[:s.top]

		return last
	}

	return nil
}

func (s *Stack) Peek() interface{} {
	if len(s.data) > 0 {
		return s.data[s.top-1]
	}

	return nil
}

func (s *Stack) Clear() {
	s.top = 0
	s.data = []interface{}{}
}

func (s *Stack) Length() int {
	return s.top
}

var history Stack
var future Stack

func visit(url string) {
	history.Push(url)
	future.Clear()
}
func show(S Stack) {
	if S.Length() == 0 {
		fmt.Println("\nNo data")
	} else {
		for _, element := range S.data {
			fmt.Print(fmt.Sprint(element) + " | ")
		}
	}
}
func back() string {
	if history.Length() > 1 {
		future.Push(history.Pop())
		return fmt.Sprint(history.Peek())
	} else {
		return "No data"
	}
}
func forward() string {
	if future.Length() >= 1 {
		history.Push(future.Pop())
		return fmt.Sprint(history.Peek())
	} else {
		return "No data"
	}
}
func currentPage() {
	if history.Length() == 0 {
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
		fmt.Print("\nYour choice:")
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
