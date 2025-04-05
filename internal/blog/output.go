package blog

import "fmt"

func Output(completeItems []string, unCompleteItems []string) {
	fmt.Println("==========================")
	fmt.Print("Successfully fetched notes: ")
	for _, id := range completeItems {
		fmt.Printf("%s, ", id)
	}
	fmt.Println()
	fmt.Print("Failed to fetch notes: ")
	for _, id := range unCompleteItems {
		fmt.Printf("%s, ", id)
	}
	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("All done!")
}
