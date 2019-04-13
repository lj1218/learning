// issueshtml.go
package main

import (
    "gopl.io/ch4/github"
    "html/template"
    "log"
    "os"
    "time"
)

const templ = `<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Age</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.CreatedAt | daysAgo}} days</td>
</tr>
{{end}}
</table>
`

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
// $ go build issueshtml.go
// $ ./issueshtml repo:golang/go commenter:gopherbot json decoder >issues.html
/* Output:
see issues.html
*/
// $ ./issueshtml repo:golang/go 3133 10535 >issues2.html
/* Output:
see issues2.html
*/
