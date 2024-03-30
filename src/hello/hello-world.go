package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	showIntroducion()

	for {
		showMenu()
		command := readCommandMenu()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Printf("Showing logs...")
		case 0:
			fmt.Printf("Exiting...")
			os.Exit(0)
		default:
			fmt.Printf("Invalid command. Plese try again.")
			os.Exit(-1)
		}
	}

	// if command == 1 {
	// 	fmt.Printf("Displaing...")
	// } else if command == 2 {
	// 	fmt.Printf("Showing logs...")
	// } else if command == 0 {
	// 	fmt.Printf("Exiting...")
	// } else {
	// 	fmt.Printf("Invalid command. Plese try again.")
	// }
}

func showIntroducion() {
	name := "Gabriel"
	version := 1.1

	fmt.Println("Welcome Sr. ", name)
	fmt.Println("Validate data. The actual version is", version)
	fmt.Println()
}

func showMenu() {
	fmt.Println("===== Menu =====")
	fmt.Println("1 -  Display monitoring ")
	fmt.Println("2 -  Show logs ")
	fmt.Println("0 -  Exit ")
}

func readCommandMenu() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}

func startMonitoring() {
	fmt.Printf("Displaing...\n")

	var sites []string
	sites = append(sites, "https://www.youtube.com")
	sites = append(sites, "https://www.google.com")
	sites = append(sites, "https://www.olhardigital.com.br")

	for i := 0; i < monitoring; i++ {
		for _, sites := range sites {
			fmt.Println("Testing site: ", sites)
			testingSites(sites)
		}
		fmt.Println("=========================================================\n")
		time.Sleep(delay * time.Second)
	}
}

func testingSites(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("The site:", site, "was load with success!\n")
	} else {
		fmt.Println("Site:", site, "contains error. Status code: ", resp.StatusCode, "\n")
	}
}
