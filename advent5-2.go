package main

import("fmt";"bufio";"os";"strings";"strconv")

func main(){
    fmt.Println("Input Almanac:")
    
    start := 1
    seedstart :=0
    soilstart :=0
    fertstart :=0
    waterstart :=0
    lightstart :=0
    tempstart :=0
    humidstart :=0
    
    var soils [][]int
    var ferts [][]int
    var wats [][]int
    var lights [][]int
    var temps [][]int
    var humids [][]int
    var locs [][]int
    
    var seednums []string
    
    reader := bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)        
        if(s=="end"){break}
        
        //Collect seeds information
        if(start==1){
            s=strings.ReplaceAll(s,"seeds: ","")
            seednums=strings.Split(s," ")
            //fmt.Println(seednums,len(seednums))
            start=0
        //Start seed-to-soil calculations
        }else if(s=="seed-to-soil map:"){
            seedstart=1
        }else if(seedstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                soils = append(soils,[]int{a, b, c})
            }else{seedstart=0}
        //Start soil-to-fertilizer calculations
        }else if(s=="soil-to-fertilizer map:"){
            soilstart=1
        }else if(soilstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                ferts = append(ferts,[]int{a, b, c})
            }else{soilstart=0}
        //Start fertilizer-to-water calculations
        }else if(s=="fertilizer-to-water map:"){
            fertstart=1
        }else if(fertstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                wats = append(wats,[]int{a, b, c})
            }else{fertstart=0}
        //Start water-to-light calculations
        }else if(s=="water-to-light map:"){
            waterstart=1
        }else if(waterstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                lights = append(lights,[]int{a, b, c})
            }else{waterstart=0}
        //Start light-to-temp calculations
        }else if(s=="light-to-temperature map:"){
            lightstart=1
        }else if(lightstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                temps = append(temps,[]int{a, b, c})
            }else{lightstart=0}
        //Start temp-to-humid calculations
        }else if(s=="temperature-to-humidity map:"){
            tempstart=1
        }else if(tempstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                humids = append(humids,[]int{a, b, c})
            }else{tempstart=0}
        //Start humid-to-location calculations
        }else if(s=="humidity-to-location map:"){
            humidstart=1
        }else if(humidstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                locs = append(locs,[]int{a, b, c})
            }else{humidstart=0}
        }
    }
    fmt.Printf("\n----------------------CALCULATING----------------------\n")
    counter:=0
    terminal:=len(seednums)-1
    minloc := -1
    totalkeys:=0
    fmt.Println("Total seeds to run:",(terminal+1)/2)
        for{
            if(counter<=terminal){
                nn1,_ := strconv.Atoi(seednums[counter])
                nn2,_ := strconv.Atoi(seednums[counter+1])
                totalkeys=totalkeys+nn2
                fmt.Println("Seed",counter/2+1,":",nn1,"\tTo calculate:",nn2,"\tSum:",totalkeys)
                nn:=nn1
                for{
                    if(nn<=nn1+nn2-1){
                        loc:= toLoc(nn,soils,ferts,wats,lights,temps,humids,locs)
                        if(loc<minloc||minloc==-1){minloc=loc}
                        nn=nn+1
                    }else{break}
                }
                counter=counter+2
            }else{break}
        }
    fmt.Println("Lowest location index:",minloc)
}


// type seed struct{
//     seed, soil, fert, water, light, temp, humid, location int
// }

func toLoc(seed int, soil [][]int, fert [][]int, water [][]int, light [][]int, temp [][]int, humid [][]int, locs [][]int) int{
    loc := seed
    //fmt.Println("seed:",loc)
    for _,a:= range soil{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            loc=(loc+(a[0]-a[1]))
            break
        }
    }
    //fmt.Println("soil:",loc)
    for _,a:= range fert{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            loc=(loc+(a[0]-a[1]))
            break
        }
    }
    //fmt.Println("fert:",loc)
    for _,a:= range water{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            loc=(loc+(a[0]-a[1]))
            break
        }
    }
    //fmt.Println("water:",loc)
    for _,a:= range light{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            loc=(loc+(a[0]-a[1]))
            break
        }
    }
    //fmt.Println("light:",loc,"\n",temp)
    for _,a:= range temp{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            //fmt.Println(loc, a[1], a[1]+a[2]-1)
            loc=(loc+(a[0]-a[1]))
            //fmt.Println(loc, a[0], a[1])
            break
        }
    }
    //fmt.Println("temp:",loc)
    for _,a:= range humid{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            loc=(loc+(a[0]-a[1]))
            break
        }
    }
    //fmt.Println("humidity:",loc)
    for _,a:= range locs{
        if(loc>=a[1]&&loc<(a[1]+a[2])){
            loc=(loc+(a[0]-a[1]))
            break
        }
    }
    //fmt.Println("location:",loc)
    return loc
}






