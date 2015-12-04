go run *.go $1
if [[ $? == 1 ]]
then
	echo '[INFO] exiting...'
else
	echo '[INFO] Initializing out.c compilation using GCC'
	make run
fi
