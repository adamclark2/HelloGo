package main

import "net/http"
import "html"
import "io/ioutil"
import "fmt"

var includes []byte;
var header string;

var topic = [...]string {"Computer Science", "Stacks", "Queues", "Cool Stuff"};

func main() {
    var err error = nil;
    includes, err = ioutil.ReadFile("src/scriptincludes.html");
    if(err != nil){
        panic(err)
    }

    var tmp []byte;
    tmp, err = ioutil.ReadFile("src/header.html");
    if(err != nil){
        panic(err)
    }
    header = string(tmp)

    http.HandleFunc("/topics/", handleTopic)
    http.HandleFunc("/", handleIndex)
    http.ListenAndServe(":8080", nil)
}

func getCard(s string) string{
    var tmp = "";
    tmp += "<a href=\"/topics/" + html.EscapeString(s) + "\"class=\"card\">"+
        "<div class=\"card-header text-dark\">" + 
            html.EscapeString(s) + 
        "</div>" + 
        "<div class=\"card-body text-dark\">" + 
            "woot" + 
        "</div>" + 
    "</a>"

    return tmp;
}

func handleTopic(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, string(includes) + header);
    fmt.Fprintf(w, "<div class=\"jumbotron container text-dark\">" + html.EscapeString(r.URL.Path) + "</div>")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, string(includes) + header);

    fmt.Fprintf(w, "<div class=\"pt-4 pb-4\"></div>");
    fmt.Fprintf(w, "<div class=container>");
    for i:=0; i < 4;i++ {
        fmt.Fprintf(w, getCard(topic[i]))
        fmt.Fprintf(w, "<div class=\"pt-4 pb-4\">");
    }

    fmt.Fprintf(w, "</div>");
}