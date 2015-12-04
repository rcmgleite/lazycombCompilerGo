echo $1
go run main.go lex.go syntatic.go semantic.go $1
make all
