all:
	go build

install:
	cp tcp-to-tls /usr/local/bin

clean:
	rm tcp-to-tls