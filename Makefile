all:
	go build

install:
	cp tcp-to-tls /usr/local/bin

clean:
	rm tcp-to-tls

remove:
	rm /usr/local/bin/tcp-to-tls