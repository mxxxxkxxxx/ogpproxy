# ogpproxy
This is a proxy to cache OGP (Open Graph Protocol) Data.

# Requirements

- glide

# Installation

```
git clone git@github.com:mxxxxkxxxx/ogpproxy.git
cd ogpproxy
glide install
go build ogpp

# launch
./ogpp
```

# Usage

```
http://[YOUR HOST]/?url=[WEBSITE]/

# ex.
$ curl http://127.0.0.1:8080/?url=http://www.google.com/
{"errors":null,"ogp":{"title":"","description":"g.co/fruit \ufffdŁA2016 Doodle \ufffdt\ufffd\ufffd\ufffd[\ufffdc\ufffdQ\ufffd[\ufffd\ufffd\ufffd\ufffd\ufffd`\ufffdF\ufffdb\ufffdN #GoogleDoodle","url":"","image":"http://www.google.com/logos/doodles/2016/2016-doodle-fruit-games-day-3-5741908677623808-thp.png","site_name":"","locale":""}}
```
