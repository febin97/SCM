package controllers

import (
	"net/http"
	//"encoding/json"
	"fmt"
	"github.com/gorilla/securecookie"
)

import "database/sql"
import "log"
import _"github.com/go-sql-driver/mysql"





func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}


var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))



func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}








func (app *Application) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	

	renderTemplate(w, r, "login.html",nil)
}


func Authenticate(username string,password string,typ string) string {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/supplychain")
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
    //setSession(typ, w)
	//fmt.Println(value["typ"])
	if(typ=="manufacturer")	{
http.Redirect(w, r, "/home1.html", 302)
		}else if(typ=="distributor") {
http.Redirect(w, r, "/home2.html", 302)
		}else {
http.Redirect(w, r, "/home3.html", 302)
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

func Register(username string,emailid string,password string,cntno string,typ string) string{
str:=""
	db, err := sql.Open("mysql", "root:saleel@tcp(127.0.0.1:3306)/supplychain")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    insert, err := db.Query("INSERT INTO Users VALUES (?,?,?,?,?)",username,emailid,password,cntno,typ)
    
    if err != nil {
        panic(err.Error())
        str ="NOT ADDED"
    }else{
     str = "ADDED"
}
    

defer insert.Close()
 return str   
}

func (app *Application) SignUpHandler(w http.ResponseWriter,r *http.Request){

name := r.PostFormValue("username")
  emailid := r.PostFormValue("emailid")
  pwd1 := r.PostFormValue("password1")

  pwd2 := r.PostFormValue("password2")
  cntno:=r.PostFormValue("contactno")
  typ :=  r.PostFormValue("type")
    //fmt.Fprintf(w, "Hello, %s! .. you are going to register %s --- %s", name,pwd1,pwd2)
//msg := "2 password don't match"
   if pwd1==pwd2{
     a := Register(name,emailid,pwd1,cntno,typ)
   
    fmt.Println(a)
    if(a=="ADDED"){
	setSession(name, w)
	fmt.Println("ADDED SUCCESSSFULLY")		
    http.Redirect(w, r, "/login.html", 302)
}else{
fmt.Println("FAILURE DURING SIGNING UP")
http.Redirect(w, r, "/login.html", 302) 
}
}else{
	fmt.Println("2 PASSWORDS DONT MATCH")
http.Redirect(w, r, "/login.html", 302) 
}


}


func (app *Application) LogOutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/home.html", 302)
}

