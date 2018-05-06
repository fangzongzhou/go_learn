package basecommand
import (
	"bufio"
	"fmt"
	"os"

)

func dup1(str []string)  {
	counts:=make(map[string] int)
	files:=str
	if len((files))==0{

	}else {
		for _,arg :=range files{
			f,err:= os.Open(arg)
			if err!=nil{
				fmt.Printf("dup1: %v\n", err)
				continue
			}
			countLinex(f, counts)
			f.Close()
		}
	}

	for line,n :=range counts{
		if n>1{
			fmt.Printf("%d\t%s",n,line)
		}
	}
	
}

func countLinex(file *os.File ,counts map[string]int)  {
	input:=bufio.NewScanner(file)
	for input.Scan(){
		counts[input.Text()]++
	}
	
}
func main() {

	dup1(os.Args[1:])
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	var i=0
	for input.Scan() {
		if(i> 10) {
			break
		}
		counts[input.Text()]++
		i++

	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		} }

}