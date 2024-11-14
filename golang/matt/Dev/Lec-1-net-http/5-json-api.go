package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Define a struct to match the JSON data structure from the API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Handler function to fetch data and serve as HTML
func fetchPosts(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make HTTP GET request to JSONPlaceholder API
	resp, err := http.Get("https://jsonplaceholde.typicode.com/posts")
	if err != nil {
		// Step 2: If there's an error, respond with http.Error()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Step 3: Use json.NewDecoder to decode the response into a slice of Post structs
	var posts []Post
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&posts); err != nil {
		// Step 4: If JSON decoding fails, return an error
		http.Error(w, "Failed to decode JSON response", http.StatusInternalServerError)
		return
	}

	// Step 5: Create an HTML template to display the data
	tmpl, err := template.New("posts").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Posts</title>
		</head>
		<body>
			<h1>Posts from JSONPlaceholder API</h1>
			<ul>
				{{range .}}
					<li>
						<h2>{{.Title}}</h2>
						<p>{{.Body}}</p>
						<p><strong>User ID:</strong> {{.UserID}} | <strong>Post ID:</strong> {{.ID}}</p>
					</li>
				{{end}}
			</ul>
		</body>
		</html>
	`)
	if err != nil {
		// If template parsing fails, return an error
		http.Error(w, "Failed to parse HTML template", http.StatusInternalServerError)
		return
	}

	// Step 6: Serve the HTML response by executing the template with the posts data
	if err := tmpl.Execute(w, posts); err != nil {
		// If template execution fails, return an error
		http.Error(w, "Failed to generate HTML", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Step 7: Set up the server and routes
	http.HandleFunc("/", fetchPosts)

	// Step 8: Start the server on port 8080
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
