package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"storyReadCounter/pkg"
	"strings"
	"time"
)

type Controller struct {
	db *sql.DB
}

func NewController(db *sql.DB)*Controller{
	return &Controller{db: db}
}


func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    userName := strings.ToLower(r.Form.Get("user_name"))
    password:= r.Form.Get("pwd")
    calcPassHash,err:= pkg.GetPasswordHash(password)
    if err!=nil{
    	fmt.Fprintf(w,"Error in login ")
		return
	}
	fetchedHash := c.fetchPassHashFromDB(userName)

	fmt.Println("calulated hash is ",calcPassHash)
	fmt.Println("fetched hash is ",fetchedHash)

	if fetchedHash==calcPassHash{
		token:= pkg.GetUniqueToken()
		fmt.Println(userName+":"+password)
		expiration := time.Now().Add(365 * 24 * time.Hour)
		hc :=&http.Cookie{
			Name:       userName,
			Value:      token,
			Expires:    expiration,
		}

		http.SetCookie(w,hc)
		cookie, _ := r.Cookie(userName)
		fmt.Println(cookie.Name," << cookie >> ",cookie.Value)
		fmt.Println("User successfully signed up ",userName)
		http.Redirect(w,r,"/story?user_name="+userName,http.StatusFound)
	}else{
		fmt.Fprintf(w,"User does not exist")
	}



}
func  (c *Controller)  SignUp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := strings.ToLower(r.Form.Get("user_name"))
	password:= r.Form.Get("pwd")

	if c.isUserExist(userName){
		fmt.Fprintf(w,"User already exist")
	}else{
		passhash,err := pkg.GetPasswordHash(password)
		if err!=nil{
			fmt.Fprintf(w,"Error in backend")
		}else{
			err:=c.insertUserToDB(userName,passhash)
			if err!=nil{
				fmt.Println("Error in insertion to db ",err)
				fmt.Fprintf(w,"Error in backend")
			}else{
				fmt.Println("User successfully signed up ",userName)

				token:= pkg.GetUniqueToken()
				fmt.Println(userName+":"+password)
				expiration := time.Now().Add(365 * 24 * time.Hour)
				hc :=&http.Cookie{
					Name:       userName,
					Value:      token,
					Expires:    expiration,
				}

				http.SetCookie(w,hc)
				_,err := r.Cookie(userName)
				if err!=nil{
					fmt.Println("error in signup cookie ",err)
				}
				w.Header().Set("Content-Type", "text/html")

				w.Write([]byte(`<html>
	                            <head>
	                             <meta http-equiv="refresh" content="2;url=/" />
                             	<title>Page Moved</title>
								<body>Signed up successfully Login to continue</body>
	                            </head>`))
			    }
		}

	}
}


func (c *Controller)Story(w http.ResponseWriter, r *http.Request){
     q :=r.URL.Query()
     userName := q.Get("user_name")
     fmt.Println("User name ",userName)

     cookie,err := r.Cookie(userName)

     if err!=nil{
     	fmt.Println("error in story cookie ",err)
     	fmt.Fprintf(w,"session expired")
		 return
	 }
     fmt.Println("cookie found on story page ",cookie.Value)
	 if len(cookie.Value)>0{
	 	 fmt.Println("Inside tpl part ")
		 templ :=template.New("stories.html")
		 templ,err = templ.ParseFiles("./static/stories.html")

		 if err!=nil{
			 fmt.Println("error in template ",err)
			 fmt.Fprintf(w,"Error in backend")
		 }else{
		 	data := c.getStories(userName)
		 	p := Stories{Story: data}
			fmt.Println("feched data is ",data)
		 	fmt.Println("stories fetched are ",len(data))
			err:= templ.Execute(w,p)
			if err!=nil{
				fmt.Println("error tpl execute ",err)
				//fmt.Fprintf(w,"error in backend")
			}
		 }
	 }




}

func (c *Controller) StoryContent(w http.ResponseWriter, r *http.Request) {
	q :=r.URL.Query()
	userName := q.Get("user_name")
	storyTitle := q.Get("story_title")
	fmt.Println("User name ",userName ," story title ",storyTitle)
	cookie,err := r.Cookie(userName)

	if err!=nil{
		fmt.Println("error in story cookie ",err)
		fmt.Fprintf(w,"session expired")
		return
	}
	fmt.Println("cookie found on story page ",cookie.Value)
	if len(cookie.Value)>0{
		fmt.Println("Inside tpl part ")
		templ :=template.New("story.html")
		templ,err = templ.ParseFiles("./static/story.html")

		if err!=nil{
			fmt.Println("error in template ",err)
			fmt.Fprintf(w,"Error in backend")
		}else{
			data := c.getStory(userName,storyTitle)
			if data!=(Story{}){
				q:= `INSERT INTO story_views(story_title,user_name)VALUES($1,$2);`
				c.db.Exec(q,storyTitle,userName)
			}
			data.Views=c.getViewsOfStory(data.StoryTitle)
			err:= templ.Execute(w,data)
			if err!=nil{
				fmt.Println("error tpl execute ",err)
			}

		}
	}
}

