package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/akhil/go-bookstore/pkg/routes"
)


type Movie Struct{
	ID string `json:"id"`
	Isbn string `json:"isbn`
	Title string `json:"title"`
	director *Director `json:"director"`

}
type Director Struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
var movies []Movie

func getmovies(w http.ResponseWriter,r *http.request){

	w.Header.Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)


}

func deletemovies(w http.ResponseWriter,r *http.request){

	w.Header.Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item := range params{
		if item.ID == params["id"]{
			movies=movies.append(movies[item],movies[item+1]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)


}
func getmovie(w http.ResponseWriter,r *http.request){

	w.Header.Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item := range params{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(movies)
			
			break
		}

	}
	


}
func createmovies(w http.ResponseWriter,r *http.request){
	w.Header.Set("Content-Type","application/json")
	var movie Movie
	json.NewEncoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000))
	movies=append(movies,movie)
	
	json.NewEncoder(w).Encode(movies)



}
func updatemovies((w http.ResponseWriter,r *http.request){

	w.Header.Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item := range params{
		if item.ID == params["id"]{
			movies=movies.append(movies[item],movies[item+1]...)
			var movie Movie
		json.NewEncoder(r.Body).Decode(&movie)
		movie.ID=params["id"]
		movies=append(movies,movie)
	
		json.NewEncoder(w).Encode(movies)
			
		}

	}
	json.NewEncoder(w).Encode(movies)

}
func main(){

	r:=mux.NewRouter()
	movies=append(movies,Movie{ID:"1",Isbn:"12345",Title:"vikram",&Director:{Firstname:"lk",Lastname:"u"}})
	movies=append(movies,Movie{ID:"2",Isbn:"12347",Title:"kaithi",&Director:{Firstname:"lc",Lastname:"u"}})
	
	r.HandleFunc("/getmovies",getmovies).Methods("GET")
	r.HandleFunc("/getmovie/{id}",getmovie).Methods("GET")
	r.HandleFunc("/createmovies",createmovies).Methods("POST")
	r.HandleFunc("/updatemovies",updatemovies).Methods("PUT")
	r.HandleFunc("/deletemovies",deletemovies).Methods("DELETE")

	fmt.Print("hi")
}