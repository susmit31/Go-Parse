package main
import (
	"fmt"
//	"strconv"
)
func main(){
 var k [2]int32 = [2]int32{1,2}
// k = fmt.Sprintf("%T",[2]string{strconv.Itoa(2),strconv.Itoa(3)})
 arr2 := new([4]int32)
 *arr2 = make([]string,16)
 fmt.Println(k)
}
