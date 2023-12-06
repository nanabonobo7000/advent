package main

import("fmt";"bufio";"os";"strings";"strconv")

type card struct{
    score, numbers int
}

var wscore = 0

func main(){
    fmt.Println("Input Cards:")
    score:=0
    
    var cards []card
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}

        ss:=strings.Split(s,":")
        s=ss[1]
        games:=strings.Split(s,"|")
        winning:=strings.Split(games[0]," ")
        enum:=strings.Split(games[1]," ")
        index,_:=strconv.Atoi((strings.Replace(ss[0],"Card ","",1)))
        
//         fmt.Println(index)
//         fmt.Println("winning:",winning)
//         fmt.Println("elf nums:",enum)
        cs:=0
        cw:=0
        for _,n := range enum{
            for _,w :=range winning{
                wn,_:=strconv.Atoi(w)
                nn,_:=strconv.Atoi(n)
                //fmt.Printf("%d-%d ",wn,nn)
                if(wn==nn&&wn!=0){
                    if(cs==0){cs=cs+1}else{cs=cs*2}
                    cw=cw+1
                    //fmt.Println(index,wn,cs)
                }
            }
        }
        score=score+cs
        fmt.Printf("Card %d\tscore:%d\twins:%d\n",index,cs,cw)
        winrar := card{score:cs, numbers:cw}
        cards  = append(cards,winrar)
    }
    
    for ind,car := range cards{
        scoreme(ind,car,cards)
    }
    
    
    
    
    
    fmt.Println("Non duplicate score:",score)
    fmt.Println("Total Cards:",wscore)
}

func scoreme(index int, c card, cc []card){
    wscore = wscore+1
    i := index+1
    //fmt.Println(i)
    for{
        if(i<=index+c.numbers){scoreme(i,cc[i],cc)}else{break}
        i=i+1
    }
}













