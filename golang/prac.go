package main

import (
	"fmt"
	"os"
	"log"
	"bytes"
	"strings"
	"io"
)

type StorageType int

const (
    DiskStorage StorageType = 1 << iota
    TempStorage
    MemoryStorage
)

// func NewStore(t StorageType) Store {
//     switch t {
    // case MemoryStorage:
    //     return newMemoryStorage( /*...*/ )
    // case DiskStorage:
    //     return newDiskStorage( /*...*/ )
    // default:
    //     return newTempStorage( /*...*/ )
//     }
// }


func (s StorageType) storeType() string {
	switch s {
		case MemoryStorage:
			return "newMemoryStorage( /*...*/ )"
		case DiskStorage:
			return "newDiskStorage( /*...*/ )"
		default:
			return "newTempStorage( /*...*/ )"
		}
	}

func main(){
	
	var b bytes.Buffer // A Buffer needs no initialization.
	// b.Write([]byte("Hello "))
	// fmt.Fprintf(&b, "world!")
	// b.WriteTo(os.Stdout)
	b.Write([]byte("Hello  "))
	i, err := b.Read([]byte("llo0"))
	t, u := b.ReadBytes(72)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(string(t), u, err)
	// fmt.Println()
	// b.WriteTo(os.Stdout)
	s := strings.NewReader("Kungfu\n Jandol")
	fmt.Println(s, s.Len()) 
	// strings.WriteString(" Genji ")
	las := make([]byte, 5)
	m, e := s.Read(las)

	fmt.Println(m, e, s.Len(), "KUNGFUL KING", string(las))
	// s.WriteTo(os.Stdout)

	var g bytes.Buffer;
	g.Write([]byte("Porn"))

	const maxInt = int(^uint(0) >> 1)
	
	fmt.Println(maxInt)
	lass := make([]byte, 5)
	q := strings.NewReader("Azeez")

	// io.WriteString(q, ", I love you king!")
	io.Copy(&g, q)


	q.ReadAt(lass, 3)

	fmt.Println(string(lass), string(g.Bytes()), "Ire", string(las), "empty")
	// io.Copy(os.Stdout, os.Stdin)

	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	// g.Seek(2, io.SeekCurrent)
	ip := MemoryStorage

	fmt.Printf("%s", StorageType.storeType(DiskStorage), ip.storeType())
}


// car := NewVechicle()
// fmt.Println(car)
// car.VechicleVtype("car").VechicleAge(4).VechicleName("Lamborghini")

// fmt.Printf("%s is a %s with %d years on the streets", car.name, car.vtype, car.age)

// car.VechicleAge(40).VechicleVtype("plane").VechicleName("Virgin")

// fmt.Printf("%s is a %s with %d years on the streets", car.name, car.vtype, car.age)


type SpawnVehicle struct {
	name string
	age int
	vtype string
}

func NewVechicle () *SpawnVehicle {
	return new(SpawnVehicle)
}

func (s *SpawnVehicle) VechicleName (name string) *SpawnVehicle {
	(*s).name = name;
	return s
}

func (s *SpawnVehicle) VechicleAge (age int) *SpawnVehicle {
	s.age = age;
	return s
}


func (s *SpawnVehicle) VechicleVtype (vtype string) *SpawnVehicle {
	s.vtype = vtype;
	return s
}


// type Tree struct {
// 	Head *Node
// }

// func main(){
// 	tree := NewNode(20)
// 	tree.AddData(3)
// 	tree.AddData(3)
// 	tree.AddData(10)
// 	tree.AddData(100)
// 	tree.AddData(1000)
// 	tree.AddData(10000000)
// 	tree.AddData(24)
// 	tree.AddData(-1)
// 	tree.AddData(-10)
// 	tree.AddData(1023)
// 	tree.AddData(1200)
// 	tree.AddData(-10)
// 	tree.AddData(-400)
// 	tree.AddData(-400)
// 	tree.AddData(-4000)
// 	tree.AddData(-400)
// 	tree.AddData(-400)	
// 	tree.AddData(1)
	
// 	search, _ := Search(tree, 1)
// 	fmt.Println(search, "Found")
// 	// Print All data
// 	tree.PrintAll()
// 	fmt.Println("Tree Len:", tree.Len)
// 	tree1 := NewNode(202)
// 	tree1.AddData(1023)
// 	tree1.AddData(1200)
// 	tree1.AddData(-10)
// 	tree1.AddData(-400)
// 	tree1.PrintAll()
// 	fmt.Println("Tree Len:", tree1.Len)
// 	search, _ = Search(tree, 12)
// 	fmt.Println(search, "Found")
// 	search, trye := Search(tree, 1200)
// 	fmt.Println(search, " Found ", trye.index, trye.Number)
// 	search, _ = Search(tree, 120)
// 	fmt.Println(search, "Found")	
// }


// func InitDB(table Node) (*Node){
// 	return new(table)
// }
// []{name: name, password: password}
// name, password string
