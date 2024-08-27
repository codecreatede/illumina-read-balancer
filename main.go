package fastq

/*

Author Gaurav Sablok
Universitat Potsdam
Date : 2024-8-27

a multiplex nextseq, novaseq read balancer for the Illumina reads using GO.
adding the support for the MongoDB file creation and a concurrent pattern.

*/

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	argsinput1 := os.Args[1:]
	argsinput2 := os.Args[2:]
	readinput1, err := os.OpenFile(argsinput1)
	if err != nil {
		panic(errors.New())
		log.Fatal(errors.New("input1 cant be empty"))
	}
	readinput2, err := os.OpenFile(argsinput2)
	if err != nil {
		panic(err)
		log.Fatal(Error.New("input2 cant be empty"))
	}
	open1 := bufio.NewScanner(readinput1)
	open2 := bufio.NewScanner(readinput2)

	openinputH1 := []string{}
	openinputSeq1 := []string{}
	openinputH2 := []string{}
	openinputSeq2 := []string{}

	for open1.Scan() {
		line := open1.Text()
		if strings.HasPrefix(line, "@") || strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH1, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			openinputSeq1 = append(openinputSeq1, line)
		}
	}

	for open2.Scan() {
		line := open2.Text()
		if strings.HasPrefix(line, "@") || strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH2, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			openinputSeq2 = append(openinputSeq2, line)
		}
	}

	balancedH1 := []string{}
	balancedH2 := []string{}
	balancedH1head := []string{}
	balancedH2head := []string{}

	for i := range openinputH1 {
		if string(openinputH1[i]) == string(openinputH2[i]) {
			balancedH1 = append(balancedH1, openinputH1[i])
			balancedH2 = append(balancedH2, openinputH2[i])
			balancedH1head = append(balancedH1head, openinputSeq1[i])
			balancedH2head = append(balancedH2head, openinputSeq2[i])
		}
	}

	checklen := len(balancedH1) == len(balancedH1)
	if checklen != true {
		panic(err.Error())
		log.Fatal(Error.new("The balancer is uneven and the reads cant be assembled"))

	} else {
		fmt.Println("The balancer is equal")
	}

}
