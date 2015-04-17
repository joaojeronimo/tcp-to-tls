# tcp-to-tls

This is a tiny simple TCP to TLS proxy.

For every tcp connection it gets, it opens a TLS connection and pipes the input of the TCP connection to the TLS connection, and the output of the TLS connection to the TCP connection.

## tcp-to-tls -h

    $ ./tcp-to-tls -h
    Usage of ./tcp-to-tls:
      -host="": Remote host to connect. Mandatory.
      -port=0: Remote port to connect. Mandatory.
      -insecure-skip-verify=false: Skip verification of server's certificate chain and host name
      -interface="": Local interface to bind to. Defaults to every interface available.
      -local-port=8888: Remote port to bind to. Defaults to port-1.

## Example

In this example we'll use tcp-to-tls and connect to example.com on port 443 where this [example website](https://example.com) is served. Then we use curl to retrieve it from [http://localhost:9001](http://localhost:8888). Notice that we will get it using HTTP instead of HTTPS (because tcp-to-tls took care of wrapping the connection in TLS).

    $ tcp-to-tls -host="example.com" -port=443 &
    $ curl http://localhost:8888 -H "Host: example.com"

## Installing

    $ make
    $ make install

## Keeping it running

If want this to be running all the time and you are running Ubuntu, there is [an upstart service sample](https://github.com/joaojeronimo/tcp-to-tls/blob/master/upstart.conf) you can copy to `/etc/init/tcp-to-tls.conf` and then start it with `sudo service tcp-to-tls start`.

If you are running a systemd based Linux distribution, then copy the [systemd service sample](https://github.com/joaojeronimo/tcp-to-tls/blob/master/systemd.service) to `/lib/systemd/system/tcp-to-tls.service` and start it with `sudo systemctl start tcp-to-tls`.

Rember to change the arguments in both scripts to make them connect to your desired host and port.

Both scripts will have the service restart if it fails and start on startup (the upstart one even starts when there is an internet connection, I don't know how to do that with sytemd, if you do please send me a PR).

## Caveat

This very simple version does not include options to specify arbitrary certificates to accept from the server, so if you are using a self-signed certificate your only option is to run with `-insecure-skip-verify=true`.