/* main package */
package main

/* importing necessary pkg */
import "net/http"

func main() {
	/* whatever is the most specific */
	http.HandleFunc("/", someFunc)
	/* nil is for default ServeMux */
	http.ListenAndServe(":8000", nil)
}

/* write my request back; incoming request */
func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe"))
}

/* RUN YOUR CODE FROM MAIN DIRECTORY!!! */