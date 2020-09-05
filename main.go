package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 {
		panic("Need file")
	}

	if len(os.Args) == 2 {
		panic("Need output file")
	}

	filename := os.Args[1]
	output := os.Args[2]

	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	r, err := regexp.Compile(`(\${.*})`)
	if err != nil {
		panic(err.Error())
	}

	ymlText := string(yml)

	varBytes := r.FindAll(yml, 10)

	for _, varByte := range varBytes {
		fmt.Println("Finded variable:", string(varByte))

		varPattern, err := regexp.Compile(`([A-Z_-az]+)`)
		if err != nil {
			panic(err.Error())
		}

		varName := string(varPattern.Find(varByte))

		ymlPattern := regexp.MustCompile(fmt.Sprintf(`\${%s}`, varName))
		ymlText = ymlPattern.ReplaceAllString(ymlText, os.Getenv(varName))
	}

	err = ioutil.WriteFile(output, []byte(ymlText), 0644)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("File has been successfully generated:", output)
}
