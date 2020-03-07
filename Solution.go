package main

import(
	"os"
	"bufio"

)

func Solution(ph []*Room) {
	fileOutput, _ :=os.Create("output.txt")
	writer := bufio.NewWriter(fileOutput)
	for _, p := range ph {
		writer.WriteString("L1-"+p.Name+"\n")
		writer.Flush() 
	}
}