package main

import("fmt";"bufio";"os";"strings";"strconv")

func main(){
    fmt.Println("Input races:")    
    reader:= bufio.NewReader(os.Stdin)
    var ss []string
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        ss = append(ss,s)
    }
    ss[0]=strings.ReplaceAll(ss[0],"Time:      ","")
    ss[1]=strings.ReplaceAll(ss[1],"Distance:  ","")
    timeS := strings.Split(ss[0]," ")
    distS := strings.Split(ss[1]," ")
    
    bigTime,_ := strconv.Atoi(strings.ReplaceAll(ss[0]," ",""))
    bigDist,_ := strconv.Atoi(strings.ReplaceAll(ss[1]," ",""))
    
    var time []int
    for _,t := range timeS{
        tt,_ := strconv.Atoi(t)
        if(tt>0){time = append(time,tt)}
    }
    var dist []int
    for _,t := range distS{
        tt,_ := strconv.Atoi(t)
        if(tt>0){dist = append(dist,tt)}
    }
    
    doftot:=1
    for i,t := range time{
        doftot = doftot*dof(t,dist[i])
    }
    
    fmt.Println("Multiplicative Degrees of Freedom:",doftot)
    fmt.Println("Big Race Answer:",dof(bigTime, bigDist))
    
}

func dof(time int, distance int) int{
    dof := 0
    for t := 0; t < time; t += 1{
        if((t*(time-t))>distance){dof=dof+1}
    }
    return dof
}












