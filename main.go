package main
import "fmt"

func main() {
	circlez()
	school()
}

func strings (){
	a := "I love to"
	b := " listen to music"
	c := ",what about you?"

	d := a + b + c
	fmt.Println(d)
}

func integers (){
	a := 39
	g := 13

	fmt.Println(a/g)

}

func boolean () {
	if 8/4 == 2 {
		fmt.Println("8 is divisible by 2")
	}
}

func floatingPoint(){
	var m1 float32
	m1 = 43.4343434343
	fmt.Println(m1)
}

func arr(){
	colors := []string{"blue", "red", "green", "white", "black","Grey","pink", "orange"}
	var numerals []int = []int{1, 2, 3, 4, 5}
	fmt.Println(numerals, colors)
}

func circlez() {	 
   var Experts = []string{"doctor", "lawyer", "policeman", "architect"} 
	for index, value := range Experts {
		fmt.Println(index, value)
	}
}

func school (){
	type student struct{
		name string
		class string
		Age  int
		session string
	}

	var Moses student

	Moses.name = "Moses"
	Moses.class = "jss2"
	Moses.Age = 67

	

	James := student{
		name: "Imaobong",
		class: "primary2",
		Age: 13,
		session: "second term",
	}

fmt.Println(Moses, James)
}
party := [Person{tobi, chuks}]
people := []Person {
	{
		Age: 1432,
		address: "djt"
	},
	{
		Age: 1432,
		address: "dfd"
	}
}

