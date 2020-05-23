package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var TEMP_DIR string = filepath.Join(os.TempDir(), "sshdiff")

func check_args() {
	if len(os.Args) != 3+1 {
		fmt.Println(`Usage:
sshdiff HOSTNAME_1 HOSTNAME_2 COMMAND
`)
		os.Exit(1)
	}
}

func dir_exist(dir string) bool {
	_, err := os.Stat(TEMP_DIR)
	if err != nil {
		return false
	} else {
		return true
	}
}

func mktempdir() {
	if !dir_exist(TEMP_DIR) {
		if err := os.Mkdir(TEMP_DIR, 0777); err != nil {
			panic(err)
		}
	}
}

func diff(a string, b string, hostname_1 string, hostname_2 string) string {
	mktempdir()
	t := time.Now()

	var (
		tempfile_1 string = filepath.Join(TEMP_DIR, hostname_1+t.Format("20060102150405"))
		tempfile_2 string = filepath.Join(TEMP_DIR, hostname_2+t.Format("20060102150405"))
	)

	err := ioutil.WriteFile(tempfile_1, []byte(a), 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(tempfile_2, []byte(b), 0644)
	if err != nil {
		panic(err)
	}

	stdout, _ := exec.Command("diff", "-u", tempfile_1, tempfile_2).Output()
	return string(stdout)
}

func run_ssh_command(host string, command string) string {
	fmt.Println("Running command in " + host + "..." + command)
	stdout, _ := exec.Command("ssh", host, command).Output()
	return string(stdout)
}

func main() {
	check_args()
	var (
		hostname_1 string = os.Args[1]
		hostname_2 string = os.Args[2]
		command    string = os.Args[3]
	)

	result1 := run_ssh_command(hostname_1, command)
	result2 := run_ssh_command(hostname_2, command)

	if result1 != result2 {
		fmt.Println(diff(result1, result2, hostname_1, hostname_2))
		os.Exit(-1)
	} else {
		fmt.Println("No difference. All results same")
		os.Exit(0)
	}
}
