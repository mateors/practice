package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	words   = []string{"Moaz", "Rezaul", "Imran", "Anonnya"}
	cmds    = make([][2]*exec.Cmd, len(words))
	writers = make([]io.Writer, len(words))
	buffers = make([]bytes.Buffer, len(words))
	err     error
)

func main() {

	//Method-1
	//app-1
	// cmd := exec.Command("ls", "-l")
	// iwc, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer iwc.Close()
	// cmd.Start()

	// //app-2
	// cmd2 := exec.Command("grep", "students.txt")
	// cmd2.Stdin = iwc

	// cout2, err := cmd2.Output()
	// if err != nil {
	// 	log.Fatal("ERROR:", err)
	// }
	// fmt.Println(string(cout2))
	// cmd.Wait()

	//Method-2
	cmd1 := exec.Command("ls", "-l")
	cmd2 := exec.Command("grep", "students.txt")
	cmd3 := exec.Command("wc", "-lm")

	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd3.Stdin, _ = cmd2.StdoutPipe()
	//cmd3.Stdout = os.Stdout
	var bs bytes.Buffer
	cmd3.Stdout = &bs

	//cmd3.Start()
	//cmd2.Start()
	//cmd1.Run()
	cmd1.Start()
	cmd2.Start()
	cmd3.Start()

	cmd1.Wait()
	cmd2.Wait()
	cmd3.Wait()
	fmt.Println(bs.String())

	os.Exit(0)
	//var buf bytes.Buffer
	//iwc.Write(&buf)

	// /cmd.Start()
	// cmd := exec.Command("cat")
	// stdin, err := cmd.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// go func() {
	// 	defer stdin.Close()
	// 	io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	// }()

	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s\n", out)
	// os.Exit(0)

	for i := range words {

		cmds[i][0] = exec.Command("grep", words[i])
		if writers[i], err = cmds[i][0].StdinPipe(); err != nil {
			log.Fatal("in pipe", i, err)
		}

		//cmds[i][0].Start()

		cmds[i][1] = exec.Command("wc", "-l")

		if cmds[i][1].Stdin, err = cmds[i][0].StdoutPipe(); err != nil {
			log.Fatal("in pipe", i, err)
		}
		cmds[i][1].Stdout = &buffers[i]
	}

	cat := exec.Command("cat", "students.txt")
	cat.Stdout = io.MultiWriter(writers...)

	//
	for i := range cmds {
		if err := writers[i].(io.Closer).Close(); err != nil {
			log.Fatalln("close 0", i, err)
		}
	}
	for i := range cmds {
		if err := cmds[i][0].Wait(); err != nil {
			log.Fatalln("grep wait", i, err)
		}
	}

	for i := range cmds {
		if err := cmds[i][1].Wait(); err != nil {
			log.Fatalln("wc wait", i, err)
		}
		count := bytes.TrimSpace(buffers[i].Bytes())
		log.Printf("%10q %s entries", cmds[i][0].Args[1], count)
	}
}
