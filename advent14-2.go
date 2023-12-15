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
    
    cnt := 0
    
    for cycle:=0;cycle<1000000000;cycle++{
        sat = north(sat)
        sat = west(sat)
        sat = south(sat)
        sat = east(sat)
        cnt++
    }
    
    //fmt.Println(cnt)
    
    //fmt.Println(len(sat))
    w := len(sat)
    weight:=0
    
    for _,ss := range sat{
        fmt.Println(ss)
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

func south(ss []string) []string{
    sl := len(ss)-1
    for i:=sl;i>=0;i--{
        for j,c:= range ss[i]{
            if(c=='O'){
                l := i
                for k:=i+1;k<=sl;k++{
                    if(ss[k][j]=='.'){
                        b := []byte(ss[k])
                        b[j] = 'O'
                        ss[k] = string(b)
                        b = []byte(ss[l])
                        b[j] = '.'
                        ss[l] = string(b)
                        l = l+1
                    }else{
                        break
                    }
                }
            }
        }
    }
    return ss
}

func west(ss []string) []string{
    sl := len(ss[0])-1
    for i,s := range ss{
        for j:=0;j<=sl;j++{
            if(s[j]=='O'){
                l := j
                for k:=j-1;k>=0;k--{
                    if(ss[i][k]=='.'){
                        b := []byte(ss[i])
                        b[k] = 'O'
                        b[l] = '.'
                        ss[i] = string(b)
                        l=l-1
                    }else{
                        break
                    }
                }
            }
        }
    }
    return ss
}

func east(ss []string) []string{
    sl := len(ss[0])-1
    for i,s := range ss{
        for j:=sl;j>=0;j--{
            if(s[j]=='O'){
                l := j
                for k:=j+1;k<=sl;k++{
                    if(ss[i][k]=='.'){
                        b := []byte(ss[i])
                        b[k] = 'O'
                        b[l] = '.'
                        ss[i] = string(b)
                        l=l+1
                    }else{
                        break
                    }
                }
            }
        }
    }
    return ss
}










