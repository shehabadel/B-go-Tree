package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1.3.1

	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// 	if input.Text() == "done" {
	// 		break
	// 	}
	// }

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d \t %s \n", n, line)
	// 	}
	// }

	//1.3.2

	// counts := make(map[string]int)
	// files := os.Args[1:]
	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 		f.Close()
	// 	}
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d \t %s \n", n, line)
	// 	}
	// }

	//1.3.3

	// counts := make(map[string]int)
	// for _, filename := range os.Args[1:] {
	// 	data, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "filename: %s - %v \n", filename, err)
	// 		continue
	// 	}
	// 	for _, line := range strings.Split(string(data), "\n") {
	// 		counts[line]++
	// 	}
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d \t %s \n", n, line)
	// 	}
	// }
	// 1.3.4

	files := os.Args[1:]
	counts := make(map[string]int)
	inMap := make(map[string][]string)
	if len(files) == 0 {
		countLines(os.Stdin, counts, inMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts, inMap)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d \t %s  found in: %v \n ", n, line, inMap[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, inMap map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if len(inMap[input.Text()]) == 0 {
			inMap[input.Text()] = []string{f.Name()}
		} else {
			if found(f.Name(), inMap[input.Text()]) {
				continue
			}
			inMap[input.Text()] = append(inMap[input.Text()], f.Name())
		}
	}
}

func found(text string, fileNames []string) bool {
	for _, file := range fileNames {
		if text == file {
			return true
		}
	}
	return false
}
