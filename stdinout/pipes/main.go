package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {

	r1, w1 := io.Pipe()
	r2, w2 := io.Pipe()

	var cmds = []*exec.Cmd{
		exec.Command("cat", "students.txt"),
		exec.Command("grep", "Moaz"),
		exec.Command("wc", "-lmc"),
	}

	var bf bytes.Buffer

	cmds[1].Stdin, cmds[0].Stdout = r1, w1
	cmds[2].Stdin, cmds[1].Stdout = r2, w2
	//cmds[2].Stdout = os.Stdout //sending output
	cmds[2].Stdout = &bf //sending output to my variable

	for i := range cmds {
		if err := cmds[i].Start(); err != nil {
			log.Fatalln("Start", i, err)
		}

	}

	//fmt.Fprintln(os.Stdout,)
	//bs, err := cmds[2].Output()
	//fmt.Println(bs, err)

	for i, closer := range []io.Closer{w1, w2, nil} {

		if err := cmds[i].Wait(); err != nil {
			log.Fatalln("wait", i, err)
		}

		if closer == nil {
			continue
		}

		if err := closer.Close(); err != nil {
			log.Fatalln("Close", i, err)
		}

	}

	fmt.Println("Output:", bf.String())
}
