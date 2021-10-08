package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

	//example -5 mini cmd app which count all the folders and files in a given path
	// if len(os.Args) != 2 { // ensure path is specified
	// 	fmt.Println("Please specify a path.")
	// 	return
	// }
	// root, err := filepath.Abs(os.Args[1]) // get absolute path
	// if err != nil {
	// 	fmt.Println("Cannot get absolute path:", err)
	// 	return
	// }
	// fmt.Println("Listing files in", root)
	// var c struct {
	// 	files int
	// 	dirs  int
	// }
	// filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	// walk the tree to count files and folders
	// 	if info.IsDir() {
	// 		c.dirs++
	// 	} else {
	// 		c.files++
	// 	}
	// 	fmt.Println("-", path)
	// 	return nil
	// })
	// fmt.Printf("Total: %d files in %d directories", c.files, c.dirs)

	//example-6
	// args := "ls -la"
	// argSlc := strings.Split(args, " ")
	// appName := argSlc[0]
	// cmd := exec.Command(appName, argSlc[1:]...)

	// stderr, _ := cmd.StderrPipe()
	// cmd.Start()

	// scanner := bufio.NewScanner(stderr)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	m := scanner.Text()
	// 	fmt.Println(m)
	// }
	// cmd.Wait()

	//example-7

	// cmdName := "passwd mostain" //apt search php7.2
	// cmdArgs := strings.Fields(cmdName)

	// cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	// stdout, _ := cmd.StdoutPipe()
	// cmd.Start()

	// go func() {
	// 	oneByte := make([]byte, 1024)
	// 	for {
	// 		_, err := stdout.Read(oneByte)
	// 		if err != nil {
	// 			break
	// 		}
	// 		fmt.Printf("%v \n", string(oneByte))
	// 	}
	// }()

	// time.Sleep(time.Second * 2)

	//example-8
	// cmdName := "passwd mostain" //ping 127.0.0.1
	// cmdArgs := strings.Fields(cmdName)

	// cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	// stdout, _ := cmd.StdoutPipe()
	// cmd.Start()
	// oneByte := make([]byte, 100)
	// num := 1
	// for {
	// 	_, err := stdout.Read(oneByte)
	// 	if err != nil {
	// 		fmt.Printf(err.Error())
	// 		break
	// 	}
	// 	r := bufio.NewReader(stdout)
	// 	line, _, _ := r.ReadLine()
	// 	fmt.Println(string(line))
	// 	num = num + 1
	// 	if num > 3 {
	// 		os.Exit(0)
	// 	}
	// }

	// cmd.Wait()

	//example -9
	//findCmd := cmd.NewCmd("find", "/", "-name", "v2ray")
	//findCmd := cmd.NewCmd("passwd", "mostain")
	//statusChan := <-findCmd.Start() // non-blocking
	//fmt.Println(statusChan, findCmd.Stderr)

	// Print each line of STDOUT from Cmd
	// for _, line := range statusChan.Stdout {
	// 	fmt.Println(line)
	// }

	// ticker := time.NewTicker(2 * time.Second)

	// // Print last line of stdout every 2s
	// go func() {
	// 	for range ticker.C {
	// 		status := findCmd.Status()
	// 		n := len(status.Stdout)
	// 		fmt.Println(status.Stdout[n-1])
	// 	}
	// }()

	// // Stop command after 1 hour
	// go func() {
	// 	<-time.After(1 * time.Hour)
	// 	findCmd.Stop()
	// }()

	// // Check if command is done
	// select {
	// case finalStatus := <-statusChan:
	// 	// done
	// 	fmt.Println("done:", finalStatus)
	// default:
	// 	// no, still running
	// 	fmt.Println("running...")
	// }

	// // Block waiting for command to exit, be stopped, or be killed
	// finalStatus := <-statusChan
	// fmt.Println(finalStatus)

	//example -10
	// Disable output buffering, enable streaming
	// cmdOptions := cmd.Options{
	// 	Buffered:  true,
	// 	Streaming: true,
	// }

	// // Create Cmd with options
	// envCmd := cmd.NewCmdOptions(cmdOptions, "passwd", "mostain") //./print-some-lines

	// // Print STDOUT and STDERR lines streaming from Cmd
	// doneChan := make(chan struct{})
	// go func() {
	// 	defer close(doneChan)
	// 	// Done when both channels have been closed
	// 	// https://dave.cheney.net/2013/04/30/curious-channels
	// 	for envCmd.Stdout != nil || envCmd.Stderr != nil {
	// 		select {
	// 		case line, open := <-envCmd.Stdout:
	// 			if !open {
	// 				envCmd.Stdout = nil
	// 				continue
	// 			}
	// 			fmt.Println(line)

	// 		case line, open := <-envCmd.Stderr:
	// 			if !open {
	// 				envCmd.Stderr = nil
	// 				continue
	// 			}
	// 			fmt.Fprintln(os.Stderr, line)
	// 		}
	// 	}
	// }()

	// // Run and wait for Cmd to return, discard Status
	// <-envCmd.Start()
	// //fmt.Println(co)

	// // Wait for goroutine to print everything
	// <-doneChan
	//fmt.Println(dc)

	//time.Sleep(time.Second * 3)

	//example-11
	//
	//cmd := exec.Command("passwd", "mostain")
	// stdoutPipe, _ := cmd.StdoutPipe()
	// stdoutReader := bufio.NewReader(stdoutPipe)
	// cmd.Start()

	// Datas := make(chan Data, 100)
	// go dataProcessor(Datas)

	// bufioReader := bufio.NewReader(stdoutReader)
	// for {

	// 	//r := bufio.NewReader(stdoutPipe)
	// 	output, _, err := bufioReader.ReadLine()
	// 	if err != nil {
	// 		fmt.Println("err:", err)
	// 	}
	// 	fmt.Println(string(output))
	// 	var tempData Data
	// 	tempData.Out = output
	// 	if err != nil || err == io.EOF {
	// 		break
	// 	}
	// 	Datas <- tempData
	// }

	//example-12
	if len(os.Args) < 3 {
		fmt.Println("Please specify a path and a search string.")
		return
	}
	root, err := filepath.Abs(os.Args[1]) // get absolute path
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}
	q := []byte(strings.Join(os.Args[2:], " "))
	fmt.Printf("Searching for %q in %s...\n", q, root)
	err = filepath.Walk(root, func(path string, info os.FileInfo,
		err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = ioutil.ReadAll(io.TeeReader(f, queryWriter{q, os.Stdout}))
		return err
	})
	if err != nil {
		fmt.Println(err)
	}

}

type Data struct {
	Out []byte
}

func dataProcessor(Datas <-chan Data) {
	for {
		select {
		case data := <-Datas:
			fmt.Println(">>", string(data.Out))
		default:
			continue
		}
	}
}

type queryWriter struct {
	Query []byte
	io.Writer
}

func (q queryWriter) Write(b []byte) (n int, err error) {
	lines := bytes.Split(b, []byte{'\n'})
	l := len(q.Query)
	for _, b := range lines {
		i := bytes.Index(b, q.Query)
		if i == -1 {
			continue
		}
		for _, s := range [][]byte{
			b[:i],              // what's before the match
			[]byte("\x1b[31m"), //star red color
			b[i : i+l],         // match
			[]byte("\x1b[39m"), // default color
			b[i+l:],            // whatever is left
		} {
			v, err := q.Writer.Write(s)
			n += v
			if err != nil {
				return 0, err
			}
		}
		fmt.Fprintln(q.Writer)
	}
	return len(b), nil
}
