install:
	go build -o calculator ./src/*.go; mv calculator ./build

run:
	./build/calculator

clean:
	rm ./build/calculator

