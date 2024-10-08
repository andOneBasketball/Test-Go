package main
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
func main() {
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url, "http://") || ! strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		fmt.Println("fetching: ", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Println("status code: ", resp.StatusCode)
		resp.Body.Close()
		if err != nil {
			fmt.Println("err data: ", err)
		}
	}
}