.PHONY: build build-sdk build-example run-example clean

CC=gcc
CFLAGS=-I.
LDFLAGS=-L. -lark

build: build-sdk build-example

build-sdk: main.go
	./build.sh

build-example: example/alice_to_bob_covenantless.c libark.so arksdk.h
	$(CC) -o example/alice_to_bob_covenantless example/alice_to_bob_covenantless.c $(CFLAGS) $(LDFLAGS)

run-example:
	./example/alice_to_bob_covenantless

clean:
	rm -f libark.so libark.h example/alice_to_bob_covenantless
