package main

import("fmt";"bufio";"os";"strings";"strconv")

type point struct{
    x int
    y int
}

func main(){
    fmt.Println("Input directions:")
    
    var ss []string
    
    trench := map[point]string{}
    p := point{0,0}
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        
        ss = append(ss,s)
        var dir string
        var mag int
        var color string
        
        ss := strings.Split(s," ")
        dir = ss[0]
        mag,_ = strconv.Atoi(ss[1])
        color = ss[2]
        
        for i:=0;i<mag;i++{
            switch dir{
                case "U":
                    p.y--
                case "D":
                    p.y++
                case "L":
                    p.x--
                case "R":
                    p.x++
            }
            trench[p] = color
        }
        
    }
    
    floodFill(trench, point{1, 1})
    //fmt.Println(trench)

	var result = len(trench)
	fmt.Println("Part 1 directions Result: ", result)
    
   
    
    nodes := make([]point,0,len(ss))
    nodes = append(nodes,point{0,0})
    p = point{0,0}
    for _,l := range ss{
        sss := strings.Split(l," ")
        color := sss[2][2:8]
        //fmt.Println(color)
        dist,_ := strconv.ParseInt(color[:5],16,0)
        dir,_ := strconv.Atoi(string(color[5]))
        switch dir{
            case 0:
                fmt.Printf("#%s = R %d\n",color,dist)
                p.x += int(dist)
            case 1:
                fmt.Printf("#%s = D %d\n",color,dist)
                p.y += int(dist)
            case 2:
                fmt.Printf("#%s = L %d\n",color,dist)
                p.x -= int(dist)
            case 3:
                fmt.Printf("#%s = U %d\n",color,dist)
                p.y -= int(dist)
        }
        nodes = append(nodes,p)
    }
    //fmt.Println("Got out",nodes)
    var result2 = findAreaShoelace(nodes)
    fmt.Println("Part 2 Hexdecimal:",result2+1)
    
}

func findAreaShoelace(nodes []point) int {
	area := 0

	for i := 0; i < len(nodes); i++ {
		pointA, pointB := nodes[i], nodes[(i+1)%(len(nodes))]
		area += (pointA.x * pointB.y) - (pointB.x * pointA.y) + max(abs(pointA.x-pointB.x), abs(pointA.y-pointB.y))
	}

	return area / 2
}

func floodFill(trench map[point]string, p point) {
	trench[p] = ""
	for _, pp := range [4]point{{p.x,p.y-1},{p.x,p.y+1},{p.x-1,p.y},{p.x+1,p.y}}{
		if _, contains := trench[pp]; !contains {
			floodFill(trench, pp)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
