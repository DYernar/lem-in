package main

import (
		"fmt"
		"os"
		"bufio"
		"strings"
		"strconv"
)
func GetIndex(name string, fullData AllData) int {
	for index, room := range fullData.Rooms {
		if room.Name == name {
			return index
		}
	}
	return -1
}


func Read(){
	if len(os.Args)== 2{

		file,err := os.Open(os.Args[1])
		if err!=nil{
			fmt.Print("Cant open file")
			return
		}
		s := bufio.NewScanner(file)
		num := false
		start := false
		end := false
		var Data AllData
		for s.Scan(){
			if !num{
				Data.AntNum, err = strconv.Atoi(s.Text())
				if err != nil{
					fmt.Print("Cant read number")
					return
				}
				num = true
				continue
			}

			if s.Text() == "##start" {
				start  = true
				continue
			}
			if s.Text() == "##end" {
				end  = true
				continue

			}
			if start {
				Data = NRoom(s.Text(), Data, start, end)
				start = false
				continue
			}
			if end {
				Data = NRoom(s.Text(), Data, start, end)
				end = false
				continue
			}
			rooms := strings.Split(s.Text(), " ")
			length := len(rooms)
			if length == 3{
				Data = NRoom(s.Text(), Data, start, end)
				continue
			}else{
				con := strings.Split(s.Text(), "-")
				if len(con) == 2{
					Data = Connection(con[0], con[1], Data)
				}
			}
		}
		GlobalData = Data
	}else{
		fmt.Println("Incorrect file name")
	}

}

func NRoom(str string, Data AllData, start bool, end bool) AllData{
	roomData := strings.Split(str, " ")
	if len(roomData) != 3 {
		fmt.Println("Invalid Room")
		return Data
	}
	var newRoom Room
	newRoom.Name = roomData[0]
	X, err := strconv.Atoi(roomData[1])
	if err != nil{
		fmt.Println("Invalid X coordinate")
	}
	Y, err := strconv.Atoi(roomData[2])
	if err != nil{
		fmt.Println("Invalid Y coordinate")
	}
	newRoom.X = X
	newRoom.Y = Y
	if start{
		newRoom.Start = true
	}else{
		newRoom.Start = false
	}
	if end {
		newRoom.End = true
	}else{
		newRoom.End = false
	}
	Data.Rooms = append(Data.Rooms, &newRoom)
	return Data
}

func Connection(name1 string, name2 string, Data AllData) AllData{
	ind1 := GetIndex(name1, Data)
	ind2 := GetIndex(name2, Data)
	contains := false
	for _, item := range Data.Rooms[ind1].Connections{
		if item.Name == name2{
			contains = true
			break
		}
	}
	if !contains{
		Data.Rooms[ind1].Connections = append(Data.Rooms[ind1].Connections, Data.Rooms[ind2])
		Data.Rooms[ind2].Connections = append(Data.Rooms[ind2].Connections, Data.Rooms[ind1])	
	}
	return Data
}