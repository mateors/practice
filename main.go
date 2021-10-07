package main

import (
	"fmt"
	"os"
)

type Anytype struct {
	name string
}

// V1PipeReader reads from the V1Pipe
type AnytypeReader Anytype

//func (r *V1PipeReader) pp() *V1Pipe                { return (*V1Pipe)(r) }
func (r *AnytypeReader) Pp() *Anytype {
	return (*Anytype)(r)
}

func NewAnytype() *AnytypeReader {

	// return &AnytypeReader{
	// 	name: "Billah",
	// }
	at := &Anytype{name: "Mostain"}
	return (*AnytypeReader)(at)
}

func main() {

	fsL, err := os.ReadDir(".") ///home/mostain/go/src/practice
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range fsL {

		//v.Info()
		fmt.Println(i, v.IsDir(), v.Name(), v.Type().Perm())

	}

	//open the folder as as file, much faster/efficient solution
	fd, err := os.Open("..")
	if err != nil {
		fmt.Println(err)
	}

	fsList, err := fd.Readdir(1024)
	if err != nil {
		fmt.Println(err)
	}
	for i, file := range fsList {
		fmt.Println(i, file.Name(), file.IsDir(), file.Mode().Perm())
	}

	//pta := &Anytype{name: "Mostain"}
	//pta := new(Anytype)
	nat := NewAnytype()

	fmt.Printf("%T\n", nat)
	fmt.Println(nat.name, nat.Pp().name, nat.Pp())
}
