package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"path"
)

func cypher(buff []byte) {
	for i :=0; i < len(buff); i++ {
		buff[i] = buff[i] ^ 255
	}
}

func processDir(directory, depth string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("%sDirectory: %s => does not exist or cannot be read\n", depth, directory)
		return
	}

	fmt.Printf("%sProcessing Directory: %s...\n", depth, directory)

	for _, fileInfo := range files {
		filePath := path.Join(directory, fileInfo.Name())
		fmt.Printf("%s  Opening: %s", depth, filePath)
		if fileInfo.IsDir() {
			fmt.Printf("... is a directory... processing recursively\n")
			processDir(filePath, depth + "  ")
			continue
		}

		fd, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("... error [%v] opening file, ignoring the file\n", err)
			continue
		}

		fmt.Printf("\n%s   ... Size: %d... Reading in a single chunk", depth, fileInfo.Size())
		buff := make([]byte, fileInfo.Size(), fileInfo.Size())
		n, err := fd.Read(buff)
		if err != nil { // Getting old fast
			fmt.Printf("... error reading file: %v ignoring\n", err)
			continue
		}
		fmt.Printf("\n%s   ... Read %d bytes... Cyphering", depth, n)
		cypher(buff)
		fd.Seek(0, 0)
		fmt.Printf("\n%s   ... Writing", depth)
		fd.Write(buff)
		fmt.Printf("... Closing")
		fd.Close()
		fmt.Printf("... Done\n")
	}
}

func main() {
	var directory string
	flag.StringVar(&directory, "dir", ".", "Specifies the directory to... er cipher")
	flag.Parse()

	fmt.Println("This tool will... cipher for the lack of a better word all the files on a dirtectory and its")
	fmt.Println("subdirectories. The ciphering is as simple as inverting all bytes so it should be easy to revert")
	fmt.Println("(hint: just run the tool again). But I cannot promise something won't broke and so I won't.")
	fmt.Println("And of course if it brokes, it's all on you")
	fmt.Println("Write 'yes' to continue, or anything else to quit")
	var maybe string
	fmt.Scanln(&maybe)
	if strings.ToLower(maybe) != "yes" {
		fmt.Println("Good Call!")
		os.Exit(0)
	}
	processDir(directory, "")
}
