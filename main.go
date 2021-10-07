package main

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

	// fsL, err := os.ReadDir(".") ///home/mostain/go/src/practice
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for i, v := range fsL {
	// 	//v.Info()
	// 	fmt.Println(i, v.IsDir(), v.Name(), v.Type().Perm())
	// }

	//open the folder as as file, much faster/efficient solution
	// fd, err := os.Open("..")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fsList, err := fd.Readdir(1024)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for i, file := range fsList {
	// 	fmt.Println(i, file.Name(), file.IsDir(), file.Mode().Perm())
	// }

	//pta := &Anytype{name: "Mostain"}
	//pta := new(Anytype)
	// nat := NewAnytype()

	// fmt.Printf("%T\n", nat)
	// fmt.Println(nat.name, nat.Pp().name, nat.Pp())

	//3 fallthrough, break how it works
	// tier := 1
	// age := 18
	// switch tier { // switch statement
	// case 1: // case statement
	// 	fmt.Println("T-shirt")
	// 	if age < 18 {
	// 		break // exits the switch block
	// 	}
	// 	fallthrough // executes the next case
	// case 2:
	// 	fmt.Println("Mug")
	// 	fallthrough // executes the next case
	// case 3:
	// 	fmt.Println("Sticker pack")
	// default: // executed if no case is satisfied
	// 	fmt.Println("no reward")

	// }

	//example-4 select channel
	// var ch1 = make(chan int)
	// var ch2 = make(chan int)
	// var a, b int

	// go func() {
	// 	time.Sleep(1 * time.Microsecond)
	// 	ch1 <- 10
	// }()

	// go func() {
	// 	time.Sleep(2 * time.Microsecond)
	// 	b = <-ch2
	// }()

	// select { // the first operation that completes is selected

	// case a = <-ch1: //receiving from a channel
	// 	fmt.Println(a)

	// case ch2 <- 20: //sending to a channel
	// 	fmt.Println(b)
	// }

	// fmt.Println(a, b)

}
