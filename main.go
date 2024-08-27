package fastq

/*

Author Gaurav Sablok
Universitat Potsdam 
Date : 2024-8-27

a multiplex nextseq, novaseq read balancer for the Illumina reads using GO.

*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	argsinput1 := os.Args[1:]
	argsinput2 := os.Args[2:]
	argsoutput1 := os.Args[3:]
	argsoutput2 := os.Args[4:]
	if argsinput1 || argsinput2 || argsinput3 || argsinput4 == nil {
		panic(err)
		log.Fatal(err.Error.new("The arguments cant be empty"))
	}
    
    readinput1, err := os.OpenFile(argsinput1)
    if err != nil {
    		panic(err.Error())
    		log.Fatal(err)
    	}
    readinput2 := os.OpenFile(argsinput2)
    if err != nil {
    		panic(err.Error())
    		log.Fatal(err)
    	}
    open1 := bufio.NewScanner(readinput1)
    open2 := bufio.NewScanner(readinput2)

    var openinputH1 []string
    var openinputSeq1 []string
    var openinputH2 []string
    var openinputSeq2 []string

    for open1.Scan() {
		line := open1.Text()
		if string(line).HasPrefix("@") || strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH1, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			openinputSeq1 = append(openinputSeq1, line)
		}
	}

	 for open2.Scan() {
		line := open2.Text()
		if string(line).HasPrefix("@") || strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH1, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			openinputSeq1 = append(openinputSeq1, line)
		}
	}
    
    balancedH1 = []string
    balancedH2 = []string
    balancedH1head = []string
    balancedH2head = []string

    for i := range openinputH1 {
        if string(openinputH1[i]) == string(openinput[H2]) {
        	balancedH1 = append(balancedH1, openinputH1[i])
        	balancedH2 = append(balancedH2, openinputH2[i])
        	balancedH1head = append(balancedH1head = openinputSeq1[i])
        	balancedH2head = append(balancedH2head = openinputSeq[i])
        }
    }

    len, err := len(balancedH1) == len(balancedH1) 
    if err != nil {
    	panic(err.Error())
    	log.Fatal(Error.new("The balancer is uneven and the reads cant be assembled"))

    } else {
    	fmt.println("The balancer is equal")
    }

}
