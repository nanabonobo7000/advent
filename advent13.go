package main

import("fmt";"bufio";"os";"strings")

func main(){
    fmt.Println("Input coordinates:")
    
    reader:= bufio.NewReader(os.Stdin)
    var rows []string
    var cols []string
    init := 1
    mirror := 0
    mirtot := 0
    
    
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        

        if(s!=""){
            rows = append(rows,s)
            if(init==1){
                for _,_ =range s{
                    cols = append(cols,"")
                }
//                 for i,c := range cols{
//                     fmt.Println(i,c)
//                 }
                init=0
            }       
            for i,c := range s{
                //fmt.Println(i,string(c))
                cols[i] = strings.Join([]string{cols[i],string(c)},"")
            }
        }else if(init==0){
            var end int
            
            //fmt.Println("ROWS")
            var r string
            for i:=0;i<len(rows)-1;i++{
                r=rows[i]
                good := 1
                if(r==rows[i+1]){
                    if(len(rows)/2>=(i+1)){
                        end = 0
                    }else{end = i-(len(rows)-1-(i+1))}
                    //fmt.Println("Row End is",end)
                    for j:=(i-1);j>=end;j--{
                        //fmt.Println("Rows ",j,"against",(i+1)+(i-j))
                        if(rows[j]!=rows[(i+1)+(i-j)]){
                            good = 0
                        }
                    }
                    if(good==1){
                        mirror = (i+1)*100
                        mirtot = mirtot+mirror
                        fmt.Println("ROWS ABOVE:",mirror)
                        break
                    }
                }
            }
            
            
            //fmt.Println("COLS")
            var c string
            for i:=0;i<len(cols)-1;i++{
                c=cols[i]
                good := 1
                if(c==cols[i+1]){
                    if(len(cols)/2>=(i+1)){
                        end = 0
                    }else{end = i-(len(cols)-1-(i+1))}
                    //fmt.Println("Col End is",end)
                    for j:=(i-1);j>=end;j--{
                        //fmt.Println("Cols ",j,"against",(i+1)+(i-j))
                        if(cols[j]!=cols[(i+1)+(i-j)]){
                            good = 0
                        }
                    }
                    if(good==1){
                        mirror = i+1
                        mirtot = mirtot+mirror
                        fmt.Println("COLS LEFT:",mirror)
                        break
                    }
                }
            }
            
            rows = []string{}
            cols = []string{}
            init = 1
            mirror = 0
            
        }
        
    }
    fmt.Println("Total results:",mirtot)
}
