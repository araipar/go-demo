package main

import ("fmt"
        "net/http"
        "time"        
)

func main() {

    links := []string{
        "http://google.com",
        "http://facebook.com",
        "http://stackoverflow.com",
        "http://twitter.com",
        "http://amazon.com",        
    }

    c := make(chan string) // channel declaration

    for _, link := range links{

     go checkLink(link, c) //initiaate new go routine 

    }
   
    for  l := range c{
       go func(link string){
        time.Sleep(5 * time.Second)
        checkLink(link,c)
       }(l)
    }

}

func checkLink(link string,  c chan string){
   _,err := http.Get(link) // BLOCKING CALL
   if err  != nil {
      fmt.Println(link,"might be down!")
       c <- link
       return
   }
   fmt.Println(link,"is Ok!")
   c <- link
}