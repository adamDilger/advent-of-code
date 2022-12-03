cd $1
echo $1.go | entr -c go run $1.go
