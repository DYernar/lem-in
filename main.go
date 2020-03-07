package main

import "fmt"

var GlobalData AllData

func main(){
	Read()
	for _, room := range GlobalData.Rooms {
		fmt.Println(room.Name)
		for _, conns := range room.Connections {
			fmt.Println(conns.Name)
		}
		fmt.Println()
	}
}