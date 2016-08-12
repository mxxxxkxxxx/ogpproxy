![build status](https://circleci.com/gh/mxxxxkxxxx/ogpproxy.svg?style=shield&circle-token=b6b5d7d9a137072aa00cd4ff757b8da488ef5673)

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
http://[YOUR HOST]/?url=[WEBSITE]
```

```
# ex1
$ curl http://localhost:8080/?url=https://mxxxxkxxxx.github.io/ogpproxy/samples/01.html
{  
   "errors":null,
   "ogp":{  
      "title":[  
         "ogp proxy sample website"
      ],
      "type":[  
         "website"
      ],
      "url":[  
         "http://mxxxxkxxxx.github.io/ogpproxy/samples/01.html"
      ],
      "image":[  
         {  
            "url":"http://mxxxxkxxxx.github.io/ogpproxy/samples/images/01.jpg",
            "secure_url":"",
            "type":"",
            "width":0,
            "height":0
         }
      ],
      "video":[  
         {  
            "url":"http://mxxxxkxxxx.github.io/ogpproxy/samples/videos/01.mp4",
            "secure_url":"",
            "type":"",
            "width":0,
            "height":0
         }
      ],
      "audio":[  
         {  
            "url":"http://mxxxxkxxxx.github.io/ogpproxy/samples/sounds/01.mp3",
            "secure_url":"",
            "type":""
         }
      ],
      "description":[  
         "This is a sample website for ogpproxy."
      ],
      "determiner":[  
         "auto"
      ],
      "locale":[  
         {  
            "locale":"en_US",
            "alternate":""
         }
      ],
      "site_name":[  
         "ogpproxy sample 01"
      ]
   }
}
```

```
# ex2
$ curl http://localhost:8080/?url=https://mxxxxkxxxx.github.io/ogpproxy/samples/02.html
{  
   "errors":null,
   "ogp":{  
      "title":[  
         "ogp proxy sample website"
      ],
      "type":[  
         "website"
      ],
      "url":[  
         "http://mxxxxkxxxx.github.io/ogpproxy/samples/01.html"
      ],
      "image":[  
         {  
            "url":"http://mxxxxkxxxx.github.io/ogpproxy/samples/images/01.jpg",
            "secure_url":"https://mxxxxkxxxx.github.io/ogpproxy/samples/images/01.jpg",
            "type":"image/jpeg",
            "width":400,
            "height":300
         }
      ],
      "video":[  
         {  
            "url":"http://mxxxxkxxxx.github.io/ogpproxy/samples/videos/01.mp4",
            "secure_url":"https://mxxxxkxxxx.github.io/ogpproxy/samples/videos/01.mp4",
            "type":"video/mp4",
            "width":400,
            "height":300
         }
      ],
      "audio":[  
         {  
            "url":"http://mxxxxkxxxx.github.io/ogpproxy/samples/sounds/01.mp3",
            "secure_url":"https://mxxxxkxxxx.github.io/ogpproxy/samples/sounds/01.mp3",
            "type":"audio/mpeg"
         }
      ],
      "description":[  
         "This is a sample website for ogpproxy."
      ],
      "determiner":[  
         "auto"
      ],
      "locale":[  
         {  
            "locale":"en_US",
            "alternate":"ja_JP"
         }
      ],
      "site_name":[  
         "ogpproxy sample 01"
      ]
   }
}
```
