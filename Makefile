# executable for linux based systems
unix-cli:
	go build -o store main.go

# executable for Windows
win-cli:
	go build -o family-tree.exe main.go	