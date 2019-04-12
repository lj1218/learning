// Issues prints a table of GitHub issues matching the search terms.
package main

import (
    "fmt"
    "log"
    "os"
    "gopl.io/ch4/github"
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %.55s\n",
            item.Number, item.User.Login, item.Title)
    }
}
/* Output:
./issues repo:golang/go is:open json decoder
33 issues:
#29688   sheerun proposal: encoding/json: add InputOffset to json decode
#29686   sheerun json: Add InputOffset for stream byte offset access
#31309  thepudds encoding/json: add sample fuzz test for prototype of "f
#28923     mvdan encoding/json: speed up the decoding scanner
#30301     zelch encoding/xml: option to treat unknown fields as an erro
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#11046     kurin encoding/json: Decoder internally buffers full input
#30701 LouAdrien encoding/json: ignore tag "-" not working on embedded s
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#16212 josharian encoding/json: do all reflect work before decoding
#28143 Carpetsmo proposal: encoding/json: add "readonly" tag
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#5901        rsc encoding/json: allow override type marshaling
#14750 cyberphon encoding/json: parser ignores the case of member names
#22752  buyology proposal: encoding/json: add access to the underlying d
#28189     adnsv encoding/json: confusing errors when unmarshaling custo
#27179  lavalamp encoding/json: no way to preserve the order of map keys
#7213  davechene cmd/compile: escape analysis oddity
#7872  extempora encoding/json: Encoder internally buffers full output
#20528  jvshahid net/http: connection reuse does not work happily with n
#17609 nathanjsw encoding/json: ambiguous fields are marshalled
#19636 josselin- encoding/base64: decoding is slow
#20754       rsc encoding/xml: unmarshal only processes first XML elemen
#22816 ganelon13 encoding/json: include field name in unmarshal error me
#21092  trotha01 encoding/json: unmarshal into slice reuses element data
#28941     mvdan cmd/compile: teach prove about slice expressions
#15808 randall77 cmd/compile: loads/constants not lifted out of loop
#28952     mvdan cmd/compile: consider teaching prove about unexported i
#5819  gopherbot encoding/gob: encoder should ignore embedded structs wi
#20206 markdryan encoding/base64: encoding is slow
*/
