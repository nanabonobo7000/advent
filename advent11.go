package main

import("fmt";"bufio";"os";"strings")

type planet struct{
    loc int
    space int
}

type loc struct{
    y int
    x int
}

var rows, cols = 0, 0

func main(){
    fmt.Println("Input space map:")
    
    var uni [][]planet
    var planets = make(map[int]loc)
    pcount := 1
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        uni = append(uni,make([]planet,len(s)))
        for col,c := range s{
            if(c=='#'){
                uni[rows][col] = planet{pcount,1}
                planets[pcount] = loc{rows,col}
                pcount++
            }else{uni[rows][col] = planet{0,1}}
        }
        rows++
        cols = len(s)
    }
    
    var xs []int
    var ys []int
    
    find:=0
    for r:=0;r<rows;r++{
        find=0
        for c:=0;c<cols;c++{
            if(uni[r][c].loc>0){find=1}
        }
        if(find==0){for c:=0;c<cols;c++{uni[r][c].space=uni[r][c].space*1000000}}
    }
    for c:=0;c<cols;c++{
        find=0
        for r:=0;r<rows;r++{
            if(uni[r][c].loc>0){find=1}
        }
        if(find==0){for r:=0;r<rows;r++{uni[r][c].space=uni[r][c].space*1000000}}
    }
    
    xs = append(xs,uni[0][0].space)
    ys = append(ys,uni[0][0].space)
    for r:=1;r<rows;r++{
        ys = append(ys,uni[r][0].space+ys[r-1])
    }
    for c:=1;c<cols;c++{
        xs = append(xs,uni[0][c].space+xs[c-1])
    }
    
    fd := 0
    
    //fmt.Println(
    for i:=1;i<(len(planets));i++{
        fmt.Printf("Planet %d [%d,%d]: ",i,planets[i].x,planets[i].y)
        x1 := planets[i].x
        y1 := planets[i].y
        for j:=i+1;j<=(len(planets));j++{
            x2 := planets[j].x
            y2 := planets[j].y
            
            dplan:=0
            if(x1<x2){            
                for x:=x1+1;x<=x2;x=x+1{
                    dplan = dplan + uni[y1][x].space
                    //fmt.Printf("%d ",dplan)
                }
            }else if(x2<x1){
                for x:=x1-1;x>=x2;x=x-1{
                    dplan = dplan + uni[y1][x].space
                    //fmt.Printf("%d ",dplan)
                }
            }
            for y:=y1+1;y<=y2;y++{
                dplan = dplan + uni[y][x2].space
                //fmt.Printf("%d ",dplan)
            }
            fd = fd+dplan
            fmt.Printf("\n%d[%d]\n",j,dplan)
        }
        fmt.Printf("\n")
    }
    
    
    for _,p := range uni{
        for _,pp := range p{
            fmt.Printf("%d",pp.space)
        }
        fmt.Printf("\n")
    }
    fmt.Printf("\n")
    
    
    fmt.Println("X:",xs)
    fmt.Println("Y:",ys)
    fmt.Println(len(planets))
    
    fmt.Println("Distance:",fd)
}
