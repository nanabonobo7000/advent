package main

import("fmt";"bufio";"os";"strings")

func main(){
    fmt.Println("Input coordinates:")
    
    reader:= bufio.NewReader(os.Stdin)
    
    var sat []string
    
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        sat = append(sat,s)
    }
    sat = north(sat)
    
    fmt.Println(len(sat))
    w := len(sat)
    weight:=0
    
    for _,ss := range sat{
        for _,c := range ss{
            if(c=='O'){
                weight = weight+w
            }
        }
        w--
    }
    
    fmt.Println("North weight:",weight)
    
}

func north(ss []string) []string{
    for i,s := range ss{
        for j,c:= range s{
            if(c=='O'){
                l := i
                for k:=i-1;k>=0;k--{
                    if(ss[k][j]=='.'){
                        b := []byte(ss[k])
                        b[j] = 'O'
                        ss[k] = string(b)
                        b = []byte(ss[l])
                        b[j] = '.'
                        ss[l] = string(b)
                        l = l-1
                    }else{
                        break
                    }
                }
            }
        }
    }
    return ss
}













