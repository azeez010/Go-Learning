package main

import (
	"fmt"
	"os/exec"
	"bytes"
	// "os"
	// "io"
	"strings"
	"log"
	
)
// make ai
func main() {
	cmd := exec.Command("py", "print.py")
	// stdin, _ := cmd.StdinPipe()

	// io.WriteString(stdin, "Azeez")
	// var r strings.Reader

	// r.Read([]byte("Hello"))
	
	// for i := 0; i < 2; i++ {
	// 	cmd.Stdin = strings.NewReader("Azeez")
	// }
	// cmd.Stdin = os.Stdin//[]*io.Reader{strings.NewReader("Azeez"), strings.NewReader("Azeez")}
	
	cmd.Stdin = strings.NewReader("some input1 ")
	// go func(){
	// 	stdin.Write([]byte("Az"))
	// }()

	path, err := exec.LookPath("print.py")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	
	fmt.Println("Complete path", path)
	
	var out bytes.Buffer
	cmd.Stdout = &out

	// cmd.Stdin = strings.NewReader("Azeez")
	// cmd.Stdout = &out
	
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %q\n", out.String())
}

// // func Calc(num ...int) int {

// // }

// func main(){
// 	cmd := exec.Command("py", "print.py")
// 	stdin, err := cmd.StdinPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// 	go func() {
// 		defer stdin.Close()
// 		io.WriteString(stdin, "Azeez")
// 		io.WriteString(stdin, "Olabode")
// 		fmt.Println(stdin)
// 	}()
	
	
// 	out, err2 := cmd.CombinedOutput()
// 	if err2 != nil {
// 		log.Fatal(err2, "Stderr")
// 	}
	
// 	fmt.Printf("%s\n", out)
// }
