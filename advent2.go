package main

import("fmt";"bufio";"os";"strings";"strconv")

func main(){
    fmt.Println("Input games:")
    
    reader:= bufio.NewReader(os.Stdin)
    summer1:=0
    summer2:=0
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        s=strings.Replace(s," ","",-1)
        if(s=="end"){break}
        
        s=strings.Replace(s,"ed","",-1)
        s=strings.Replace(s,"reen","",-1)
        s=strings.Replace(s,"lue","",-1)
        
        ss:=strings.Split(s,":")
        s=ss[1]
        
        games:=strings.Split(s,";")
        
        index,_:=strconv.Atoi((strings.Replace(ss[0],"Game","",-1)))
        //fmt.Println("index:",index)
        
        good:=1
        high := [3]int{0,0,0} //rgb
        for _,c := range games{
            
            num := [3]int{0,0,0} // rgb
            rolls := strings.Split(c,",")
            for _,d := range rolls{
                if(string(d[len(d)-1])=="r"){
                    num[0],_ = strconv.Atoi(strings.TrimSuffix(d,"r"))
                    if(num[0]>high[0]){high[0]=num[0]}
                }else if(string(d[len(d)-1])=="g"){
                    num[1],_ = strconv.Atoi(strings.TrimSuffix(d,"g"))
                    if(num[1]>high[1]){high[1]=num[1]}
                }else if(string(d[len(d)-1])=="b"){
                    num[2],_ = strconv.Atoi(strings.TrimSuffix(d,"b"))
                    if(num[2]>high[2]){high[2]=num[2]}
                }
            }
             if(num[0]>12||num[1]>13||num[2]>14){
                 good=0
             }
            
        }
        power:=high[0]*high[1]*high[2]
        if(good==1){
            summer1 = summer1+index
        }
        summer2 = summer2+power
        //fmt.Println(index,":",high,power)
    }
    fmt.Println("Compatible index sum [12,13,14]:",summer1)
    fmt.Println("High Power sum:",summer2)
}
