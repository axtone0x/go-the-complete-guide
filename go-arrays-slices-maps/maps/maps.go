package maps

import "fmt"

func main() {
	// websites := map[string]string{} //empty map {}
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["Linkedln"] = "https://linkedin.com"

	delete(websites,"Google")
	fmt.Println(websites)
}
