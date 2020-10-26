package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "story"
)

func NewDbClient() *sql.DB {
		psqlInfo := os.Getenv("DATABASE_URL")

		if len(psqlInfo)==0{
			psqlInfo=fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)
		}

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		//defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")
	return db
}


func SeedDB(db *sql.DB){
	q:=`INSERT INTO story (story_title, story_content, story_image_link) VALUES
	('Dogs and the world','Contrary to popular belief, 
                    Lorem Ipsum is not simply random text. 
					It has roots in a piece of classical Latin literature from 45 BC,
                    making it over 2000 years old. Richard McClintock, a Latin professor 
                    at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin 
                    words, consectetur, from a Lorem Ipsum passage, and going through the cites of the 
                    word in classical literature, discovered the undoubtable source. Lorem Ipsum comes 
                    from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" 
                    (The Extremes of Good and Evil) by Cicero, written in 45 BC. 
                    This book is a treatise on the theory of ethics, very popular during the Renaissance.
                     The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", 
                      comes from a line in section 1.10.32.',
                   'https://picsum.photos/id/1012/3973/2639'
   ),
('Travel world','Contrary to popular belief, 
                    Lorem Ipsum is not simply random text. 
					It has roots in a piece of classical Latin literature from 45 BC,
                    making it over 2000 years old. Richard McClintock, a Latin professor 
                    at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin 
                    words, consectetur, from a Lorem Ipsum passage, and going through the cites of the 
                    word in classical literature, discovered the undoubtable source. Lorem Ipsum comes 
                    from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" 
                    (The Extremes of Good and Evil) by Cicero, written in 45 BC. 
                    This book is a treatise on the theory of ethics, very popular during the Renaissance.
                     The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", 
                      comes from a line in section 1.10.32.',
                   'https://picsum.photos/id/1015/6000/4000'
   ),
('Parks','Contrary to popular belief, 
                    Lorem Ipsum is not simply random text. 
					It has roots in a piece of classical Latin literature from 45 BC,
                    making it over 2000 years old. Richard McClintock, a Latin professor 
                    at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin 
                    words, consectetur, from a Lorem Ipsum passage, and going through the cites of the 
                    word in classical literature, discovered the undoubtable source. Lorem Ipsum comes 
                    from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" 
                    (The Extremes of Good and Evil) by Cicero, written in 45 BC. 
                    This book is a treatise on the theory of ethics, very popular during the Renaissance.
                     The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", 
                      comes from a line in section 1.10.32.',
                   'https://picsum.photos/id/1029/367/267'
   ),
('Trains','Contrary to popular belief, 
                    Lorem Ipsum is not simply random text. 
					It has roots in a piece of classical Latin literature from 45 BC,
                    making it over 2000 years old. Richard McClintock, a Latin professor 
                    at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin 
                    words, consectetur, from a Lorem Ipsum passage, and going through the cites of the 
                    word in classical literature, discovered the undoubtable source. Lorem Ipsum comes 
                    from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" 
                    (The Extremes of Good and Evil) by Cicero, written in 45 BC. 
                    This book is a treatise on the theory of ethics, very popular during the Renaissance.
                     The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", 
                      comes from a line in section 1.10.32.',
                   'https://picsum.photos/id/1026/367/267'
   )
	;`

	_,err:=db.Exec(q)
	if err!=nil{
		fmt.Println(err)
	}
}

func DbMigration(db *sql.DB){
	q := `create table users(user_name text,password text, primary key(user_name));`

	db.Exec(q)

	q= `create table story(story_title text ,story_content text ,story_image_link text ,total_views integer default 0, primary key(story_title));`

	db.Exec(q)

	q=`create table story_views(story_title text,user_name text,primary key(story_title,user_name));`

	db.Exec(q)
}