package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type MessageInfo struct {
	timestamp	time.Time
	sender		string
	message		string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s %s", os.Args[0], "chat.txt")
		os.Exit(1)
	}
	filename := os.Args[1]
	// warning: reading the entire file at once
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(d)
	r := regexp.MustCompile(`\[.*?\]`)
	res := r.FindAllStringIndex(data, -1)
	//fmt.Printf("Number of matches for the pattern: %d\n", len(res))
	layout := "1/2/06, 3:4:5 PM"
	var msgs []MessageInfo

	var start2 int
	for i:=0; i<len(res); i++ {
		start1, end1 := res[i][0], res[i][1]
		if i < len(res)-1 {
			//start2, _:= res[i+1][0], res[i+1][1]
			start2 = res[i+1][0]
		} else {
			start2 = len(data)
		}
		m := data[end1:start2]
		timestr := data[start1+1:end1-1]
		t, _ := time.Parse(layout, timestr)
		s := strings.Split(m, ":")
		if len(s) < 2 {
			// There are messages like -- "user1 added user2" etc -- this does not have any : symbol; excluded this kind of data
			//fmt.Printf(">>>> %s", m)
			continue
		}
		msg := MessageInfo{
			timestamp:	t,
			sender:		s[0],
			message:	m,
		}
		msgs = append(msgs, msg)
	}

	//fmt.Printf("# of messageInfo structs built: %d\n", len(msgs))

	analyzeMessages(msgs)
}


func analyzeMessages(msgs []MessageInfo) {
	cm := make(map[string]int)
	for _, m := range msgs {
		//fmt.Printf("%s\n", m.sender)
		cm[m.sender] += 1
	}

	// a string slice for storing sorted names
	var keys []string
	for k := range cm {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println()
	for _, k := range keys {
		fmt.Printf("%30s : %10d\n", k, cm[k])
	}
}

