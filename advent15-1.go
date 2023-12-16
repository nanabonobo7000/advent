package main

import("fmt";"bufio";"os";"strings";"math")

func main(){
    var summer float64
    summer = 0
    cnt:=0
//     iiii := 1
    
    fmt.Println("Input HASH:")
    
    
    //NEED TO USE CAT INTO PROGRAM ELSE IT ENDS STREAM EARLY
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString(',')
        s=strings.ReplaceAll(s,",","")
        s=strings.TrimSpace(s)
        if(s=="end"||s==""){break}
        
        
//         chars := make(map[int]rune)
        
//         for _,c := range s{
//             if chars[int(c)]==0{
//                 chars[int(c)]=c
//             }else{
//                 if((chars[int(c)])!=c){
//                     fmt.Println(string(c),"is not equal to",string(chars[int(c)]),"please change the overlap")
//                 }
//             }
//         }
        
        
        
//         if(s[0]==','){
//             s1 := []byte(s)
//             s1[0]=' '
//             s = strings.TrimSpace(string(s1))
//         }
        //ss := strings.Split(s,",")
        
        
//         for i,c := range chars{
//             fmt.Printf("%d:%c ",i,c)
//         }
//         fmt.Printf("\n")
        
        //for _,sss := range ss{
            v := hash(s)
            summer = summer + float64(v)
            cnt++
            fmt.Println(cnt,s,v,summer)
        //}

    }
    fmt.Println("Final Sum:",summer)
}

func hash(s string) int{
    r := 0
    s = strings.TrimSpace(s)
    for _,c := range s{
        r = r+int(c)
        r = r*17
        r = int(math.Mod(float64(r),float64(256)))
    }
    return r
}
        
