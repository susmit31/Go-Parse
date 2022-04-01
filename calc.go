package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var OPS []string = strings.Split("+ - * /", " ")
var NUMS = srange(0,10)

func main(){
	var inpstr string

	//---------------------------------- 
	// Accepting full input string
	//----------------------------------
	if !(len(os.Args) > 1){
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inpstr = scanner.Text()
	} else {
		inpstr = arr2str(os.Args[2:len(os.Args)])
	}

	//-----------------------------------
	// Tokenizing the input string
	//-----------------------------------
	parsed := make_ast(inpstr)

	parsed.traverse()
	fmt.Println()
	fmt.Println(parsed.eval())
}


type Node struct{
	parent *Node
	children []*Node
	operators []string
	content string
}

func (n *Node) isleaf() bool{
	for _, child := range n.children{
		if child != nil {
			return false
		}
	}
	return true
}

func (n *Node) traverse() {
	if n.isleaf() {
		fmt.Printf(" %s ", n.content)
	} else {
		fmt.Print(" (")
		for _, child := range n.children{
			if child != nil {
				child.traverse()
			}
		}
		fmt.Printf(" # ")
		for _, op := range n.operators{
			fmt.Printf("%s ",op)
		}
		fmt.Print(") ")
	}
}

func (n *Node) eval() float64{
	var val float64 = float64(0)
	if n.isleaf() {
		//fmt.Sscanf(n.content, "%f", &val)
		v, _ := strconv.ParseFloat(n.content, 64)
		val = v
		return val
	}

	// Addition ; modify later to include operation
	ops := n.operators
	operands := n.children

	val += operands[0].eval()
	for i, op := range ops {
		if op == "+"{
			val += operands[i+1].eval() 
		} else if op == "-" {
			val -= operands[i+1].eval()
		} else if op == "*" {
			val *= operands[i+1].eval()
		} else if op == "/" {
			val /= operands[i+1].eval()
		}
	}
	return val
}

func make_ast(expr string) *Node{
	var root_node *Node = new(Node)
	var ops []string
	var operands []*Node
	var cursor int = 0

	for cursor<len(expr) {
		character := string(expr[cursor])
		if isin(append(NUMS,"."), character) {
			num :=  make_num(expr, cursor)
			children :=  []*Node{nil,nil,nil}
			operands = append(operands, 
						&Node{parent: root_node, 
							content: num,
							children:children, 
							operators: []string{""}})
			cursor += len(num)
		} else if "(" == character {
			subtree := make_ast(expr[cursor+1:len(expr)])
			operands = append(operands, subtree)
			cursor += len(subtree.content)+1
		} else if ")" == character {
			expr = expr[0:(cursor+1)]
			break
		} else if isin(OPS, character) {
			ops = append(ops, character)
			cursor++
		} else {
			cursor++
		}	
	}
	*root_node = Node{parent:nil, 
					children: operands, 
					content: expr, 
					operators: ops}
	return root_node
}

func make_num(expr string, pos int) string{
	var num string = ""
	for i:= pos; i < len(expr) ; i++ {
		character := string(expr[i])
		if isin(append(NUMS, "."), character) {
			num += character
		} else {
			break
		}
	}
	return num
}

func srange(start, end int) []string{
	var arr []string
	for i:=start; i<end; i++ {
		arr = append(arr, strconv.Itoa(int(i)))
	}
	return arr
}

func str2arr(str string) []string{
	var arr []string
	for _, ch := range str{
		arr = append(arr, string(ch))
	}
	return arr
}

func isin[T comparable] (arr []T, el T) bool{
	if indexof(arr, el, 0) != -1 {
		return true
	} else {
		return false
	}
}

func indexof[T comparable] (arr []T, el T, startat int) int{
	if startat > len(arr)-1{
		return -2
	}
	for i:=startat; i<int(len(arr)); i++{
		if arr[i] == el {
			return i
		}
	}
	return -1
}

func arr2str(arr []string) string{
	var result string 
	for _, str := range arr {
		result += str
	}
	return result
}
