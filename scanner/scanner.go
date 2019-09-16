package scanner

import(
	"os"
	"bufio"
)
// Input is what the user types into the console
var Input string
//Scanner creates a Scanner which will read user input from the console
func Scanner()string{
	reader := bufio.NewReader(os.Stdin)
	Input, _ := reader.ReadString('\n')
	return Input
}