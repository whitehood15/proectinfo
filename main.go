package main

import (
	"fmt"
	"net/http"	
	"io/ioutil"
    "log"
)
func logic(word string, input string, used string, mistakes int){
        
}


func root(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("./layouts/login.html")
    if err==nil {
    	w.Header().Add("Content-Type","text/html")
    	w.Write(data)
    } else {
    	w.WriteHeader(404)
    	w.Write([]byte("404 Page not found - "+http.StatusText(404)))
    }
}
func gamepage(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("Spanzhash")
    if(hashExists(cookie.Name))
	data, err := ioutil.ReadFile("./layouts/gamepage.html")
    if err==nil {
    	w.Header().Add("Content-Type","text/html")
    	w.Write(data)
    } else {
    	w.WriteHeader(404)
    	w.Write([]byte("404 Page not found - "+http.StatusText(404)))
    }
}
func dashboard(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("Spanzhash")
    if(hashExists(cookie.Name))
    
	data, err := ioutil.ReadFile("./layouts/dashboard.html")
    if err==nil {
    	w.Header().Add("Content-Type","text/html")
    	w.Write(data)
    } else {
    	w.WriteHeader(404)
    	w.Write([]byte("404 Page not found - "+http.StatusText(404)))
    }
}
func register(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./layouts/register.html")
    if err==nil {
    	w.Header().Add("Content-Type","text/html")
    	w.Write(data)
    } else {
    	w.WriteHeader(404)
    	w.Write([]byte("404 Page not found - "+http.StatusText(404)))
    }
}
func login(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    password := r.URL.Query().Get("password")
    log.Println(email,password)
    if email=="whitehood15@gmail.com" {
        w.Write([]byte("id"))
    } else {
        w.Write([]byte("usrNA"))
    }
}
func game(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func insert(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func modify(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}



func main() {
	//serve pages
	http.HandleFunc("/",root);	
	http.HandleFunc("/gamepage",gamepage);
	http.HandleFunc("/dashboard",dashboard);
	http.HandleFunc("/register",register);
	//serve Resources	
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img")))) 
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css")))) 
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js")))) 
	//serve API
	http.HandleFunc("/login",login);
	http.HandleFunc("/game",game);
	http.HandleFunc("/insert",root);
	http.HandleFunc("/modify",root);		

	// start server
    http.ListenAndServe(":8080", nil)
}