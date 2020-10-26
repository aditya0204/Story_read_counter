package controller

import (
	"fmt"
)

func(c *Controller)isUserExist(userName string)bool{
	q:= `SELECT count(*) from users where user_name=$1;`

	rows,err:=c.db.Query(q,userName)
	if err!=nil{
		fmt.Println("Error in query ",err)
	}else {
		for rows.Next(){
			var cnt int
			rows.Scan(&cnt)
			if cnt==0{
				rows.Close()
				return false
			}else{
				rows.Close()
				return true
			}
		}
	}
	rows.Close()
	return true

}

func(c *Controller)insertUserToDB(userName,password string)error{
	q:=`INSERT INTO users(user_name,password)VALUES($1,$2);`
	_,err:=c.db.Exec(q,userName,password)
	return err
}

func (c *Controller)fetchPassHashFromDB(userName string)string{
	q:= `SELECT password from users where user_name = $1;`
	rows ,err := c.db.Query(q,userName)

	if err!=nil{
		fmt.Println("Error in fetching hash")
		rows.Close()
		return ""
	}

	for rows.Next(){
		var hash string
		rows.Scan(&hash)
		rows.Close()
		return hash
	}

	rows.Close()
	return ""
}

func (c *Controller)getStories(username string)[]Story{
   var stories []Story

   q := `SELECT * from story ;`

   rows,err:=c.db.Query(q)

   if err!=nil{
   	fmt.Println("error in getting stories ",err)
   	rows.Close()
   }else{
   	 for rows.Next(){
   	 	var st Story
   	 	err:=rows.Scan(&st.StoryTitle,&st.StoryContent,&st.ImageLink,&st.Views)
   	 	if err!=nil{
   	 		fmt.Println("error in scan stories ",err)
		}else{
			st.Username=username
			stories=append(stories,st)
		}
	 }
   }
   rows.Close()
	return stories
}
func (c *Controller)getStory(username ,storyName string)Story{
	var st Story

	q := `SELECT * from story where story_title=$1 ;`

	rows,err:=c.db.Query(q,storyName)

	if err!=nil{
		fmt.Println("error in getting stories ",err)
		rows.Close()
	}else{
		for rows.Next(){
			err:=rows.Scan(&st.StoryTitle,&st.StoryContent,&st.ImageLink,&st.Views)
			if err!=nil{
				fmt.Println("error in scan stories ",err)
			}else{
				st.Username=username
			}
		}
	}
	rows.Close()
	return st
}

func (c *Controller)getViewsOfStory(storyTitle string)int{
	q:= `SELECT count(*) from story_views where story_title=$1;`

	rows,err:=c.db.Query(q,storyTitle)
	if err!=nil{
		fmt.Println("Error in query ",err)
	}else {
		for rows.Next(){
			var cnt int
			rows.Scan(&cnt)
			if cnt==0{
				rows.Close()
				return 0
			}else{
				rows.Close()
				return cnt
			}
		}
	}
	rows.Close()
	return 0
}