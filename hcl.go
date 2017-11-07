package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "strconv"
import "fmt"
import "golang.org/x/crypto/bcrypt"
//import "html/template"
import "net/http"

var db *sql.DB
var err error

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		res.Write([]byte("User created!"))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	res.Write([]byte("Hello" + databaseUsername))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
var templateStr=`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Enter the Secret Number</title 
</head>
<body>
    <h1>Input Number</h1>
    <form method="POST" action="/input">
        <input type="text" name="inputNumber" placeholder="inputNumber">
        <input type="submit" value="Submit">
    </form>
    <h2>Attempts : {{count}}</h1>
</body>
</html>`

func inputPage(res http.ResponseWriter, req *http.Request) {
        if req.Method != "POST" {
//	        var count string
//	        err = db.QueryRow("SELECT count(*) FROM attempts;").Scan(&count)
//		tmpl, _ := template.New("test").Parse(templateStr)
//		tmpl.Execute(res, &count)
//		return
                http.ServeFile(res, req, "input.html")
                return
        }

        number := req.FormValue("inputNumber")
	fmt.Println(number)
        getint,err:=strconv.Atoi(number)
	if err!=nil{
	    res.Header().Set("Content-Type", "text/html; charset=utf-8")
            fmt.Println(err)
	    //res.Write([]byte("Input not a number"))
	}
    	// check range of number
    	if getint <0 || getint >100{
    	    res.Header().Set("Content-Type", "text/html; charset=utf-8")
            res.Write([]byte("Input Number not within limits"))
    	}
        var databaseNumber string
        err = db.QueryRow("SELECT value FROM number WHERE allowed=1;").Scan(&databaseNumber)

        if err != nil {
	    res.Header().Set("Content-Type", "text/html; charset=utf-8")
            res.Write([]byte("Error."))
            return
        }

        if number == databaseNumber{ 
             res.Header().Set("Content-Type", "text/html; charset=utf-8")
             res.Write([]byte("Success. You have entered the correct number"))
        }else{
             res.Header().Set("Content-Type", "text/html; charset=utf-8")
             res.Write([]byte("Oops. You have entered the incorrect number"))
	}
        db.QueryRow("INSERT INTO attempts values(default,?)",number)


//res.Write([]byte("Hello" + databaseUsername))

}

func main() {
	// NOTE: Replace the below with the mysql DB credentials
//	db, err = sql.Open("mysql", "<root>:<password>@/<dbname>")
      db, err = sql.Open("mysql", "root:@/hcl")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", homePage)
        http.HandleFunc("/input",inputPage)
	http.ListenAndServe(":8080", nil)
}
