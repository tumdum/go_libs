CFLAGS=-fPIC

all: inspect run

libstatic.a: static.o
	ar rc libstatic.a static.o

libdynamic.so: dynamic.o
	$(CC) -shared -o libdynamic.so dynamic.o

main_dynamic: libdynamic.so
	go build main_dynamic.go

main_static: libstatic.a
	go build main_static.go

inspect: main_dynamic main_static
	ldd main_dynamic main_static

run: main_dynamic main_static
	LD_LIBRARY_PATH=. ./main_dynamic
	./main_static

clean:
	rm -rf *.o *.a *.so
	rm -f main_dynamic
	rm -f main_static
