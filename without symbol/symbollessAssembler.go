package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	c := make(map[string]string)
	d := make(map[string]string)
	j := make(map[string]string)

	a0 := "0"
	a1 := "1"

	c["M"] = "110000"
	c["!M"] = "110001"
	c["-M"] = "110011"
	c["M+1"] = "110111"
	c["M-1"] = "110010"
	c["D+M"] = "000010"
	c["D-M"] = "010011"
	c["M-D"] = "000111"
	c["D&M"] = "000000"
	c["D|M"] = "010101"

	c["0"] = "101010"
	c["1"] = "111111"
	c["-1"] = "111010"
	c["D"] = "001100"
	c["A"] = "110000"
	c["!D"] = "001101"
	c["!A"] = "110001"
	c["-D"] = "001111"
	c["-A"] = "110011"
	c["D+1"] = "011111"
	c["A+1"] = "110111"
	c["D-1"] = "001110"
	c["A-1"] = "110010"
	c["D+A"] = "000010"
	c["D-A"] = "010011"
	c["A-D"] = "000111"
	c["D&A"] = "000000"
	c["D|A"] = "010101"

	d["null"] = "000"
	d["M"] = "001"
	d["D"] = "010"
	d["MD"] = "011"
	d["A"] = "100"
	d["AM"] = "101"
	d["AD"] = "110"
	d["AMD"] = "111"

	j["null"] = "000"
	j["JGT"] = "001"
	j["JEQ"] = "010"
	j["JGE"] = "011"
	j["JLT"] = "100"
	j["JNE"] = "101"
	j["JLE"] = "110"
	j["JMP"] = "111"

	//_, _ = os.Create("Add.asm")
	_, _ = os.Create("PongL.hack")

	file, err := os.OpenFile("PongL.hack", os.O_RDWR, 0666)
	fileAsm, err := os.OpenFile("PongL.asm", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer fileAsm.Close()

	bufferedWriter := bufio.NewWriter(file)

	if err != nil {
		log.Fatal(err)
	}

	bufferedReader := bufio.NewReader(fileAsm)

	dataString := "data"
	for len(dataString) > 0 {

		dataString, err = bufferedReader.ReadString('\n')
		var compare []string
		if !strings.Contains(dataString, ";") {
			compare = strings.Split(removeSpaces(dataString), "=")
		}

		if len(removeSpaces(dataString)) > 1 && !(strings.Contains(dataString, "//")) {

			binaryNumber := convertToBinary(dataString)

			if strings.Contains(dataString, "@") {

				_, err := bufferedWriter.WriteString(
					addZerosTo16Bit(binaryNumber) + "\n",
				)

				bufferedWriter.Flush()

				if err != nil {
					log.Fatal(err)
				}
			} else {
				if strings.Contains(dataString, ";") {
					splitted := strings.Split(removeSpaces(dataString), ";")

					_, err := bufferedWriter.WriteString(
						"111" + a0 + c[splitted[0]] + d["null"] + j[splitted[1]] + "\n",
					)

					bufferedWriter.Flush()

					if err != nil {
						log.Fatal(err)
					}
				} else if (strings.Contains(compare[1], "A") && strings.Contains(dataString, "=")) || !(strings.Contains(compare[1], "M")) {
					splitted := strings.Split(removeSpaces(dataString), "=")

					_, err := bufferedWriter.WriteString(
						"111" + a0 + c[splitted[1]] + d[splitted[0]] + j["null"] + "\n",
					)

					bufferedWriter.Flush()

					if err != nil {
						log.Fatal(err)
					}
				} else if strings.Contains(dataString, "M") && strings.Contains(dataString, "=") {
					splitted := strings.Split(removeSpaces(dataString), "=")

					_, err := bufferedWriter.WriteString(
						"111" + a1 + c[splitted[1]] + d[splitted[0]] + j["null"] + "\n",
					)

					bufferedWriter.Flush()

					if err != nil {
						log.Fatal(err)
					}
				}

			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}

}

func trimFirstChar(s string) string {
	_, size := utf8.DecodeLastRuneInString(s)
	return removeSpaces(s[size:])
}

func convertToBinary(dataString string) string {
	trimmedData := trimFirstChar(dataString)
	number, _ := strconv.ParseInt(trimmedData, 10, 64)
	binaryNumber := strconv.FormatInt(number, 2)
	return binaryNumber
}

func addZerosTo16Bit(binaryNumber string) string {
	count := 16 - len(binaryNumber)
	for count != 0 {
		binaryNumber = "0" + binaryNumber
		count = count - 1
	}
	return binaryNumber
}

func removeSpaces(dataString string) string {
	return strings.Join(strings.Fields(dataString), "")
}
