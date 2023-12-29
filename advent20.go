package main

import("fmt";"bufio";"os";"strings")

type gate struct{
    t string
    status int
    routes []string
    emit int
    inputs map[string]int
}

func main(){
    fmt.Println("Input coordinates:")
    gates := make(map[string]gate)
    init_ := make(map[string]int)
    
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        s=strings.ReplaceAll(s,",","")
        if(s=="end"){break}
        ss := strings.Split(s," ")
        name := "broadcaster"
        t := "start"
        if(ss[0]!="broadcaster"){
            name = string(ss[0][1:])
            t = string(ss[0][0])
        }
        r := ss[2:]
        gates[name] = gate{t,0,r,0,init_}
    }
    
    //fmt.Println(gates)
    
    for i,g := range gates{
        if(g.t=="&"){
            //fmt.Println("Adding",i,g)
            m := make(map[string]int)
            for ii,gg := range gates{
                for _,ggg := range gg.routes{
                    if(ggg==i){
                        //fmt.Println("Found",ii,gg)
                        m[ii] = 0
                    }
                }
            }
            g.inputs = m
            gates[i] = g
        }
    }
    
    
    //fmt.Println(gates)
    
    
    
    
    
    var emitting []string
    lows:=0
    highs:=0
    rx := 0
    
    for i:=0;i<1000;i++{
        lows++
        emitting = append(emitting,"broadcaster")
        for(len(emitting)>0){
            rx = 0
            for _,i := range emitting{
                pulse := gates[i].emit
                for _,j := range gates[i].routes{
                    if(pulse==1){
                        highs++
                    }else if(pulse==0){
                        lows++
                        if(j=="rx"){rx++}
                    }
                    t := gates[j].t
                    status := gates[j].status
                    g := gates[j]
                    
                    
                    switch t{
                        case "%":
                            if(pulse==1){
                                
                            }else if(pulse==0&&status==1){
                                g.status = 0
                                g.emit = 0
                                gates[j] = g
                                emitting = append(emitting,j)
                            }else if(pulse==0&&status==0){
                                g.status = 1
                                g.emit = 1
                                gates[j] = g
                                emitting = append(emitting,j)
                            }
                        case "&":
                            ii := g.inputs
                            ii[i] = pulse
                            g.inputs = ii
                            gates[j] = g
                            sig := 0
                            for _,in := range gates[j].inputs{
                                if(in==0){
                                    sig=1
                                }
                            }
                            g.emit = sig
                            gates[j] = g
                            emitting = append(emitting,j)
                    }
                }
                emitting = emitting[1:]
            }
            if(rx==1){
                fmt.Println("rx received only 1 pulse at:",i)
                break
            }
        }
    }
    
    
    fmt.Println(lows,highs,lows*highs)
}
