package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func main()  {
	fmt.Println("Henlo uold!")

	test := []Person{
		{Name: "rafael", Age: 29},
		{Name: "rayssa", Age: 28},
		{Name: "dandan", Age: 1},
	}

	fmt.Println("with converter")
	for _, v := range test {
		// converting int to string with this strconv.FormatInt(int64(v.Age), 10)
		fmt.Println("my name is " + v.Name + " and my age is " + strconv.FormatInt(int64(v.Age), 10))
	}
	fmt.Println("with formatter")
	for _, v := range test {
		// using only Printf to print all formatted values
		fmt.Printf("my name is %s and my age is %d\n", v.Name ,v.Age)
	}
}

