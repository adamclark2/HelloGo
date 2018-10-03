
build: Main

Main: src/Main.go
	go build -o Main.o src/Main.go

run: build
	./Main.o

clean: 
	rm ./Main.o