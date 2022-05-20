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

    c := make(chan string) // channel declaration

    for _, link := range links{

     go checkLink(link,c) //initiaate new go routine 

    }
    fmt.Println(<- c)
    fmt.Println(<- c)
    fmt.Println(<- c)
    fmt.Println(<- c)
    fmt.Println(<- c)
    fmt.Println(<- c)
}

func checkLink(link string,  c chan string){
   _,err := http.Get(link) // BLOCKING CALL
   if err  != nil {
       fmt.Println(link,"might be down!")
       c <- "Might be down"
       return
   }
   fmt.Println(link,"is Ok!")
   c <- "is Ok!"
}