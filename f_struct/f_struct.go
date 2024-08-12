package main
 
import "fmt"

type Student struct{
	rollno int
	name string
	marks int
}

func main(){
 var s1 = Student{1,"Ram",100}
 fmt.Println(s1)
 fmt.Println(s1.marks)

 var s2 = Student{rollno: 2,marks: 100}
 fmt.Println(s2)
 fmt.Println(s2.name)
}