package main

import (
	"encoding/xml"
	"fmt"

	// "io"

	"io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	ID       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	ID      int    `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// decoder := xml.NewDecoder(xmlFile)
	// for {
	// 	t, err := decoder.Token()
	// 	if err == io.EOF {
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error decoding XML into tokens:", err)
	// 		return
	// 	}

	// 	switch se := t.(type) {
	// 	case xml.StartElement:
	// 		// if se.Name.Local == "comment" {
	// 		// 	var comment Comment
	// 		// 	decoder.DecodeElement(&comment, &se)
	// 		// 	fmt.Println(comment)
	// 		// }

	// 		if se.Name.Local == "post" {
	// 			var post Post
	// 			decoder.DecodeElement(&post, &se)
	// 			fmt.Println(post)
	// 		}
	// 	}
	// }

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
