package main

import ("fmt"
        "net/http"        
)

func main() {

    links := []string{
        "http://google.com",
        "http://facebook.com",
        "http://stackoverflow.com",
        "http://twitter.com",
        "http://amazon.com",        
    }

    for _, link := range links{

     go checkLink(link) //initiaate new go routine 

    }


}

func checkLink(link string){
   _,err := http.Get(link) // BLOCKING CALL
   if err  != nil {
       fmt.Println(link,"might be down!")
       return
   }
   fmt.Println(link,"is Ok!")
}