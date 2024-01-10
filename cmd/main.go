package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNum float32
	var floatNum2 float64
	floatNum = 3.1415926
	floatNum2 = 3.1415926
	fmt.Println(floatNum)
	fmt.Println(floatNum2)

	var result = float32(intNum) * floatNum
	fmt.Println(result)

	var byteNum byte
	byteNum = 255
	fmt.Println(byteNum)

	var boolNum bool
	boolNum = true
	fmt.Println(boolNum)

	var str string
	str = "hello world"
	fmt.Println(str)

	var myString string = "hello world" + "hello world"
	fmt.Println(myString)

	var myString2 = "hello world" + "hello world"
	fmt.Println(myString2)

	var myText = "hello"
	var myText2 = "world"
	fmt.Println(myText + myText2)

	myText3 := "hello"
	myText4 := "world"
	fmt.Println(myText3 + myText4)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	var1, var2 = var2, var1
	fmt.Println(var1 + var2)

	var1, var2 = var2, var1
	fmt.Println(var1 + var2)

	const var12 = 2
}
