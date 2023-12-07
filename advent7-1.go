package main

import("fmt";"bufio";"os";"strings";"strconv")

type Node struct {
	value int
	cards string
	bet int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

var card = map[rune]int{
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
    'T': 10,
    'J': 11,
    'Q': 12,
    'K': 13,
    'A': 14,
}

var sum int = 0
var rank int = 1

func main(){
    start:=1
    var tree BST
    
    fmt.Println("Input hands:")
    reader:= bufio.NewReader(os.Stdin)
    for{
        s,_ := reader.ReadString('\n')
        s=strings.TrimSpace(s)
        if(s=="end"){break}
        
        ss := strings.Split(s," ")
        b,_ := strconv.Atoi(ss[1])
        strength := stren(ss[0])
        if(start==1){
            tree = *New(strength,ss[0],b)
            start=0
        }else{
            Add(&tree, &Node{value:strength,cards:ss[0],bet:b})
            //fmt.Println(ss[0])
        }
    }
    count(tree.root)
    //p(tree.root)
    fmt.Println("Total winnings:",sum)
}

func New(t int, s string, b int) *BST {
	return &BST{&Node{value:t,cards:s,bet:b}}
}

func Add(t *BST, n *Node){
	t.root = add(t.root, n)
}

func add(r *Node, n *Node) *Node{
    if r == nil {
		return n
	}
	if(r.value > n.value){
		r.left = add(r.left,n)
	}else if(r.value < n.value){
		r.right = add(r.right,n)
	}else if(r.value==n.value){
        rootS := r.cards
        newS := n.cards
        found := 0
        for i,char := range rootS{
            //fmt.Println(card[char],card[rune(newS[i])])
            if(card[char]<card[rune(newS[i])]){
                found = 1
                r.right = add(r.right,n)
                break
            }else if(card[char]>card[rune(newS[i])]){
                found = 1
                r.left = add(r.left,n)
                break
            }
        }
        if(found == 0){r.left = add(r.left,n)}
    }
	return r
}

func p(node *Node){
    fmt.Println(node)
    if(node.left!=nil){
        fmt.Printf("L: ")
        p(node.left)
    }
    if(node.right!=nil){
        fmt.Printf("R: ")
        p(node.right)
    }
}

func count(node *Node){
    if(node.left!=nil){
        count(node.left)
    }
    if(node.left==nil){
        fmt.Println(node.cards, node.value, rank, (node.bet*rank), sum, sum + (node.bet*rank))
        sum = sum + (node.bet*rank)
        rank = rank + 1
    }
    if(node.left!=nil&&node.right==nil){
        fmt.Println(node.cards, node.value, rank, (node.bet*rank), sum, sum + (node.bet*rank))
        sum = sum + (node.bet*rank)
        rank = rank + 1
    }
    if(node.left!=nil&&node.right!=nil){
        fmt.Println(node.cards, node.value, rank, (node.bet*rank), sum, sum + (node.bet*rank))
        sum = sum + (node.bet*rank)
        rank = rank + 1
    }
    if(node.right!=nil){
        count(node.right)
    }
}
    

func stren(s string) int{
    var pow = make(map[rune]int)
    for _,char := range s{
        pow[char] = pow[char]+1
    }
    twos, threes := 0,0
    for _,j := range pow{
        if(j==2){twos=twos+1}
        if(j==3){threes=threes+1}
        if(j==4){return 6}
        if(j==5){return 7}
    }
    if(twos==1&&threes==1){
        return 5
    }else if(threes==1){
        return 4
    }else if(twos==2){
        return 3
    }else if(twos==1){
        return 2
    }else{
        return 1
    }
}
