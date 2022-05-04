package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), "|")
	match, _ := regexp.MatchString(strs[0], strs[1])
	fmt.Println(match)
}
