package main

import (
    "os"
    "os/exec"
)

func main() {
    c1 := exec.Command("ls", "-l")
    c2 := exec.Command("grep", "-i", "o")
    c3 := exec.Command("wc", "-l")
    c2.Stdin, _ = c1.StdoutPipe()
    c3.Stdin, _ = c2.StdoutPipe()
    c3.Stdout = os.Stdout
    _ = c3.Start()
    _ = c2.Start()
    _ = c1.Run()
    _ = c2.Wait()
    _ = c3.Wait()
}
