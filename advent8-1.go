package main

import("fmt";"bufio";"os";"strings")

type dirs struct{
    left string
    right string
}

func main(){
    fmt.Println("Input map:")
    
    starter:=1
    var dir string
    var directions = make(map[string]dirs)
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        if(starter==1){
            dir=s
            starter=0
        }else if(s!=""){
            s = strings.ReplaceAll(s," ","")
            ss := strings.Split(s,"=")
            key := ss[0]
            sss := strings.ReplaceAll(ss[1],"(","")
            sss = strings.ReplaceAll(sss,")","")
            vals := strings.Split(sss,",")
            directions[key] = dirs{left:vals[0],right:vals[1]}
        }
    }
    dest := "ZZZ"
    start := "AAA"
    movements := 0
    fmt.Println("START:",start)
    for{
        if(start!=dest){
            fmt.Printf("%s ",start)
            for _,c := range dir{
                if(c=='L'){
                    start=directions[start].left
                    fmt.Printf("L:%s ",start)
                    movements = movements+1
                }else if(c=='R'){
                    start=directions[start].right
                    fmt.Printf("R:%s ",start)
                    movements = movements+1
                }
            }
            fmt.Printf("%d\n",movements)
        }else{break}
    }
    fmt.Println("DEST:",start)
    fmt.Println("Total moves from AAA to ZZZ:",movements)
}
