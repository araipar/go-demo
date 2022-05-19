package main
import ("fmt"
		"net/http"
		"os"
		"io"
	)
type logWriter struct{}

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

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just Wrote this many bytes : " , len(bs))
	return len(bs),nil
}






