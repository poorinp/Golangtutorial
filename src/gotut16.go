package main
 
import ("fmt"
		"net/http"
		"html/template"
)

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, err := template.ParseFiles("basictemplating.html")
	// fmt.Println(err)
	fmt.Println(t.Execute(w,p))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Whoa, Go is neat</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}