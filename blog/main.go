package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
        "golang.org/x/crypto/bcrypt"
	"fmt"
)
type user struct {
	Username string
	Password string
}
var tpl *template.Template
var dbusers=map[string]user{}
var dbsessions=map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	//check if logged in
	c,err:=r.Cookie("session")
	if err == http.ErrNoCookie {
		tpl.ExecuteTemplate(w,"index.html",nil)
		return
	}

	if _,ok:=dbsessions[c.Value];ok{
		http.Redirect(w,r,"/secret",http.StatusSeeOther)
		return
	}


}

func register(w http.ResponseWriter, r *http.Request)  {
	var u user
	if r.Method==http.MethodPost {
		un:=r.FormValue("username")
		p:=r.FormValue("password")
		if _,ok:=dbusers[un];ok{
			tpl.ExecuteTemplate(w,"register.html","username is already registered")
			return
		}
		sid:=uuid.NewV4().String()
		http.SetCookie(w, &http.Cookie{
			Name:"session",
			Value: sid,
		})

		pb,err:=bcrypt.GenerateFromPassword([]byte(p),10)
		check(err)
		u=user{un,string(pb)}
		dbusers[un]=u
		dbsessions[sid]=un
		http.Redirect(w,r,"/secret",http.StatusSeeOther)
		return
	}

	c,err:=r.Cookie("session")
	if err ==http.ErrNoCookie {
		tpl.ExecuteTemplate(w,"register.html", nil)
		return
	}
	if _,ok:=dbsessions[c.Value];ok{
		http.Redirect(w,r,"/secret",http.StatusSeeOther)
		return
	}


}

func secret(w http.ResponseWriter, r *http.Request)  {
	c,err:=r.Cookie("session")
	if err!=nil{
		http.Redirect(w,r, "/",http.StatusForbidden)
		return
	}
	if un,ok1:=dbsessions[c.Value];ok1 {
		if u,ok2:=dbusers[un];ok2 {
			tpl.ExecuteTemplate(w,"secret.html", u)
		}
	}


}

func login(w http.ResponseWriter, r *http.Request)  {
	e:="wrong username/password, retry"
	if r.Method==http.MethodPost {
		un:=r.FormValue("username")
		p:=r.FormValue("password")
		psb:=[]byte(p)
		if u,ok1:=dbusers[un];ok1 {
			err:=bcrypt.CompareHashAndPassword([]byte(u.Password), psb)
			if err!=nil {
				tpl.ExecuteTemplate(w,"login.html", e)
				return
			}
			sid:=uuid.NewV4().String()
			http.SetCookie(w, &http.Cookie{
				Name:"session",
				Value: sid,
			})
			dbsessions[sid]=un
			http.Redirect(w,r,"/secret",http.StatusSeeOther)
			return
		}
		tpl.ExecuteTemplate(w,"login.html", e)
		return
	}


	c,err:=r.Cookie("session")
	if err ==http.ErrNoCookie {
		tpl.ExecuteTemplate(w,"login.html", nil)
		return
	}
	if un,ok1:=dbsessions[c.Value];ok1{
		if u,ok2:=dbusers[un];ok2 {
			tpl.ExecuteTemplate(w,"secret.html", u)
			return
		}
	}

}

func logout(w http.ResponseWriter, r *http.Request)  {
	c,err:=r.Cookie("session")
	if err==http.ErrNoCookie {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	c.MaxAge =-1
	http.SetCookie(w,c)
	http.Redirect(w,r,"/",http.StatusSeeOther)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}