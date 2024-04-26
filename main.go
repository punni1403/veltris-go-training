package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/punni1403/traing-golang/hello"
)

func main() {

	hello.LargeCasefunc()

	// training_io()

	// fmt.Println("Starting main....")

	// str := "Veltris"

	// // str = "This is a \"quoted\" string."

	// l := len(str)       // find the length
	// substr1 := str[0:5] // substing

	// var substr string
	// substr = str[0:5]

	// unicodestr := "こんにちは、世界!"

	// le := len(unicodestr)

	// fmt.Printf(str, l, substr, substr1, unicodestr, le)

	// fmtstr := fmt.Sprintf("My compnay is %s", str)

	// fmt.Println(str, l, substr, substr1, unicodestr, le, fmtstr)

	// // intstr := "123a"
	// // inte, _:= strconv.Atoi(intstr)
	// // str_an := strconv.Itoa(inte)

	// // rune/runes - int32
	// unicodestr = "こんにちは、世界!"
	// for x := 0; x < len(unicodestr); x++ {
	// 	fmt.Println(string(unicodestr[x]))
	// }

	// for _, val := range unicodestr {
	// 	fmt.Print(string(val), " ")
	// }
}

func training_io() {

	// var name string

	// fmt.Println("Enter your name")
	// fmt.Scan(&name)

	// fmt.Printf("Entered name : %s", name)

	// data := []byte("This si traing on day 2 topi read wirte io")

	// ioutil.WriteFile("train.txt", data, 0644) // 777, 644

	// fmt.Println("file write successful")

	// fileData, _ := ioutil.ReadFile("train.txt")

	// fmt.Println("read data: ", string(fileData))

	file, err := os.Open("train.txt")
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
		fmt.Println("---")
	}

}
