package web

import (
	"fmt"
	"github.com/chainHero/heroes-service/web/controllers"
	"net/http"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/login.html", app.LoginPageHandler)
	http.HandleFunc("/internal.html",app.LoginFunctionHandler)
	http.HandleFunc("/home1.html", app.HomeHandler1)
	http.HandleFunc("/home2.html", app.HomeHandler2)
	http.HandleFunc("/home3.html", app.HomeHandler3)
	http.HandleFunc("/request1.html", app.RequestHandler1)
	http.HandleFunc("/request2.html", app.RequestHandler2)
	http.HandleFunc("/request2keng.html",app.RequestHandlerKeng)
    	http.HandleFunc("/history1.html", app.HistoryHandler1)
	http.HandleFunc("/history2.html", app.HistoryHandler2)
	http.HandleFunc("/history3.html", app.HistoryHandler3)
	http.HandleFunc("/register.html", app.SignUpHandler)
	http.HandleFunc("/logout.html", app.LogOutHandler)
	
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//http.Redirect(w, r, "/login.html", http.StatusTemporaryRedirect)
		//renderTemplate(w, r, "home.html")
	//})
	
	http.HandleFunc("/",app.LoginPageHandler)
	fmt.Println("Listening (http://localhost:4000/) ...")
	http.ListenAndServe(":4000", nil)
}
