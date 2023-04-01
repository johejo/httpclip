# httpclip

Paste to clipboard over HTTP

```
client -> server
```

## Install

server

```
go install github.com/johejo/httpclip/cmd/httpclipserver@latest
```

client (Optional, you can use curl, wget or any http client)

```
go install github.com/johejo/httpclip/cmd/httpclipclient@latest
```

## Usage

server

```
httpclipserver -addr $PORT
```

client

```
echo $CONTENT | httpclipclient -endpoint http://$SERVER_HOST:$PORT
```

curl client example

```
echo $CONTENT | curl -d @- http://$SERVER_HOST:$PORT
```

## License

MIT
