package main

import (
    "net/http"
    "fmt"
    "log"
    "flag"
    "os"
    "pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main(){
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
    
        if err != nil {
        log.Fatal(err)
    }

    pprof.StartCPUProfile(f)
    
    defer pprof.StopCPUProfile()
    
    http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request){
        fmt.Fprintf(w, "Hello World")
    })

    http.HandleFunc("/add", func(w http.ResponseWriter, _ *http.Request){
        fmt.Fprintf(w, "Writing from add")
    })

    err := http.ListenAndServe(":5000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
// package main

// import (
//     "flag"
//     "fmt"
//     "os"
// )

// func main() {

//     subOne := flag.NewFlagSet("one", flag.ExitOnError)
//     oneCream := subOne.Bool("cream", false, "Cream")
//     oneSugar := subOne.String("sugar", "", "Sugar")

//     subTwo := flag.NewFlagSet("two", flag.ExitOnError)
//     twoTea := subTwo.Int("tea", 0, "Tea")

//     if len(os.Args) < 2 {
//         fmt.Println("expected 'one' or 'two' subcommands")
//         os.Exit(1)
//     }

// 	fmt.Println(os.Args)

//     switch os.Args[1] {

//     case "one":
//         subOne.Parse(os.Args[2:])
//         fmt.Println("subcommand 'one'")
//         fmt.Println("  Cream:", *oneCream)
//         fmt.Println("  Sugar:", *oneSugar)
//         fmt.Println("  tail:", subOne.Args())
// 		i := subOne.Args()[0]
// 		fmt.Println(i)
//     case "two":
//         subTwo.Parse(os.Args[2:])
//         fmt.Println("subcommand 'two'")
//         fmt.Println("  tea:", *twoTea)
//         fmt.Println("  tail:", subTwo.Args())
//     default:
//         fmt.Println("expected 'one' or 'two' subcommands")
//         os.Exit(1)
//     }
// }
// package main

// import (
//     "flag"
//     "fmt"
// )

// func main() {

//     wordPtr := flag.String("flavor", "vanilla", "select shot flavor")
//     numbPtr := flag.Int("quantity", 2, "quantity of shots")
//     boolPtr := flag.Bool("cream", false, "decide if you want cream")

//     var order string
//     flag.StringVar(&order, "order", "complete", "status of order")

//     flag.Parse()

//     fmt.Println("flavor:", *wordPtr)
//     fmt.Println("quantity:", *numbPtr)
//     fmt.Println("cream:", *boolPtr)
//     fmt.Println("order:", order)
//     fmt.Println("tail:", flag.Args())
// 	// fmt.Println("D: ", flag.Args()[0])
// }
// package main

// import (
// // "strings"
// // "io"
// //"encoding/json"
// "fmt"
// "math/rand"
// // "time" b

// )



// const (
// 	Name = 1 << iota
// 	Age
// 	Live

// )

// func Writer(m *[]byte){
// 	// rand.Seed(time.Time)
// 	table := []byte("abcdefghijklmnopqrstuvwxyz1234567890@#$%{}ABCDEFGHIJKLMNOPQRSTUVWXYZ")
// 	fmt.Println(len(table))
// 	// fmt.Println(r, m)
// 	for i := 0; i < 300; i++ {
// 		r := rand.Intn(len(table))
// 		g := table[r]
// 		fmt.Println(g)
// 		a := append(*m, g)
// 		fmt.Println(a)
// 	}
// }

// func main() {
// 	fmt.Println(Name, Age, Live)
// 	fmt.Println(Name | Age | Live)
// 	fmt.Println(Name & Age)
	

// 	var mem []byte;

// 	Writer(&mem)


// 	// 	go func(){
// 	// 	fmt.Println("Println")
// 	// }()

	
// }

// // func main(){
// // 	type Person struct{
// // 		Name string
// // 		Age int
// // 		Fav string
// // 	}
	
// // 	var people []Person;
// // 	bytep, _ := os.ReadFile("a.json")
// // 	fmt.Println(string(bytep))
// // 	err := json.Unmarshal(bytep, &people)
	
// // 	if err == nil{
// // 		for i, e := range people {
// // 			fmt.Println(i, e.Name)
// // 		} 
// // 	}
// 	// per := Person{Name: "A", Age: 199, Fav: "food"}

// 	// en, _ := json.Marshal(per)
	
// 	// fmt.Println(string(en), per)
	
// 	// people := []Person{{"Ola", 21, "code"}, {"ade", 8, "food"}}
// 	// pe, _ := json.Marshal(people)
	
// 	// os.WriteFile("a.json", pe, 0666)
// 	// err := os.Remove("txt.log")
	
// 	// fmt.Println(err)
// 	// red, err := io.Writer([]byte{"Hello world"})
// 	// fmt.Println(string(read[:]), err)
// // }