package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{ID: 1, Content: "First Post", Author: "Jamil"},
		{ID: 2, Content: "Second Post", Author: "Fahim"},
		{ID: 3, Content: "Third Post", Author: "Fahim"},
		{ID: 4, Content: "Fourth Post", Author: "Jamil"},
	}

	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()

	file, err := os.Open("post.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	for _, item := range record {
		id, err := strconv.ParseInt(item[0], 0, 0)
		if err != nil {
			log.Fatal(err)
		}
		post := Post{ID: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Printf("ID: %d, Content: \"%s\", Author: \"%s\"\n", post.ID, post.Content, post.Author)
	}
}
