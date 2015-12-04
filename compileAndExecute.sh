echo $1
go run main.go lex.go syntatic.go semantic.go $1
echo ''
echo ''
echo '[INFO] Initializing out.c compilation using GCC'
make run
