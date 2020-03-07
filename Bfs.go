package main

func BrFS(s *Room, e *Room) []*Room {
	prev := Solve(s)

	return ReconstructPath(s,e, prev)
}


func Solve(s *Room) []*Room {
	var q []*Room
	q = append(q, s)

	size := len(GlobalData.Rooms)

	visited := make([]bool, size)
	ind := GetIndex(s.Name, GlobalData)

	visited[ind] =true

	prev := make([]*Room, size)

	for len(q) != 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]

		for _, next := range node.Connections {
			index := GetIndex(next.Name, GlobalData)
			if !visited[index] && !next.Occupied {
				q = append(q, next)
				visited[index] = true
				prev[index] = node
			}
		}
	}
	return prev
}

func ReconstructPath(s *Room, e *Room, prev []*Room) []*Room {
	var path []*Room
	for at := e; at != nil; at = prev[GetIndex(at.Name, GlobalData)] {
		path =append(path, at)
	}

	return path
}