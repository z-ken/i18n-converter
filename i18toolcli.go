package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)



func main1() {

AGAIN1:
	fmt.Println("Please input the i18n json file name: ")

	var filename string
	fmt.Scanln(&filename)

	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("An error occurred on opening the inputfile:", inputError)
		goto AGAIN1
	}
	defer inputFile.Close()

	finfo, _ := inputFile.Stat()
	if finfo.IsDir() {
		fmt.Println("Please input a file path , directory path is invalid.")
		goto AGAIN1
	}

	var s []string
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		s = append(s, inputString)
		if readerError == io.EOF {
			break
		}
	}

	jsonContent := strings.Join(s, "")

AGAIN2:
	fmt.Println("Please input the locale no. (1:en_US, 2:zh_CN) : ")
	var localeNo int
	fmt.Scanln(&localeNo)

	var locale string
	if localeNo == 1 {
		locale = "en_US"
	} else if localeNo == 2 {
		locale = "zh_CN"
	} else {
		fmt.Println("Invalid locale number.")
		goto AGAIN2
	}

	var maps map[string]string
	if err := json.Unmarshal([]byte(jsonContent), &maps); err != nil {
		fmt.Println("Invalid json format !")
		goto AGAIN1
	}

	var sql []string
	sql = append(sql, "INSERT INTO core_sys.`sys_i18n_appearance` \n")
	sql = append(sql, "(`key`, `content`, `locale`, `system_code`, `created_time`, `created_by`) \n")
	sql = append(sql, "VALUES \n")

	ml := len(maps)
	i := 0
	for k,v := range maps {
		sql = append(sql, "(\"", k, "\", \"", v, "\", \"", locale, "\", \"CSL\", NOW(), 0)")
		if i++; i < ml {
			sql = append(sql, ",\n")
		}
	}

	result := strings.Join(sql, "")

	fmt.Println("The SQL statement is below:")
	fmt.Println()
	fmt.Println(result)

	fmt.Println()
	goto AGAIN1

}
