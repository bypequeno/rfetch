// rfetch -> https://github.com/bypequeno/rfetch
// Licensed under GNU 2.0
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func reset() {
	fmt.Println("\x1b[0m")
}

func rainbow(freq float64, i float64) []int {
	r := int(math.Sin(freq*i)*127 + 128)
	g := int(math.Sin(freq*i+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(freq*i+4*math.Pi/3)*127 + 128)
	return []int{r, g, b}
}

func main() {
	global_vars := []string{"USER", "HOSTNAME", "USER", "LOGNAME", "SHELL", "HOME"}
	prints := []string{}
	spaces := []string{}
	for _, variable := range global_vars {
		now_var := os.Getenv(variable)
		prints = append(prints, now_var)
		space := strings.Repeat(" ", 20-len([]rune(now_var)))
		spaces = append(spaces, space)
	}
	line := strings.Repeat("-", 18+len([]rune(spaces[1])))
	fetch := fmt.Sprintf(`
	"%s@%s"
	%s
	|user: %q %s       |
	|os: %q   %s       |         
	|shell: %q%s       |   
	|home: %q %s       |    
	%s
	`, prints[0], prints[1], line, prints[2], spaces[2], prints[3], spaces[3], prints[4], spaces[4], prints[5], spaces[5], line)
	for i := range fetch {
		rgb := rainbow(float64(0.1), float64(i))
		fmt.Printf("\x1b[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
		fmt.Print(string(fetch[i]))
	}
	reset()
}
