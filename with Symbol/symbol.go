// package main

// import (
// 	"bufio"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"unicode/utf8"
// )

// func main() {

// 	c := make(map[string]string)
// 	d := make(map[string]string)
// 	j := make(map[string]string)
// 	sTable := make(map[string]string)
// 	a0 := "0"
// 	a1 := "1"

// 	sTable["R0"] = addZerosTo16Bit(strconv.FormatInt(int64(0), 2))
// 	sTable["R1"] = addZerosTo16Bit(strconv.FormatInt(int64(1), 2))
// 	sTable["R2"] = addZerosTo16Bit(strconv.FormatInt(int64(2), 2))
// 	sTable["R3"] = addZerosTo16Bit(strconv.FormatInt(int64(3), 2))
// 	sTable["R4"] = addZerosTo16Bit(strconv.FormatInt(int64(4), 2))
// 	sTable["R5"] = addZerosTo16Bit(strconv.FormatInt(int64(5), 2))
// 	sTable["R6"] = addZerosTo16Bit(strconv.FormatInt(int64(6), 2))
// 	sTable["R7"] = addZerosTo16Bit(strconv.FormatInt(int64(7), 2))
// 	sTable["R8"] = addZerosTo16Bit(strconv.FormatInt(int64(8), 2))
// 	sTable["R9"] = addZerosTo16Bit(strconv.FormatInt(int64(9), 2))
// 	sTable["R10"] = addZerosTo16Bit(strconv.FormatInt(int64(10), 2))
// 	sTable["R11"] = addZerosTo16Bit(strconv.FormatInt(int64(11), 2))
// 	sTable["R12"] = addZerosTo16Bit(strconv.FormatInt(int64(12), 2))
// 	sTable["R13"] = addZerosTo16Bit(strconv.FormatInt(int64(13), 2))
// 	sTable["R14"] = addZerosTo16Bit(strconv.FormatInt(int64(14), 2))
// 	sTable["R15"] = addZerosTo16Bit(strconv.FormatInt(int64(15), 2))
// 	sTable["SCREEN"] = addZerosTo16Bit(strconv.FormatInt(int64(16384), 2))
// 	sTable["KBD"] = addZerosTo16Bit(strconv.FormatInt(int64(24576), 2))
// 	sTable["SP"] = addZerosTo16Bit(strconv.FormatInt(int64(0), 2))
// 	sTable["LCL"] = addZerosTo16Bit(strconv.FormatInt(int64(1), 2))
// 	sTable["ARG"] = addZerosTo16Bit(strconv.FormatInt(int64(2), 2))
// 	sTable["THIS"] = addZerosTo16Bit(strconv.FormatInt(int64(3), 2))
// 	sTable["THAT"] = addZerosTo16Bit(strconv.FormatInt(int64(4), 2))

// 	c["M"] = "110000"
// 	c["!M"] = "110001"
// 	c["-M"] = "110011"
// 	c["M+1"] = "110111"
// 	c["M-1"] = "110010"
// 	c["D+M"] = "000010"
// 	c["D-M"] = "010011"
// 	c["M-D"] = "000111"
// 	c["D&M"] = "000000"
// 	c["D|M"] = "010101"

// 	c["0"] = "101010"
// 	c["1"] = "111111"
// 	c["-1"] = "111010"
// 	c["D"] = "001100"
// 	c["A"] = "110000"
// 	c["!D"] = "001101"
// 	c["!A"] = "110001"
// 	c["-D"] = "001111"
// 	c["-A"] = "110011"
// 	c["D+1"] = "011111"
// 	c["A+1"] = "110111"
// 	c["D-1"] = "001110"
// 	c["A-1"] = "110010"
// 	c["D+A"] = "000010"
// 	c["D-A"] = "010011"
// 	c["A-D"] = "000111"
// 	c["D&A"] = "000000"
// 	c["D|A"] = "010101"

// 	d["null"] = "000"
// 	d["M"] = "001"
// 	d["D"] = "010"
// 	d["MD"] = "011"
// 	d["A"] = "100"
// 	d["AM"] = "101"
// 	d["AD"] = "110"
// 	d["AMD"] = "111"

// 	j["null"] = "000"
// 	j["JGT"] = "001"
// 	j["JEQ"] = "010"
// 	j["JGE"] = "011"
// 	j["JLT"] = "100"
// 	j["JNE"] = "101"
// 	j["JLE"] = "110"
// 	j["JMP"] = "111"

// 	_, _ = os.Create("Add.hack")

// 	file, err := os.OpenFile("Add.hack", os.O_RDWR, 0666)
// 	fileAsm, err := os.OpenFile("Add.asm", os.O_RDWR, 0666)
// 	fileAsm2, err := os.OpenFile("Add.asm", os.O_RDWR, 0666)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()
// 	defer fileAsm.Close()

// 	bufferedWriter := bufio.NewWriter(file)
// 	firstPassReader := bufio.NewReader(fileAsm)
// 	bufferedReader := bufio.NewReader(fileAsm2)

// 	sTable = firstPass(sTable, firstPassReader)

// 	dataString := "data"
// 	symbolCounter := 16
// 	for {
// 		dataString, _ = bufferedReader.ReadString('\n')
// 		if len(dataString) == 0 { //Is it end of file?

// 			break
// 		}

// 		if strings.Contains(dataString, "//") {
// 			if isCommentAndCode(dataString) {
// 				dataString = removeCodeFromCommand(dataString)
// 			} else {
// 				continue
// 			}
// 		}

// 		if len(removeSpaces(dataString)) == 0 { //Is it empty line?
// 			continue
// 		}

// 		var compare []string

// 		if !strings.Contains(dataString, ";") {
// 			compare = strings.Split(removeSpaces(dataString), "=")
// 		}

// 		//if len(removeSpaces(dataString)) > 1 && !(strings.Contains(dataString, "")) {

// 		if strings.Contains(dataString, "@") {
// 			dataString = strings.ReplaceAll(dataString, "@", "")
// 			dataString = removeSpaces(dataString)

// 			if len(sTable[dataString]) == 0 {
// 				_, err := strconv.Atoi(dataString)
// 				if err != nil {
// 					sTable[dataString] = addZerosTo16Bit(strconv.FormatInt(int64(symbolCounter), 2))
// 					symbolCounter++
// 					_, _ = bufferedWriter.WriteString(
// 						sTable[dataString] + "\n",
// 					)

// 					bufferedWriter.Flush()
// 					continue
// 				}
// 				number, _ := strconv.ParseInt(dataString, 10, 64)
// 				binaryNumber := strconv.FormatInt(number, 2)

// 				_, _ = bufferedWriter.WriteString(
// 					addZerosTo16Bit(binaryNumber) + "\n",
// 				)

// 				bufferedWriter.Flush()
// 				continue
// 			} else {
// 				_, _ = bufferedWriter.WriteString(
// 					sTable[dataString] + "\n",
// 				)

// 				bufferedWriter.Flush()
// 				continue

// 			}
// 		} else {
// 			if strings.Contains(dataString, ")") && strings.Contains(dataString, "(") {
// 				continue
// 			} else if strings.Contains(dataString, ";") {
// 				splitted := strings.Split(removeSpaces(dataString), ";")

// 				_, err := bufferedWriter.WriteString(
// 					"111" + a0 + c[splitted[0]] + d["null"] + j[splitted[1]] + "\n",
// 				)

// 				bufferedWriter.Flush()

// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			} else if strings.Contains(compare[1], "A") && strings.Contains(dataString, "=") {
// 				splitted := strings.Split(removeSpaces(dataString), "=")

// 				_, err := bufferedWriter.WriteString(
// 					"111" + a0 + c[splitted[1]] + d[splitted[0]] + j["null"] + "\n",
// 				)

// 				bufferedWriter.Flush()

// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			} else if strings.Contains(compare[1], "M") && strings.Contains(dataString, "=") {
// 				splitted := strings.Split(removeSpaces(dataString), "=")

// 				_, err := bufferedWriter.WriteString(
// 					"111" + a1 + c[splitted[1]] + d[splitted[0]] + j["null"] + "\n",
// 				)

// 				bufferedWriter.Flush()

// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			}

// 		}
// 		//}
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// func trimFirstChar(s string) string {
// 	_, size := utf8.DecodeLastRuneInString(s)
// 	return removeSpaces(s[size:])
// }

// func convertToBinary(dataString string) string {
// 	trimmedData := trimFirstChar(dataString)
// 	number, _ := strconv.ParseInt(trimmedData, 10, 64)
// 	binaryNumber := strconv.FormatInt(number, 2)
// 	return binaryNumber
// }

// func addZerosTo16Bit(binaryNumber string) string {
// 	count := 16 - len(binaryNumber)
// 	for count != 0 {
// 		binaryNumber = "0" + binaryNumber
// 		count = count - 1
// 	}
// 	return binaryNumber
// }

// func removeSpaces(dataString string) string {
// 	return strings.Join(strings.Fields(dataString), "")
// }

// func isCommentAndCode(dataString string) bool {
// 	splittedText := strings.Split(dataString, "//")

// 	if len(removeSpaces(splittedText[0])) > 0 {
// 		return true
// 	} else {
// 		return false
// 	}

// }

// func removeCodeFromCommand(dataString string) string {
// 	splittedText := strings.Split(dataString, "//")
// 	return splittedText[0]
// }

// func firstPass(sTable map[string]string, bufferedReader *bufio.Reader) map[string]string {
// 	counter := 0

// 	for {
// 		dataString, _ := bufferedReader.ReadString('\n')
// 		if len(dataString) == 0 { //Is it end of file?
// 			break
// 		}

// 		if strings.Contains(dataString, "//") {
// 			if isCommentAndCode(dataString) {
// 				dataString = removeCodeFromCommand(dataString)
// 			} else {
// 				continue
// 			}
// 		}
// 		dataString = removeSpaces(dataString)

// 		if len(dataString) == 0 { //Is it empty line?
// 			continue
// 		}

// 		counter++
// 		if !strings.Contains(dataString, "(") {
// 			continue
// 		}
// 		counter--
// 		dataString = strings.ReplaceAll(dataString, "(", "")
// 		dataString = strings.ReplaceAll(dataString, ")", "")
// 		sTable[dataString] = addZerosTo16Bit(strconv.FormatInt(int64(counter), 2))

// 	}
// 	return sTable
// }
