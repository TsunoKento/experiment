package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	cmd := exec.Command("docker", "ps")

	cmd.Dir = filepath.Join("C:\\", "Users")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("終了するには何かキーを入力してください。")
	var input string
	fmt.Scanln(&input)

	os.Exit(0)
}
