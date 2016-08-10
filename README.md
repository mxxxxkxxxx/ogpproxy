# OGP Proxy
This is a proxy to cache OGP (Open Graph Protocol) Data.

# Requirements

- glide

# Installation

```
git clone git@github.com:mxxxxkxxxx/ogpproxy.git
cd ogpproxy
glide install
go build -o ogpp

# run as development (configure by config.development.json)
./ogpp

# run as production (configure by config.production.json)
ENV=production ./ogpp
```

# Usage

```
http://[YOUR HOST]/?url=[WEBSITE]/

# ex.
$ curl http://127.0.0.1:8080/?url=http://www.businessinsider.com/
{"errors":null,"ogp":{"title":"Business Insider","type":"blog","url":"http://www.businessinsider.com/","image":"http://static5.businessinsider.com/assets/images/us/logos/og-image-logo-social.png","audio":"","description":"The latest news from Business Insider","determiner":"","locale":"","locale_alternate":"","site_name":"Business Insider","video":""}}
```
