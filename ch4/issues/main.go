package main

import (
	"fmt"
	"log"
	"os"

	"self.gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s \n", item.Number, item.User.Login, item.Title)
	}
}

// 21 issues:
// #23331  mspiegel proposal: encoding/json: export the offset method of th
// #11046     kurin encoding/json: Decoder internally buffers full input
// #12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
// #5901        rsc encoding/json: allow override type marshaling
// #16212 josharian encoding/json: do all reflect work before decoding
// #7872  extempora encoding/json: Encoder internally buffers full output
// #17609 nathanjsw encoding/json: ambiguous fields are marshalled
// #22369     Splik encoding/json: add the full path to the field in Unmars
// #22752  buyology proposal: encoding/json: add access to the underlying d
