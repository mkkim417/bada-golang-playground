package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	createFolder()
}

func loadFolderNames() []string {
	names, err := ioutil.ReadFile("./folderNames.txt")
	checkError(err)
	namesSlice := strings.Split(string(names), "\n")

	return namesSlice
}

func getFolderName() string {
	rand.Seed(time.Now().UnixNano())

	folderNames := loadFolderNames()
	choosedNum := rand.Intn(len(folderNames))
	folderName := folderNames[choosedNum]

	return folderName
}

func createFolder() {
	folderName := getFolderName()
	err := os.Mkdir(folderName, 0755)
	checkError(err)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
