package main
import ("fmt"
		"net/http"
		"os"
		"io"
	)
func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	// bs := make([]byte,99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//3 line up there can be replaced by 1 line down here

	io.Copy(os.Stdout, resp.Body)

}






