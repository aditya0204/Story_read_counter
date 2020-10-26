package web

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"storyReadCounter/web/controller"
)

type WebServer struct {
   cont *controller.Controller
}



func NewWebServer(cont *controller.Controller)  *WebServer{
	return &WebServer{cont: cont}
}

var MapStoryVsViews = make(map[string]int)
var SockIdVsStory  = make(map[string]string)
func (ws *WebServer) InitRoutes(){

	r:=mux.NewRouter()
	server, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.On("connection", func(so socketio.Socket) {

		log.Println("on connection")
       var storyName string
        so.On("views", func(storyTitle string) {
        	fmt.Println("View received on ",storyTitle)
        	storyName = storyTitle
        	so.Join(storyName)
        	MapStoryVsViews[storyName]++
        	SockIdVsStory[so.Id()]=storyName
        	so.Emit("views",MapStoryVsViews[storyName])
        	server.BroadcastTo(storyTitle,"views",MapStoryVsViews[storyName])
		})

		so.On("disconnection", func() {
			storyName=SockIdVsStory[so.Id()]
			so.Leave(storyName)
			fmt.Println(storyName," storyname ",MapStoryVsViews[storyName])
			MapStoryVsViews[storyName]--

			server.BroadcastTo(storyName,"views",MapStoryVsViews[storyName])
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})


	r.HandleFunc("/signup",ws.cont.SignUp).Methods("POST")
	r.HandleFunc("/login",ws.cont.Login).Methods("POST")
	r.HandleFunc("/story",ws.cont.Story).Methods("GET")
	r.HandleFunc("/content",ws.cont.StoryContent).Methods("GET")
	r.Handle("/socket.io/", server)


	fs := http.FileServer(http.Dir("static"))
	r.Handle("/", fs)

   PORT := os.Getenv("PORT")

   if PORT==""{
   	 PORT="5000"
   }
   var ch = make(chan os.Signal)

   go  http.ListenAndServe(":"+PORT,r)

   fmt.Println("Listening on port :",PORT)

   <-ch


}


