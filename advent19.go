package main

import("fmt";"bufio";"os";"strings";"strconv")

type part struct{
    x int
    m int
    a int
    s int
}

func main(){
    fmt.Println("Input coordinates:")
    
    runner:=0
    funcmap := make(map[string][]string)
    var parts []part
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        if(runner==2){runner=1}
        if(s==""){runner=2}
        
        
        if(runner==0){
            ss:= strings.Split(s,"{")
            f := ss[0]
            dir := strings.ReplaceAll(ss[1],"}","")
            dirs := strings.Split(dir,",")
            for _,d := range dirs{
                funcmap[f] = append(funcmap[f],d)
            }
            
        }else if(runner==1){
            //fmt.Printf("%s\t",s)
            ss := strings.ReplaceAll(s,"}","")
            ss = strings.ReplaceAll(ss,"{","",)
            //fmt.Printf("%s\n",ss)
            ss = strings.ReplaceAll(ss,"=","")
            ss = strings.ReplaceAll(ss,"x","")
            ss = strings.ReplaceAll(ss,"m","")
            ss = strings.ReplaceAll(ss,"a","")
            ss = strings.ReplaceAll(ss,"s","")
            
            sss := strings.Split(ss,",")
            //fmt.Println(sss)
            x1,_ := strconv.Atoi(sss[0])
            m1,_ := strconv.Atoi(sss[1])
            a1,_ := strconv.Atoi(sss[2])
            s1,_ := strconv.Atoi(sss[3])
            
            parts = append(parts,part{x1,m1,a1,s1})
        }
    }
    
    sum := 0
    for _,p := range parts{
        valient := flow(p,funcmap,"in")
        if(valient==-1){valient=0}
        sum = sum + valient
        //fmt.Println(p,valient,sum)
    }
    fmt.Println("Accepted Partval:",sum)
    
    sum2 := grinder(funcmap,"in")
    fmt.Println("[Unfinished] Total distinct combos:",sum2)
    
    
}

func grinder(funcmap map[string][]string,dir string) int{
    return 0
}




func flow(p part, funcmap map[string][]string, dir string) int{
    fin := 0
    partcode := p.x+p.m+p.a+p.s
    f := funcmap[dir]
    //fmt.Println("Checking part",p,f)
    for i:=0;i<(len(f)-1);i++{
        if(fin!=0){
            break
        }
        s := strings.Split(f[i],":")
        xmas := string(s[0][0])
        cond := string(s[0][1])
        val1 := string(s[0][2:])
        val,_ := strconv.Atoi(val1)
        //fmt.Println(dir,i,":If",xmas,"is",cond,"than",val)
        switch xmas{
            case "x":
                switch cond{
                    case ">":
                        if(p.x>val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                    case "<":
                        if(p.x<val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                }
            case "m":
                switch cond{
                    case ">":
                        if(p.m>val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                    case "<":
                        if(p.m<val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                }
            case "a":
                switch cond{
                    case ">":
                        if(p.a>val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                    case "<":
                        if(p.a<val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                }
            case "s":
                switch cond{
                    case ">":
                        if(p.s>val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                    case "<":
                        if(p.s<val){
                            if(s[1]=="A"){
                                return partcode
                            }else if(s[1]=="R"){
                                return -1
                            }else{
                                fin = flow(p,funcmap,s[1])
                            }
                        }
                }
        }
    }
    //fmt.Println("GOT HERE",fin,f[len(f)-1])
    if(fin==0){
        if(f[len(f)-1]=="A"){
            return partcode
        }else if(f[len(f)-1]=="R"){
            return -1
        }else{
            fin = flow(p,funcmap,f[len(f)-1])
        }
    }
    
    if(fin==-1){
        return -1
    }
    //fmt.Println(p,fin)
    return fin
}


