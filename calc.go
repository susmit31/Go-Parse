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
	fmt.Println("keilla dilam")
	var inpstr string

	//---------------------------------- 
	// Accepting full input string
	//----------------------------------
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inpstr = scanner.Text()
	fmt.Printf("lulami laib %s\n", inpstr)

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

/*func (n* Node) isroot() bool{
	if n.parent == nil {
		return true
	} else {
		return false
	}
}*/

func (n *Node) traverse() {
	if n.isleaf() {
		fmt.Printf(" %s ", n.content)
	} else {
		fmt.Print(" [")
		for _, child := range n.children{
			if child != nil {
				child.traverse()
			}
		}
		fmt.Print("] ")
	}
}

func (n *Node) eval() float64{
	var val float64
	if n.isleaf() {
		//fmt.Sscanf(n.content, "%f", &val)
		v, _ := strconv.ParseFloat(n.content, 64)
		val = v
		return val
	}

	// Addition ; modify later to include operation
	op := n.children[len(n.children)-1]
	operands := n.children[0:(len(n.children)-1)]
	
	if op.content == "+"{
		val = float64(0)
		for _, operand := range operands{
			val += operand.eval()
		}
	} else if op.content == "-" {
		val = float64(0)
		for _, operand := range operands{
			val -= operand.eval()
		}
	} else if op.content == "*" {
		v, _ := strconv.ParseFloat(operands[0].content, 64)
		val = v
		for i, operand := range operands{
			if i > 0{
				val *= operand.eval()
			}
		}
	} else if op.content == "/" {
		v, _ := strconv.ParseFloat(operands[0].content, 64)
		val = v
		for i, operand := range operands{
			if i > 0{
				val /= operand.eval()
			}
		}
	}
	return val
}

func make_ast(expr string) *Node{
	var root_node *Node = new(Node)
	var op *Node
	var operands []*Node
	var cursor int = 0

	for cursor<len(expr) {
		character := string(expr[cursor])
		if isin(append(NUMS,"."), character) {
			num :=  make_num(expr, cursor)
			children :=  []*Node{nil,nil,nil}
			operands = append(operands, &Node{parent: root_node, content: num, children:children})
			cursor += len(num)
		} else if "(" == character {
			// Incorrect
			// Try working through a situation
			// with nested brackets
			close_parenthesis := indexof(str2arr(expr), ")", cursor)
			//--------------------------------------------------------//
			//----------Breaks when close... is -1 -------------------//
			//--------------------------------------------------------//
			if close_parenthesis != -1 {
				operands = append(operands, make_ast(expr[cursor+1:close_parenthesis]))
				cursor = close_parenthesis+1
			} else {
				cursor++
			}
		} else if isin(OPS, character) {
			op_content := character
			op_children := []*Node {nil,nil,nil}
			op = &Node{parent: root_node, children: op_children, content: op_content}		
			cursor += len(op_content)
		} else {
			cursor++
		}	
	}
	*root_node = Node{parent:nil, children: append(operands, op), content: expr}
	fmt.Println("finished")
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

/*func indicesof[T comparable] (arr []T, el T) []int{
	var indices []int
	for i:=0; i < len(arr); i++ {
		if arr[i] == el {
			indices = append(indices, int(i))
		}
	}
	return indices
}
*/
