package main
import "fmt"

func main(){
	strings()
	integers()
	boolean ()
	floatingPoint()
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