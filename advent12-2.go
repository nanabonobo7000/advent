package main

import("fmt";"bufio";"os";"strings";"strconv")

//var level = 0

func main(){
    fmt.Println("[This works but it takes infinity years]")
    fmt.Println("Input springs:")
    counttot := 0
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        var nums1 []int
        ss := strings.Split(s," ")
        springs1 := ss[0]
        ss2 := strings.Split(ss[1],",")
        for _,i := range ss2{
            ii,_ := strconv.Atoi(i)
            nums1 = append(nums1,ii)
        }   
        
        var springs = springs1
        var nums []int
        for _,n := range nums1{
            nums = append(nums,n)
        }
        for i:=0;i<4;i++{
            springs = strings.Join([]string{springs,springs1},"?")
            for _,n := range nums1{
                nums = append(nums,n)
            }
        }
        fmt.Println(springs,nums)
        
        
        var sp1 string
        var myst []int
        for i,c := range springs{
            if(c=='?'){
                myst = append(myst,i)
                sp1 = strings.Join([]string{sp1,"#"},"")
            }else{
                sp1 = strings.Join([]string{sp1,string(c)},"")
            }
        }
        var strops []string
        strops = append(strops,sp1)
        
        
        strops = apper(sp1,strops,0,myst)        
        
        //fmt.Println(strops)
        countsame := 0
        
        for _,s := range strops{
            ints := count(s)
            //fmt.Println(s,ints)
            if(len(ints)==len(nums)){
                same := 1
                for i,_ := range ints{
                    if(ints[i]!=nums[i]){
                        same=0
                        break
                    }
                }
                if(same==1){
                    //fmt.Println(s,springs,nums,"They are the same")
                    countsame++
                }else{}//fmt.Println(s,springs,nums,"They are not the same")}
            }
        }
        fmt.Println("Same",countsame)
        counttot = counttot + countsame
    }
    fmt.Println("Tot Same",counttot)
}


func count(s string) []int{
    var ints []int
    counter:=0
    for _,c := range s{
        if(c=='#'){
            counter++
        }else if(c=='.'){
            if(counter>0){ints = append(ints,counter)}
            counter=0
        }
    }
    if(s[len(s)-1]=='#'){ints=append(ints,counter)}
    return ints
}

func apper(sp1 string, ss []string, ii int, myst []int) []string{
    var sp string
    for i:=ii;i<len(myst);i++{
        sp2 := sp1
        sp = ""
        focus := myst[i]
//          for z:=0;z<level;z++{
//              fmt.Printf("\t")
//          }
//          fmt.Printf("%d[%d]\t",level,focus)
        for j,c := range sp2{
            if(j!=focus){
                sp = strings.Join([]string{sp,string(c)},"")
            }else{
                sp = strings.Join([]string{sp,"."},"")
            }
        }
        //fmt.Printf("%s\n",sp)
        ss = append(ss,sp)
        if(i+1<len(myst)){
            //level++
            ss = apper(sp,ss,i+1,myst)
            //level--
        }
    }
    return ss
}
