package main

import("fmt";"bufio";"os";"strings";"strconv")

func main(){
    fmt.Println("Input series:")
    
    sum := 0
    sum2:= 0
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        ss := strings.Split(s," ")
        
        var il1 []int
        for _,i := range ss{
            n,_ := strconv.Atoi(i)
            il1 = append(il1,n)
        }
        l := len(il1)+2
        //fmt.Println(l)
        
        var il [][]int
        for i:=1;i<=l-2;i++{
            il = append(il,make([]int,l))
        }
        
        for ind,inti := range il1{
            il[0][ind+1] = inti
        }
        
        
        for i:=1; i<=l-4; i++{
            for j:=1; j<=l-1-i; j++{
                il[i][j] = il[i-1][j+1]-il[i-1][j]
            }
        }
        for i:=l-4;i>=0;i--{
            il[i][l-1-i] = il[i+1][l-2-i]+il[i][l-2-i]
        }
        
        for i:=l-4;i>=0;i--{
            il[i][0] = il[i][1]-il[i+1][0]
        }
        
        
        
        sum = sum + il[0][l-1]
        sum2 = sum2 + il[0][0]
         for _,r := range il{
             fmt.Println(r)
         }
    }
    fmt.Println("OASIS READING:",sum)
    fmt.Println("OASIS EXTRAPOLATION:",sum2)
}
