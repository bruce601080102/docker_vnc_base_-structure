package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Conf struct {
	originGoFileName string
	outputFileName   string
	SSH_PASSWORD     string
	VNC_PASSWORD     string
}

func ReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func WriteFile(filename, txt string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println("刪除成功")

	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file is fail", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(txt)
	write.Flush()
	fmt.Println("Write success", filename)
}

func Generate(conf Conf, wait *sync.WaitGroup) {
	main_dockerfile := ReadFile("./os_version/" + conf.originGoFileName)

	install_package1 := ReadFile("./unit/install_package1")
	install_vnc2 := ReadFile("./unit/install_vnc2")
	install_vnc2 = strings.Replace(install_vnc2, "16313302", conf.VNC_PASSWORD, -1)
	install_ssh3 := ReadFile("./unit/install_ssh3")
	install_ssh3 = strings.Replace(install_ssh3, "16313302", conf.SSH_PASSWORD, -1)
	copy_srcipt4 := ReadFile("./unit/copy_srcipt4")

	all := install_package1 + install_vnc2 + install_ssh3 + copy_srcipt4

	main_dockerfile = strings.Replace(main_dockerfile, "# This is a positioning point", all, -1)
	// fmt.Println(diia_main_js)

	fileName := conf.outputFileName
	WriteFile(fileName, main_dockerfile)
	wait.Done()
}

func main() {
	path, _ := os.Getwd()
	err := godotenv.Load(path + "/env/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	TARGET_FileName := os.Getenv("TARGET_FileName")
	OUTPUT_FileName := os.Getenv("OUTPUT_FileName")
	SSH_PASSWORD := os.Getenv("SSH_PASSWORD")
	VNC_PASSWORD := os.Getenv("VNC_PASSWORD")

	wait := sync.WaitGroup{}
	wait.Add(1)
	start := time.Now()
	conf := Conf{
		outputFileName:   OUTPUT_FileName,
		originGoFileName: TARGET_FileName,
		SSH_PASSWORD:     SSH_PASSWORD,
		VNC_PASSWORD:     VNC_PASSWORD,
	}

	go Generate(conf, &wait)

	wait.Wait()
	elapsed := time.Now().Sub(start)
	fmt.Println("生成時間:", elapsed)
}
