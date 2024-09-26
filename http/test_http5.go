package main  
  
import (  
	"encoding/json"  
	"fmt"  
	"io/ioutil"  
	"log"  
	"net/http"  
	"os"  
)  
  
// GitHubAPIResponse 是一个通用的结构体，用于解码GitHub API的响应  
type GitHubAPIResponse struct {  
	TotalCount int    `json:"total_count"`  
	Items      []Item `json:"items"`  
}  
  
// Item 是Issues或Milestones的通用表示  
type Item struct {  
	Title  string `json:"title"`  
	Number int    `json:"number"`  
	// 可以根据需要添加更多字段  
}  
  
// fetchGitHubData 获取GitHub数据  
func fetchGitHubData(url string) ([]Item, error) {  
	client := &http.Client{}  
	req, err := http.NewRequest("GET", url, nil)  
	if err != nil {  
		return nil, err  
	}  
	// 设置你的GitHub Personal Access Token  
	req.Header.Set("Authorization", "token YOUR_GITHUB_ACCESS_TOKEN")  
  
	resp, err := client.Do(req)  
	if err != nil {  
		return nil, err  
	}  
	defer resp.Body.Close()  
  
	body, err := ioutil.ReadAll(resp.Body)  
	if err != nil {  
		return nil, err  
	}  
  
	var ghResp GitHubAPIResponse  
	err = json.Unmarshal(body, &ghResp)  
	if err != nil {  
		return nil, err  
	}  
  
	return ghResp.Items, nil  
}  
  
func mainHandler(w http.ResponseWriter, r *http.Request) {  
	// 示例URL，替换为你的GitHub仓库和路径  
	url := "https://api.github.com/repos/yourusername/yourrepo/issues"  
	items, err := fetchGitHubData(url)  
	if err != nil {  
		http.Error(w, err.Error(), http.StatusInternalServerError)  
		return  
	}  
  
	// 将数据以JSON格式返回  
	json.NewEncoder(w).Encode(items)  
}  
  
func main() {  
	http.HandleFunc("/", mainHandler)  
	port := os.Getenv("PORT")  
	if port == "" {  
		port = "8080"  
	}  
	log.Printf("Server starting on port %s\n", port)  
	if err := http.ListenAndServe(":"+port, nil); err != nil {  
		log.Fatal("ListenAndServe: ", err)  
	}  
}