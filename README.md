### Golang HCL Test

#### Requires: 

* ![golang.org/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

* ![github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

* ![Golang installed] (This is required to run the go files)
### How To Run
MANUAL METHOD: 
1. First install golang

2. Have Mysql server installed and restore the sql file attached to database "hcl"
   You can restore the mysqldump.sql file by using the following command :
	mysql --host=localhost --user=root --database=hcl -e "source <PATH TO sqldump.sql>"
3. Install go dependencies
	$- go get golang.org/x/crypto/bcrypt
	$- go get github.com/go-sql-driver/mysql
4. Inside of **hcl_test.go** line **148** replace values with your own credentials

```go
db, err = sql.Open("mysql", "<root>:<password>@/<dbname>")
// Replace with 
db, err = sql.Open("mysql", "myUsername:myPassword@/hcl")
```


Simple Binary run(Preferred):
If you want to run the simple binary file named : hcl_test, then it assumes mysql has root user as login and no password with hcl database dump restored
I would really recommend you to use the simple binary run for this since it has all deps compiled within the binary. 
command: ./hcl_test
Note: the ./hcl_test will only work if you have mysql installed with root user without password and hcl db already restored.
	In order to restore the hcl db, execute the following command: 
	mysql --host=localhost --user=root --database=hcl -e "source <PATH TO sqldump.sql>"


I tried to implement login and signup to track per user tries but it took too much time for me to complete it and i think i made a mistake by over complicating things. 
Also, the counter for the number of tries was intended to be implemented using golang templates but falling short of time, i could not complete that too. 









