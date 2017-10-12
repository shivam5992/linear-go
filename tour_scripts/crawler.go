package main 
// https://jdanger.com/build-a-web-crawler-in-go.html
import (
	"fmt"
	"flag"
	"os"
	"crypto/tls"
	"net/url" 
	"net/http"    
  	// "io/ioutil" 
  	"io"
	"strconv"
	"strings"
	"golang.org/x/net/html"
	// "net/html"
)

var visited = make(map[string]bool)  


func collectlinks(httpBody io.Reader) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}


func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}


func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Enter Start Page")
		os.Exit(1)
	}

	 queue := make(chan string)   

	 go func() {            
                         
    queue <- args[0]     
  }()

  for uri := range queue {     
    enqueue(uri, queue)  
  }

	fmt.Println("Hello World")
	// retrieve(args[0])
}

func enqueue(uri string, queue chan string) {
  fmt.Println("fetching", uri)
  visited[uri] = true   
  tlsConfig := &tls.Config{
    InsecureSkipVerify: true,
  }
  transport := &http.Transport{
    TLSClientConfig: tlsConfig,
  }
  client := http.Client{Transport: transport}
  resp, err := client.Get(uri)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  links := collectlinks(resp.Body)

   for _, link := range links {
    absolute := fixUrl(link, uri)      
                                       
    if uri != "" {                     
                                       
      if !visited[absolute] {          
        go func() { queue <- absolute }()
      }
    }
  }
}



func fixUrl(href, base string) (string) {  
  uri, err := url.Parse(href)              
  if err != nil {                          
    return ""                              
  }                                        
  baseUrl, err := url.Parse(base)          
  if err != nil {                          
    return ""
  }
  uri = baseUrl.ResolveReference(uri)
  return uri.String()                      
}      