
build: Main

Main: src/Main.go
	go build -o Main.o src/Main.go

run: build
	./Main.o

genConfig: 
	go build -o GenConfig.o src/GenConfig.go src/Config.go
	./GenConfig.o

clean: 
	rm ./Main.o
	rm ./GenConfig.o

help: 
	@cat makeHelp.txt