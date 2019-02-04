package controllers

import (
	"net/http"
	//"encoding/json"
	"fmt"
)

import "database/sql"
import "log"
import _"github.com/go-sql-driver/mysql"


func (app *Application) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	

	renderTemplate(w, r, "login.html",nil)
}


func Authenticate(username string,password string,typ string) string {
	db, err := sql.Open("mysql","root:saleel@tcp(127.0.0.1:3306)/supplychain")
    if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(username)
	defer db.Close()
	results, err := db.Query("SELECT * FROM Users where username=? and password=? and type = ?",username,password,typ)
	if err != nil {
		panic(err.Error())
	}
	if(results.Next()){
		return "VALID"
	}
	return "INVALID"

}

func (app *Application) LoginFunctionHandler(w http.ResponseWriter, r *http.Request){ 

  
  fmt.Println("Continue 2")
  name := r.PostFormValue("username")
  pwd := r.PostFormValue("password")
  typ := r.PostFormValue("type")
   redirectTarget := "/"
    
    a:=Authenticate(name,pwd,typ)
    fmt.Println(a)
    

    if(a=="VALID"){
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //w.WriteHeader(http.StatusAccepted)
    fmt.Println("Logged In")
    //setSession(name, w)
	if(typ=="manufacturer")	{
http.Redirect(w, r, "/home.html", 302)
		}else if(typ=="distributor") {
http.Redirect(w, r, "/request.html", 302)
		}else {
http.Redirect(w, r, "/history.html", 302)
		}
    
 //http.Redirect(w, r, redirectTarget, 302)
    }else{
	//Pop up saying didn't logged in successfully
    http.Redirect(w, r, "/", 302)
    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //w.WriteHeader(http.StatusUnauthorized)
    fmt.Println("Incorrect credentials")
    
    //http.Redirect(w, r, "/login.html", 302)
    }
   
 fmt.Println(redirectTarget)


}

