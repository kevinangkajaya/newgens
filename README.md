This is assignment project for newgens interview test.

## Step by Step Guide
- Install golang from https://go.dev/doc/install.
- Install mysql from https://dev.mysql.com/downloads/installer/.
- Install any mysql GUI, ex: [MySQL Workbench](https://dev.mysql.com/downloads/workbench/) or [HeidiSQL](https://www.heidisql.com/download.php).
- Run SQL Script from folder `database/setup_mysql.sql` on this project on to restore MT202 database and table.
- Copy `config.yaml.example` into `config.yaml`. Change database parameters as needed.
- You have two ways to run the program, either by using the program directly or use golang command lines.

###### Use Program Directly
- Simply run main.exe from the root folder of this project. Run `./main.exe` on console.

###### Use Golang Command Lines
- From the root folder of this project, run `go run main.go` on console.

##### Running the Program
- Input the file location of MT202 file when asked (for example: C:\Users\User\Documents\MT202.txt).
- The program will show success message when it's completed or give error message when it has problems converting the file.
- Check if the database of MT202 has added the new data.

### To Do Code Testing
- On console, run `go test` or `go test -v`. The `-v` means verbose (showing more logs).
- To do specific test, use `go test -v -run {Test Name}`. The `{Test Name}` is the function name. The test function list:
    - TestGetDataMt202
    - TestInsertDataMt202
    - TestInsertDataMt202Raw
    - TestReadFilesMt202