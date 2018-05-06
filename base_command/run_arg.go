package basecommand

import (
	"fmt"
"os"
	"strings"
)


func main() {
	var res,split string
	for i:=1;i<len(os.Args);i++{
		res+=split+os.Args[i]
		split=" "
	}
	fmt.Println(res)

	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Println(os.Args[0:])
}
