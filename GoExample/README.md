# node-go-comparison

## Quick Summary

This folder holds the source code that supports the Go Example

Please visit https://ivocosme.github.io/node-go-concurrency/ for more information on this post.

## Requirements

* Node - 8.4.0
* npm - 5.4.2
* Go - 1.9
* curl - 7.54.0

## Setting up everything

### Setup assumptions

You already have Go installed (see above the versions I've used).

* [Go download/installation guide](https://golang.org/doc/install)

### Download the project

Open a terminal window inside and run:

```bash
git clone https://github.com/ivocosme/node-go-comparison.git
```

If you enter the new folder, you should see two folders with the following structure

```bash
.
├── GoExample
│   └── main.go
└── NodeExample
    ├── main.js
    └── package.json
```

### Setting up the project

For the sake of simplicity, we will:

* Use [concurrently](https://www.npmjs.com/package/concurrently), a NodeJS library that allows us to run two commands in paralel;
* cURL, so that we can access the API using the terminal.

To install concurrently, just run from the terminal

```bash
npm install -g concurrently
```

For cURl, you should have some version installed... Try to access it by opening the terminal and running

```bash
curl --version
```

You should see an output similar to this

```bash
curl 7.54.0 (x86_64-apple-darwin16.0) libcurl/7.54.0 SecureTransport zlib/1.2.8
Protocols: dict file ftp ftps gopher http https imap imaps ldap ldaps pop3 pop3s rtsp smb smbs smtp smtps telnet tftp
Features: AsynchDNS IPv6 Largefile GSS-API Kerberos SPNEGO NTLM NTLM_WB SSL libz UnixSockets
```

You're now good to go!

## Up and Running

### Running the go server example

```bash
cd GoExample
go run main.go
```

## Executing examples

### Go example

Open a terminal window and run

```bash
concurrently -r "curl http://localhost:8080/" "curl http://localhost:8080/" "curl http://localhost:8080/" "curl http://localhost:8080/" "curl http://localhost:8080/"
```

#### Expected result

All messages will appear randomly (because they have a random timer between 1 and 5 seconds) before the first 5 seconds. (Just like the non-blocking example using timeout).