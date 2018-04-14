.PHONY: clean

TARGET=howto-go-with-cpp

$(TARGET): ld2_build ld_build libfoo.a 
	echo ok

libfoo.a: foo.o cfoo.o
	ar r $@ $^
ld_build:
	g++ -dynamiclib -flat_namespace foo.cpp cfoo.cpp -o libfoo.so
gold_build:
	go build -buildmode=c-shared -o libgoo.so foo.go
%.o: %.cpp
	g++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
