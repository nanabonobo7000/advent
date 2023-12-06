package main

import("fmt";"bufio";"os";"strings";"strconv")

type seed struct{
    seed, soil, fert, water, light, temp, humid, location int
}

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
    
    var seeds []seed
    
    reader := bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)        
        if(s=="end"){break}
        
        //Collect seeds information
        if(start==1){
            s=strings.ReplaceAll(s,"seeds: ","")
            seednums:=strings.Split(s," ")
            //fmt.Println(seednums,len(seednums))
            counter:=0
            terminal:=len(seednums)-1
            for{
                if(counter<=terminal){
                    //fmt.Println("l1:",counter,terminal)
                    nn1,_ := strconv.Atoi(seednums[counter])
                    nn2,_ := strconv.Atoi(seednums[counter+1])
                    nn:=nn1
                    for{
                        if(nn<=nn1+nn2-1){
                            //fmt.Println("\tl2:",nn,nn1+nn2-1)
                            seeds = append(seeds,seed{seed:nn,soil:nn,fert:nn,water:nn,light:nn,temp:nn,humid:nn,location:nn})
                            nn=nn+1
                        }else{break}
                    }
                    counter=counter+2
                }else{break}
            }
            start=0
        //Start seed-to-soil calculations
        }else if(s=="seed-to-soil map:"){
            seedstart=1
        }else if(seedstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.seed>=b&&sss.seed<(b+c)){
                        seeds[i].soil=(sss.seed+(a-b))
                        seeds[i].fert=seeds[i].soil
                        seeds[i].water=seeds[i].soil
                        seeds[i].light=seeds[i].soil
                        seeds[i].temp=seeds[i].soil
                        seeds[i].humid=seeds[i].soil
                        seeds[i].location=seeds[i].soil
                    }
                }
            }else{seedstart=0}
        //Start soil-to-fertilizer calculations
        }else if(s=="soil-to-fertilizer map:"){
            soilstart=1
        }else if(soilstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.soil>=b&&sss.soil<(b+c)){
                        seeds[i].fert=(sss.soil+(a-b))
                        seeds[i].water=seeds[i].fert
                        seeds[i].light=seeds[i].fert
                        seeds[i].temp=seeds[i].fert
                        seeds[i].humid=seeds[i].fert
                        seeds[i].location=seeds[i].fert
                    }
                }
            }else{soilstart=0}
        //Start fertilizer-to-water calculations
        }else if(s=="fertilizer-to-water map:"){
            fertstart=1
        }else if(fertstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.fert>=b&&sss.fert<(b+c)){
                        seeds[i].water=(sss.fert+(a-b))
                        seeds[i].light=seeds[i].water
                        seeds[i].temp=seeds[i].water
                        seeds[i].humid=seeds[i].water
                        seeds[i].location=seeds[i].water
                    }
                }
            }else{fertstart=0}
        //Start water-to-light calculations
        }else if(s=="water-to-light map:"){
            waterstart=1
        }else if(waterstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.water>=b&&sss.water<(b+c)){
                        seeds[i].light=(sss.water+(a-b))
                        seeds[i].temp=seeds[i].light
                        seeds[i].humid=seeds[i].light
                        seeds[i].location=seeds[i].light
                    }
                }
            }else{waterstart=0}
        //Start light-to-temp calculations
        }else if(s=="light-to-temperature map:"){
            lightstart=1
        }else if(lightstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.light>=b&&sss.light<(b+c)){
                        seeds[i].temp=(sss.light+(a-b))
                        seeds[i].humid=seeds[i].temp
                        seeds[i].location=seeds[i].temp
                    }
                }
            }else{lightstart=0}
        //Start temp-to-humid calculations
        }else if(s=="temperature-to-humidity map:"){
            tempstart=1
        }else if(tempstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.temp>=b&&sss.temp<(b+c)){
                        seeds[i].humid=(sss.temp+(a-b))
                        seeds[i].location=seeds[i].humid
                    }
                }
            }else{tempstart=0}
        //Start humid-to-location calculations
        }else if(s=="humidity-to-location map:"){
            humidstart=1
        }else if(humidstart==1){
            if(s!=""){
                var a, b, c int
                fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
                for i,sss := range seeds{
                    if(sss.humid>=b&&sss.humid<(b+c)){
                        seeds[i].location=(sss.humid+(a-b))
                    }
                }
            }else{humidstart=0}
        }
    }
    lowcation := seeds[0].location
    for _,locs := range seeds{
        fmt.Println(locs)
        if(locs.location<lowcation){lowcation=locs.location}
    }
    fmt.Println("Lowest location index:",lowcation)
    
}
