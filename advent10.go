package main

import("fmt";"bufio";"os";"strings")

type pipe struct{
    c rune
    north int
    south int
    east int
    west int
}

var count int

func main(){
    fmt.Println("Input coordinates:")
    start := 1
    row := 0
    sss := 0
    var sr, sc int
    var pipes [][]pipe
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        if(start==1){
            for i:=0;i<len(s);i++{
                pipes = append(pipes,make([]pipe,len(s)))
            }
            sss = len(s)-1
            start = 0
        }
        for col,c := range s{
            pipes[row][col] = piper(c)
             if(c=='S'){
                 sr = row
                 sc = col
             }
        }
        row++
    }
    
    find(&pipes,sr,sc,sss)
   
    cc := 0
    var open []int
    
    for _,pp := range pipes[0]{
        if(pp.c!='0'){
            open = append(open,1)
        }else{open=append(open,0)}
    }
    fmt.Println("PIPES",open)
    in := 0
    insider:=0
    
    for i,p := range pipes{
        if(i>0){
            in=0
            for j,pp := range p{   
                if(pp.c>'0'){
                    if(open[j]==1){
                        if(in==0){
                            in=1
                        }else if(in==1){
                            in=1
                        }else if(in==2){
                            in=0
                            cc=cc+insider
                            insider=0
                        }
                    }
                }else if(pp.c=='0'){
                    if(open[j]==1){
                        if(in==0){
                            in=0
                            open[j]=0
                        }else if(in==1){
                            in=2
                            insider++
                            pipes[i][j].c = '1'
                        }else if(in==2){
                            in=2
                            insider++
                            pipes[i][j].c = '1'
                        }
                    }else{
                        if(in==0){
                            in=0
                        }else if(in==1){
                            in=2
                            insider++
                            pipes[i][j].c = '1'
                        }else if(in==2){
                            in=2
                            insider++
                            pipes[i][j].c = '1'
                        }
                    }
                }
            }
            //cc=cc+insider
        }
    }        
        
    for i,p := range pipes{
        fmt.Printf("%d: ",i)
        for _,pp := range p{
            if(pp.c=='1'){
                fmt.Printf("X")
            }else if(pp.c=='0'){
                fmt.Printf("0")
            }else{fmt.Printf(".")}
        }
        fmt.Printf("\n")
    }
    
    fmt.Printf("\nFarthest distance:%d\n",(count/2)+1)
    fmt.Printf("Enclosed count:%d\n",cc)
}

func find(p *[][]pipe, r int, c int, ss int){
    n := (*p)[r][c].north
    s := (*p)[r][c].south
    e := (*p)[r][c].east
    w := (*p)[r][c].west
    
    if(n==1){
        rl := r-1
        //fmt.Println("Searching",r,c,"in",rl,c)
        if(rl>=0){
            if((*p)[rl][c].south==1&&(*p)[rl][c].c!=2){
                (*p)[rl][c].south=0
                //if((*p)[rl][c].c=='0'){return}
                (*p)[rl][c].c = (*p)[r][c].c+1
                count++
                find(p,rl,c,ss)
            }
        }
        (*p)[r][c].north = 0
    }
    if(s==1){
        rl := r+1
        //fmt.Println("Searching",r,c,"in",rl,c)
        if(rl<=ss){
            if((*p)[rl][c].north==1&&(*p)[rl][c].c!=2){
                (*p)[rl][c].north=0
                //if((*p)[rl][c].c=='0'){return}
                (*p)[rl][c].c = (*p)[r][c].c+1
                count++
                find(p,rl,c,ss)
            }
        }
        (*p)[r][c].south = 0
    }
    if(e==1){
        cl := c+1
        //fmt.Println("Searching",r,c,"in",r,cl)
        if(cl<=ss){
            if((*p)[r][cl].west==1&&(*p)[r][cl].c!=2){
                (*p)[r][cl].west=0
                //if((*p)[r][cl].c=='0'){return}
                (*p)[r][cl].c = (*p)[r][c].c+1
                count++
                find(p,r,cl,ss)
            }
        }
        (*p)[r][c].east = 0
    }
    if(w==1){
        cl := c-1
        //fmt.Println("Searching",r,c,"in",r,cl)
        if(cl>=0){
            if((*p)[r][cl].east==1&&(*p)[r][cl].c!=2){
                (*p)[r][cl].east=0
                //if((*p)[r][cl].c=='0'){return}
                (*p)[r][cl].c = (*p)[r][c].c+1
                count++
                find(p,r,cl,ss)
            }
        }
        (*p)[r][c].west = 0
    }
}


func piper(c rune) pipe{
    if(c=='|'){
        return pipe{'0',1,1,0,0}
    }else if(c=='-'){
        return pipe{'0',0,0,1,1}
    }else if(c=='L'){
        return pipe{'0',1,0,1,0}
    }else if(c=='J'){
        return pipe{'0',1,0,0,1}
    }else if(c=='7'){
        return pipe{'0',0,1,0,1}
    }else if(c=='F'){
        return pipe{'0',0,1,1,0}
    }else if(c=='.'){
        return pipe{'0',0,0,0,0}
    }else if(c=='S'){
        return pipe{'2',1,1,1,1}
    }
    return pipe{0,0,0,0,0}
}


