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
    //dest := "ZZZ"
    var starts []string
    var zeds []int
    for f,_ := range directions{
        if(f[2]=='A'){
            starts = append(starts,f)
        }
    }
    movements := 0
    fmt.Println(starts)
    for _,st := range starts{
        movements = 0
        fmt.Printf("%s ",st)
        done:=0
        for{
            if(done==0){
                for _,c := range dir{
                    if(c=='L'){
                        st=directions[st].left
                        movements = movements+1
                        if(st[2]=='Z'){
                            done=1
                            break
                        }
                    }else if(c=='R'){
                        st=directions[st].right
                        movements = movements+1
                        if(st[2]=='Z'){
                            done=1
                            break
                        }
                    }
                }
            }else{break}
        }
        zeds = append(zeds,movements)
        fmt.Printf("%s %d\n",st,movements)
    }
    fmt.Println("I used an online LCM to calculate\nI think this funtion works too but takes a while\nThe answer is likely in the trillions:")
    lstcm := zeds[0]
    for i,z := range zeds{
        if(i!=0){lstcm = lcm(lstcm,z)}
    }
    
    //fmt.Println(starts)
    fmt.Println("Total ghost moves from --A to --Z:",lstcm)
    //629379415 R [LLC BNL KRN VNJ GTK SCN]
    //if you read Erics mind you would know:
    //1, start -> Z == Z -> Z
    //2, there is only 1 Z in each cycle
    //my answer: 13,830,919,117,339
}


func lcm(temp1 int,temp2 int) int{ 
    var lcmnum int = 1
    if(temp1>temp2)    {
        lcmnum=temp1
    }else{
        lcmnum=temp2
    }
    for {        
        if(lcmnum%temp1==0 && lcmnum%temp2==0) {
            fmt.Printf("LCM of %d and %d is %d\n",temp1,temp2,lcmnum)            
            break
        }
        lcmnum++
    }
    return lcmnum
}
