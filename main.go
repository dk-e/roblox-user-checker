package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var (
	available []string
	names     []string
	threadc   = 10
	mutex     sync.Mutex
)

func main() {
	file, err := os.Open("usernames.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Starting")

	var wg sync.WaitGroup
	for i := 0; i < threadc; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			validate(divide(names, threadc)[index])
		}(i)
	}
	wg.Wait()

	fmt.Println("Writing")
	writeToFile("available.txt", available)
	fmt.Println("Finished!")
}

func divide(stuff []string, parts int) [][]string {
	var divided [][]string
	for i := 0; i < parts; i++ {
		var part []string
		for j := i; j < len(stuff); j += parts {
			part = append(part, stuff[j])
		}
		divided = append(divided, part)
	}
	return divided
}

func validate(users []string) {
	for _, user := range users {
		resp, err := http.Get(fmt.Sprintf("https://auth.roblox.com/v1/usernames/validate?request.username=%s&request.birthday=1337-04-20", user))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			mutex.Lock()
			available = append(available, user)
			mutex.Unlock()
		}
	}
}

func writeToFile(filename string, data []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, noob := range data {
		_, err := writer.WriteString(noob + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writer.Flush()
}
