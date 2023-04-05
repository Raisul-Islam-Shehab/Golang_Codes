package main

import "fmt"

type Post struct {
	ID      int
	Content string
	Author  string
}

var postById map[int]*Post
var postByAuthor map[string][]*Post

func store(post *Post) {
	postById[post.ID] = post
	postByAuthor[post.Author] = append(postByAuthor[post.Author], post)
}

func main() {
	postById = make(map[int]*Post)
	postByAuthor = make(map[string][]*Post)

	post1 := &Post{ID: 1, Content: "First Post", Author: "Jamil"}
	post2 := &Post{ID: 2, Content: "Second Post", Author: "Fahim"}
	post3 := &Post{ID: 3, Content: "Third Post", Author: "Fahim"}
	post4 := &Post{ID: 4, Content: "Fourth Post", Author: "Jamil"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	// fmt.Println(postById[1])
	// fmt.Println(postById[2])

	_, ok := postByAuthor["Jamil"]
	if ok {
		for _, post := range postByAuthor["Jamil"] {
			fmt.Println(post)
		}
	} else {
		fmt.Println("There is no author named 'Jamil'")
	}

	_, ok = postByAuthor["Shehab"]
	if ok {
		for _, post := range postByAuthor["Shehab"] {
			fmt.Println(post)
		}
	} else {
		fmt.Println("There is no author named 'Shehab'")
	}
}
