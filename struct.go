package main

type Room struct {
	Name string 
	X int
	Y int
	Connections []*Room
	Start bool
	End bool
}

type AllData struct {
	AntNum int
	Rooms []*Room
}