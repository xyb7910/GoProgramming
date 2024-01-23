package main

import (
	"LearingGo/crawier/engine"
	"LearingGo/crawier/persist"
	"LearingGo/crawier/scheduler"
	"LearingGo/crawier/zhenai/parser"
	"fmt"
	"regexp"
)

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
	}
	fmt.Printf("Match found: %d\n",
		len(matches))
}

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
