pkill -f main.go &&
pkill -f go-build &&
nohup go run main.go >log.log 2>&1 &
