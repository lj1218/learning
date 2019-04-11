package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
    edges := graph[from]
    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }
    edges[to] = true
}

func hasEdge(from, to string) bool {
    return graph[from][to]
}

func main() {
    addEdge("a", "1")
    addEdge("b", "2")
    addEdge("c", "3")
    addEdge("c", "4")
    fmt.Println(graph)
    fmt.Println(hasEdge("c", "3"))  // true
    fmt.Println(hasEdge("c", "4"))  // true
    fmt.Println(hasEdge("c", "5"))  // false
    fmt.Println(hasEdge("d", "1"))  // false
}
/* Output:
map[a:map[1:true] b:map[2:true] c:map[3:true 4:true]]
true
true
false
false
*/
