package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "htps://google.com",
		"amazon web services": "http://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "amazon web services")
	fmt.Println(websites)

}
