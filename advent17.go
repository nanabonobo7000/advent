package main

import("fmt";"bufio";"os";"strings";"strconv")

type point struct{
    x, y int
}

func main(){
    fmt.Println("Input coordinates:")
    
    reader:= bufio.NewReader(os.Stdin)
    var m [][]int
    row:=0
    //valoss
    
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        m = append(m,make([]int,len(s)))
        for i,c := range s{
            m[row][i],_ = strconv.Atoi(string(c))
        } 
        row++

    }
    
    start, end := point{0, 0}, point{len(m[0]) - 1, len(m) - 1}
    heatLoss := countHeatLoss(m,start,end,1,3)
    fmt.Println("Min Heat loss with max distance 3 is:",heatLoss)
    
    heatLoss = countHeatLoss(m,start,end,4,10)
    fmt.Println("Min Heat loss with min distance 4 and max distance 10 is:",heatLoss)
    
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type HeatState struct {
	p, dir point
	streak     int
}

func countHeatLoss(m [][]int, start point, end point, minS int, maxS int) int{
    pointsToCheck := []HeatState{{p:start, dir:point{1, 0}, streak:0}, {p:start, dir:point{0, 1}, streak:0}}
    visited := map[HeatState]int{{p:start, dir:point{0, 0}, streak:0}: 0}
    minHeatLoss := 999999999
    
    for len(pointsToCheck) > 0 {
		current := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]
		
		if current.p == end && current.streak >= minS {
            minHeatLoss = min(minHeatLoss, visited[current])
        }
        
        for _,dir := range [3]point{current.dir, dirLeft(current.dir), dirRight(current.dir)}{
            nextP := point{current.p.x + dir.x, current.p.y + dir.y}
            if !(nextP.x>=0&&nextP.x<=end.x&&nextP.y>=0&&nextP.y<=end.y){
                continue
            }
            
            totalHeatLoss := visited[current] + getHeatLoss(m,nextP)
            nextStreak := 1
            if dir == current.dir{
                nextStreak = current.streak + 1
            }
            if(dir==current.dir && current.streak < maxS)||(dir!=current.dir && current.streak>= minS){
                nextState := HeatState{nextP,dir,nextStreak}
                if val, found := visited[nextState]; !found||val > totalHeatLoss{
                    visited[nextState] = totalHeatLoss
                    pointsToCheck = append(pointsToCheck, nextState)
                }
            }
            
        }
    }
    
    return minHeatLoss
}

func getHeatLoss(grid [][]int, p point) int {
	return int(grid[p.y][p.x])
}

func dirLeft(p point) point {
	return point{p.y, -p.x}
}

func dirRight(p point) point {
	return point{-p.y, p.x}
}




