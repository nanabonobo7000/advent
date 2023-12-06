package main

import("fmt";"bufio";"os";"strings";"strconv")

type val struct{
    sum, r1, c1, c2 int
}

func main(){
    fmt.Println("Input Engine Matrix:")
    
    var sarr []string
    var narr []val
    rows:=-1
    reader := bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        sarr = append(sarr,s)
        rows=rows+1
    }

    for row,ss := range sarr{
        inside:=0
        var asset = val{sum:0,r1:row,c1:-1,c2:-1}
        for col,c := range ss{
            intv,err := strconv.Atoi(string(c))
            if err==nil {
                //fmt.Printf("r%dc%d ",row,col)
                if(inside==0){asset.c1=col}
                asset.sum = (asset.sum*10)+intv
                inside=1
                
                
            }else{
                //fmt.Printf("-")
                if(inside==1){
                    asset.c2 = col-1
                    narr = append(narr,asset)
                }
                inside=0
                asset = val{sum:0,r1:row,c1:-1,c2:-1}
                
            }
        }
        //fmt.Printf("\n")
        if(inside==1){
            asset.c2 = len(ss)-1
            narr = append(narr,asset)
        }
    }
    
//     type val struct{
//         sum, r1, c1, c2 int
//     }
    
    summer := 0
    grsum := 0
    for row,ss := range sarr{
        for col,c := range ss{
            _,err := strconv.Atoi(string(c))
            if(err!=nil&&c!='.'){
                gr:=1
                adj:=0
                //fmt.Println(string(c),row,col)
                for _,part := range narr{
                    if(part.r1==row-1||part.r1==row||part.r1==row+1){
                        if(part.c1==col-1||part.c1==col||part.c1==col+1||part.c2==col-1||part.c2==col||part.c2==col+1){
                            summer=summer+part.sum
                            //fmt.Println(string(c),row,col,part)
                            adj=adj+1
                            gr=gr*part.sum
                        }
                    }
                }
                if(adj==2&&c=='*'){grsum=grsum+gr}
            }
        }
    }
    
    fmt.Println("Total Enginer Schema:",summer)
    fmt.Println("Gear Ratio Sum:",grsum)
    //fmt.Println(narr)
}
