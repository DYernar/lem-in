package main

import "fmt"

var GlobalData AllData
var queue []*Room
var weights []int

func main(){
	Read()
	queue = append(queue, GlobalData.Rooms[0])


	for i := 0; i < len(GlobalData.Rooms); i++ {
		weights = append(weights, 0)
	}
	BFS(1, 0, len(GlobalData.Rooms))
	for _, i := range GlobalData.Rooms {
		fmt.Print("room " + i.Name+" ")
		fmt.Print(i.Distance)
		fmt.Println(" away " )
	}

}