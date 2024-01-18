package main

import (
	"LearingGo/advanced/retriever/mock"
	"LearingGo/advanced/retriever/real"
	"fmt"
	"time"
)

const URL = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get(URL)
}

func post(poster Poster) {
	poster.Post(URL,
		map[string]string{
			"name":   "ypb",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(URL, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(URL)
}

// 使用switch
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)

	fmt.Println("> Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever

	mockRetriever := mock.Retriever{
		Contents: "this is a fake web"}
	r = &mockRetriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// 使用 assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}
