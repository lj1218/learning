// gopl.io/ch5/wait
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    urls := []string{
        "https://www.baidu.com",
        "bad.gopl.io",
        "https://127.0.0.1",
    }
    for _, url := range urls {
        if err := WaitForServer(url); err != nil {
            log.Fatalf("Site is down: %v\n", err)
        }
    }
}

func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil  // success
        }
        sleeptime := time.Second << uint(tries)  // exponential back-off
        log.Printf("server not responding (%s); sleep %d seconds before retrying...",
            err, sleeptime / 1000000000)
        time.Sleep(sleeptime)  // exponential back-off
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
