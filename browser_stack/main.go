package main

import "fmt"

//khai báo 1 struct gồm 1 mảng kiểu string và 1 biến count kiểu int
type Stack struct {
	data  []string //mảng chứa các phần tử của stack
	count int      //biến đếm số lượng phần tử trong stack
}

//tăng count 1 đơn vị và nối 1 phần tử vào đuôi của mảng
func (s *Stack) Push(element string) {
	s.count++
	s.data = append(s.data, element)
}

//
func (s *Stack) Pop() string {
	if len(s.data) > 0 {
		s.count--                 //count giảm 1 đơn vị
		last := s.data[s.count]   //tạo 1 biến last lấy giá trị tại index thứ count của mảng đưa vào biến last
		s.data = s.data[:s.count] //gán mảng thành 1 slice có giá trị từ đầu mảng đến index count
		return last
	}

	return "No data"
}

func (s *Stack) Peek() string {
	if len(s.data) > 0 {
		return s.data[s.count-1]
	}
	return "No data"
}

func (s *Stack) Clear() {
	s.count = 0
	s.data = nil
}

func (s *Stack) Length() int {
	return s.count
}

var history Stack
var future Stack

func Visit(url string) {
	history.Push(url)
	future.Clear()
}
func Show(s Stack) string {
	if s.Length() == 0 {
		return "No data"
	} else {
		show := ""
		for _, element := range s.data {
			show = show + "|" + element + "|"
		}
		return show
	}

}
func Back() string {
	if history.Length() > 1 {
		future.Push(history.Pop())
		return history.Peek()
	} else {
		return "No data"
	}
}
func Forward() string {
	if future.Length() >= 1 {
		history.Push(future.Pop())
		return history.Peek()
	} else {
		return "No data"
	}
}

func CurrentPage() string {
	if history.Length() == 0 {
		return "No data"
	} else {
		return history.Peek()
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
			fmt.Print("\nEnter address :")
			var name string
			fmt.Scanln(&name)
			Visit(name)
			fmt.Println(CurrentPage())
		case 2:
			fmt.Println("\n" + Back() + "\n")
		case 3:
			fmt.Println("\n" + Forward() + "\n")
		case 4:
			fmt.Println(Show(history))
		case 5:
			fmt.Println(Show(future))
		case 6:
			fmt.Println(CurrentPage())
		case 7:

		}
	}
}
