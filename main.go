package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date : 2024-8-27

a multiplex nextseq, novaseq read balancer for the Illumina reads using GO.
adding the support for the MongoDB file creation and a concurrent pattern.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	argsinput1 := os.Args[1:]
	argsinput2 := os.Args[2:]
	readinput1, err := os.Open(argsinput1)
	if err != nil {
		log.Fatal(err)
	}
	readinput2, err := os.Open(argsinput2)
	if err != nil {
		log.Fatal(err)
	}
	open1 := bufio.NewScanner(readinput1)
	open2 := bufio.NewScanner(readinput2)

	openinputH1 := []string{}
	openinputSeq1 := []string{}
	openinputH2 := []string{}
	openinputSeq2 := []string{}

	for open1.Scan() {
		line := open1.Text()
		if strings.HasPrefix(string(line[0]), ">") && strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH1, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" ||
			string(line[0]) == "C" {
			openinputSeq1 = append(openinputSeq1, string(line))
		}
	}
	for open2.Scan() {
		line := open2.Text()
		if strings.HasPrefix(string(line[0]), ">") && strings.Contains(string(line), "length") {
			openinputH2 = append(openinputH2, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" ||
			string(line[0]) == "C" {
			openinputSeq2 = append(openinputSeq2, string(line))
		}
	}

	if len(openinputH1) == len(openinputSeq1) && len(openinputH2) == len(openinputSeq2) {
		fmt.Println("True")
	}

	balancedH1 := []string{}
	balancedH2 := []string{}
	balancedH1head := []string{}
	balancedH2head := []string{}

	for i := range openinputH1 {
		for j := range openinputH2 {
			if string(openinputH1[i]) == string(openinputH2[j]) {
				balancedH1 = append(balancedH1, openinputH1[i])
				balancedH2 = append(balancedH2, openinputH2[j])
				balancedH1head = append(balancedH1head, openinputSeq1[j])
				balancedH2head = append(balancedH2head, openinputSeq2[j])
			}
		}
	}
	if len(balancedH1) == len(balancedH2) {
		fmt.Println("Your reads are flushed and are present in the iters for further use")
	}
}
