// issuesreport.go
package main

import (
    "gopl.io/ch4/github"
    "log"
    "os"
    "text/template"
    "time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ))

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
        return
    }
    if err := report.Execute(os.Stdout, result); err != nil {
        log.Fatal(err)
    }
}
// $ go build issuesreport.go
// $ ./issuesreport repo:golang/go is:open json decoder
/* Output:
33 issues:
---------------------------------------
Number: 29688
User:   sheerun
Title:  proposal: encoding/json: add InputOffset to json decoder
Age:    91 days
---------------------------------------
Number: 29686
User:   sheerun
Title:  json: Add InputOffset for stream byte offset access
Age:    91 days
---------------------------------------
Number: 31309
User:   thepudds
Title:  encoding/json: add sample fuzz test for prototype of "fuzzing as
Age:    6 days
---------------------------------------
Number: 28923
User:   mvdan
Title:  encoding/json: speed up the decoding scanner
Age:    141 days
---------------------------------------
Number: 30301
User:   zelch
Title:  encoding/xml: option to treat unknown fields as an error
Age:    53 days
---------------------------------------
Number: 11046
User:   kurin
Title:  encoding/json: Decoder internally buffers full input
Age:    1409 days
---------------------------------------
Number: 29035
User:   jaswdr
Title:  proposal: encoding/json: add error var to compare  the returned
Age:    133 days
---------------------------------------
Number: 30701
User:   LouAdrien
Title:  encoding/json: ignore tag "-" not working on embedded sub struct
Age:    34 days
---------------------------------------
Number: 12001
User:   lukescott
Title:  encoding/json: Marshaler/Unmarshaler not stream friendly
Age:    1348 days
---------------------------------------
Number: 16212
User:   josharian
Title:  encoding/json: do all reflect work before decoding
Age:    1017 days
---------------------------------------
Number: 28143
User:   Carpetsmoker
Title:  proposal: encoding/json: add "readonly" tag
Age:    184 days
---------------------------------------
Number: 26946
User:   deuill
Title:  encoding/json: clarify what happens when unmarshaling into a non
Age:    243 days
---------------------------------------
Number: 5901
User:   rsc
Title:  encoding/json: allow override type marshaling
Age:    2095 days
---------------------------------------
Number: 14750
User:   cyberphone
Title:  encoding/json: parser ignores the case of member names
Age:    1128 days
---------------------------------------
Number: 22752
User:   buyology
Title:  proposal: encoding/json: add access to the underlying data causi
Age:    513 days
---------------------------------------
Number: 28189
User:   adnsv
Title:  encoding/json: confusing errors when unmarshaling custom types
Age:    181 days
---------------------------------------
Number: 27179
User:   lavalamp
Title:  encoding/json: no way to preserve the order of map keys
Age:    232 days
---------------------------------------
Number: 7213
User:   davecheney
Title:  cmd/compile: escape analysis oddity
Age:    1902 days
---------------------------------------
Number: 7872
User:   extemporalgenome
Title:  encoding/json: Encoder internally buffers full output
Age:    1812 days
---------------------------------------
Number: 20528
User:   jvshahid
Title:  net/http: connection reuse does not work happily with normal use
Age:    682 days
---------------------------------------
Number: 17609
User:   nathanjsweet
Title:  encoding/json: ambiguous fields are marshalled
Age:    898 days
---------------------------------------
Number: 19636
User:   josselin-c
Title:  encoding/base64: decoding is slow
Age:    752 days
---------------------------------------
Number: 20754
User:   rsc
Title:  encoding/xml: unmarshal only processes first XML element
Age:    659 days
---------------------------------------
Number: 22816
User:   ganelon13
Title:  encoding/json: include field name in unmarshal error messages wh
Age:    508 days
---------------------------------------
Number: 21092
User:   trotha01
Title:  encoding/json: unmarshal into slice reuses element data between
Age:    632 days
---------------------------------------
Number: 28941
User:   mvdan
Title:  cmd/compile: teach prove about slice expressions
Age:    138 days
---------------------------------------
Number: 15808
User:   randall77
Title:  cmd/compile: loads/constants not lifted out of loop
Age:    1054 days
---------------------------------------
Number: 28952
User:   mvdan
Title:  cmd/compile: consider teaching prove about unexported integer fi
Age:    138 days
---------------------------------------
Number: 5819
User:   gopherbot
Title:  encoding/gob: encoder should ignore embedded structs with no exp
Age:    2112 days
---------------------------------------
Number: 20206
User:   markdryan
Title:  encoding/base64: encoding is slow
Age:    711 days
*/

