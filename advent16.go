package main

import("fmt";"bufio";"os";"strings")

type mir struct{
    ch rune
    enrg int
    north int
    south int
    east int
    west int
}

func main(){
    fmt.Println("Input coordinates:")
    
    var m [][]mir
    var m1 [][]mir
    rr:=0
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        m = append(m,make([]mir,len(s)))
        m1 = append(m1,make([]mir,len(s)))
        
        for i,c := range s{
            m[rr][i]=mir{ch:c,enrg:0,north:0,south:0,east:0,west:0}
        }
        rr++
    }
    
    enermax := 0
    for i,_ := range m1{
        copy(m1[i],m[i])
    }
    energized := 0
    
    for _,mm := range m1{
        for _,mmm := range mm{
                fmt.Printf("%c",mmm.ch)
        }
        fmt.Printf("\n")
    }
    
    
    for r:=0;r<len(m);r++{
        for i,_ := range m1{
            copy(m1[i],m[i])
        }
        energized = 0
        east(r,0,&m1)
        for _,m1m := range m1{
            for _,m1mm := range m1m{
                if(m1mm.enrg==1){
                    //fmt.Printf("#")
                    energized++
                }else{
                    //fmt.Printf(".")
                }
            }
            //fmt.Printf("\n")
        }
        fmt.Printf("[r%d][c%d]:%d\n",r,0,energized)
        if(energized>enermax){enermax=energized}
        for i,_ := range m1{
            copy(m1[i],m[i])
        }
        energized = 0
        west(r,(len(m[0]))-1,&m1)
        for _,m1m := range m1{
            for _,m1mm := range m1m{
                if(m1mm.enrg==1){
                    //fmt.Printf("#")
                    energized++
                }else{
                    //fmt.Printf(".")
                }
            }
            //fmt.Printf("\n")
        }
        fmt.Printf("[r%d][c%d]:%d\n",r,len(m[0])-1,energized)
        if(energized>enermax){enermax=energized}
    }
    
    for c:=0;c<len(m[0]);c++{
        for i,_ := range m1{
            copy(m1[i],m[i])
        }
        energized = 0
        south(0,c,&m1)
        for _,m1m := range m1{
            for _,m1mm := range m1m{
                if(m1mm.enrg==1){
                    //fmt.Printf("#")
                    energized++
                }else{
                    //fmt.Printf(".")
                }
            }
            //fmt.Printf("\n")
        }
        fmt.Printf("[r%d][c%d]:%d\n",0,c,energized)
        if(energized>enermax){enermax=energized}
        for i,_ := range m1{
            copy(m1[i],m[i])
        }
        energized = 0
        north(len(m)-1,c,&m1)
        for _,m1m := range m1{
            for _,m1mm := range m1m{
                if(m1mm.enrg==1){
                    //fmt.Printf("#")
                    energized++
                }else{
                    //fmt.Printf(".")
                }
            }
            //fmt.Printf("\n")
        }
        fmt.Printf("[r%d][c%d]:%d\n",len(m)-1,c,energized)
        if(energized>enermax){enermax=energized}
    }
    
    
    
    east(0,0,&m)
    energized=0
    
    for _,mm := range m{
        for _,mmm := range mm{
            if(mmm.enrg==1){
                //fmt.Printf("#")
                energized++
            }else{
                //fmt.Printf(".")
            }
        }
        //fmt.Printf("\n")
    }
    fmt.Println("Energized [0,0]:",energized)
    fmt.Println("Max Energy:",enermax)
}

func east(r int, c int, m *[][]mir){
    char := (*m)[r][c].ch
    (*m)[r][c].enrg = 1
    if((*m)[r][c].east==1){
        return
    }
    (*m)[r][c].east = 1
    if(char=='\\'){
        if(r+1<len(*m)){
            south(r+1,c,m)
        }
    }else if(char=='/'){
        if(r-1>=0){
            north(r-1,c,m)
        }
    }else if(char=='|'){
        if(r+1<len(*m)){
            south(r+1,c,m)
        }
        if(r-1>=0){
            north(r-1,c,m)
        }
    }else if(char=='.'||char=='-'){
        if(c+1<len((*m)[r])){
            (*m)[r][c].east = 1
            east(r,c+1,m)
        }
    }
}

func north(r int, c int, m *[][]mir){
    char := (*m)[r][c].ch
    (*m)[r][c].enrg = 1
    if((*m)[r][c].north==1){
        return
    }
    (*m)[r][c].north = 1
    if(char=='\\'){
        if(c-1>=0){
            west(r,c-1,m)
        }
    }else if(char=='/'){
        if(c+1<len((*m)[r])){
            east(r,c+1,m)
        }
    }else if(char=='-'){
        if(c+1<len((*m)[r])){
            east(r,c+1,m)
        }
        if(c-1>=0){
            west(r,c-1,m)
        }
    }else if(char=='.'||char=='|'){
        if(r-1>=0){
            north(r-1,c,m)
        }
    }
}

func west(r int, c int, m *[][]mir){
    char := (*m)[r][c].ch
    (*m)[r][c].enrg = 1
    if((*m)[r][c].west==1){
        return
    }
    (*m)[r][c].west = 1
    if(char=='/'){
        if(r+1<len(*m)){
            south(r+1,c,m)
        }
    }else if(char=='\\'){
        if(r-1>=0){
            north(r-1,c,m)
        }
    }else if(char=='|'){
        if(r+1<len(*m)){
            south(r+1,c,m)
        }
        if(r-1>=0){
            north(r-1,c,m)
        }
    }else if(char=='.'||char=='-'){
        if(c-1>=0){
            west(r,c-1,m)
        }
    }
}

func south(r int, c int, m *[][]mir){
    char := (*m)[r][c].ch
    (*m)[r][c].enrg = 1
    if((*m)[r][c].south==1){
        return
    }
    (*m)[r][c].south = 1
    if(char=='/'){
        if(c-1>=0){
            west(r,c-1,m)
        }
    }else if(char=='\\'){
        if(c+1<len((*m)[r])){
            east(r,c+1,m)
        }
    }else if(char=='-'){
        if(c+1<len((*m)[r])){
            east(r,c+1,m)
        }
        if(c-1>=0){
            west(r,c-1,m)
        }
    }else if(char=='.'||char=='|'){
        if(r+1<len(*m)){
            south(r+1,c,m)
        }
    }
}
