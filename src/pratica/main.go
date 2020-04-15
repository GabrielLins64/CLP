package main
import(
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "VIADU")
		})
	
	http.LintenAndServe(":1337", nil)
}