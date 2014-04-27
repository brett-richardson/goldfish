Goldfish
========

The goldfish that remembers. 

_Goldfish is designed to be a long-term permanent logging solution._

Go HTTP server that remembers things (arbitrary blocks of text) in chronological order.
Currently uses Amazon S3 for persistant storage.


Usage
-----

1. Run the server.

```
go run main.go
```

2. POST to memories/:type/:id

For example, a POST to `memories/product/123`
with POST values:

```
POST http://localhost:3000/memories/product/123
   datetime=2014-01-01T21:14:00
   memory=Arbitrary block of text.
   
POST http://localhost:3000/memories/product/123
   datetime=2014-01-02T23:14:00
   memory=Another block of text.
```

3. Now a get request to `http://localhost:3000/memories/product/123`.

The response:

```
==== 2014-01-01T21:14:00 ====
Arbitrary block of text.

==== 2014-01-02T23:14:00 ====
Another block of text.
```
