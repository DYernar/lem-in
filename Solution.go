package main

import(
	"os"
	"bufio"
	"strconv"
)

func Solution(ph []*Room) {
	fileOutput, _ :=os.Create("output.txt")
	writer := bufio.NewWriter(fileOutput)
		for _, p := range ph {
			for i := 0; i < GlobalData.AntNum; i++ {

			writer.WriteString(" L"+strconv.Itoa(i)+"-"+p.Name)
		}
		writer.WriteString("\n")
		writer.Flush() 

	}

}