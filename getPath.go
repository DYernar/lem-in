package main

import(
	"fmt"
)

var visited []string
var r []Room
func BFS(distance int, num int, length int) {
	zero := 0
	for _, i := range weights {
		if i == 0 {
			zero ++
		}
	}
	if zero == 1 {
		return
	}

	if num >= len(queue) {
		fmt.Println("incorrect connections")
		return
	}
	queue[num].Visited = true

	for _, q := range queue {
		for _ , room := range q.Connections {
		if !room.Visited {
			room.Visited = true
			alreadyContains := false
			for _, contains := range queue {
				if room == contains {
					alreadyContains = true
				}
			}
			if !alreadyContains {
				if (distance == 0) {
					continue
				} else {
					indx := GetIndex(room.Name, GlobalData)
					if weights[indx] != 0{
						if weights[indx] > distance {
							room.Distance = distance
							weights[indx] = distance
						}
					} else {
						room.Distance = distance
						weights[indx] = distance
					}
				}
				queue = append(queue, room)
			}
		}

	}
	}

	fmt.Println()

	distance += 1
	num++
	
	BFS(distance, num, length)
}

