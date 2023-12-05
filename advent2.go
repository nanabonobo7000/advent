package main

import("fmt";"bufio";"os";"strings";"strconv")

func main(){
    fmt.Println("Input coordinates:")
    
    reader:= bufio.NewReader(os.Stdin)
    sum:=0
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSuffix(s,"\n")
        
        
        first := -1
        last := -1
        l := len(s)
        for i,c := range s{
            //fmt.Println(int(c)-48)
            _,err := strconv.Atoi(string(c))
            if err==nil {
                if(first==-1){
                    first = (int(c)-48)
                    last = (int(c)-48)
                }else{last = (int(c)-48)}
            }
            if(c=='o'&&l>=i+3){
                if(string(s[i:i+3])=="one"){
                    if(first==-1){
                        first = 1
                        last = 1
                    }else{last = 1}
                }
            }
            if(c=='t'&&l>=i+3){
                if(string(s[i:i+3])=="two"){
                    if(first==-1){
                        first = 2
                        last = 2
                    }else{last = 2}
                }else if(l>=i+5){ 
                        if(string(s[i:i+5])=="three"){
                            if(first==-1){
                                first = 3
                                last = 3
                            }else{last = 3}
                    }
                }
            }
            if(c=='f'&&l>=i+4){
                if(string(s[i:i+4])=="four"){
                    if(first==-1){
                        first = 4
                        last = 4
                    }else{last = 4}
                }else if(string(s[i:i+4])=="five"){
                    if(first==-1){
                        first = 5
                        last = 5
                    }else{last = 5}
                }
            }
            if(c=='s'&&l>=i+3){
                if(string(s[i:i+3])=="six"){
                    if(first==-1){
                        first = 6
                        last = 6
                    }else{last = 6}
                }else if(l>=i+5){
                        if(string(s[i:i+5])=="seven"){
                            if(first==-1){
                                first = 7
                                last = 7
                            }else{last = 7}
                        }
                }
            }
            if(c=='e'&&l>=i+5){
                if(string(s[i:i+5])=="eight"){
                    if(first==-1){
                        first = 8
                        last = 8
                    }else{last = 8}
                }
            }
            if(c=='n'&&l>=i+4){
                if(string(s[i:i+4])=="nine"){
                    if(first==-1){
                        first = 9
                        last = 9
                    }else{last = 9}
                }
            }
        }
        if(first!=-1){
            sum=sum+(first*10+last)
            fmt.Printf("Thanks here: %s\t%d %d\n",s,(first*10+last),sum)
        }
    
        if(s=="end"){break}
    }
}
