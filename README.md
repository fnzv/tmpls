# TMPLS (trackmepls)

![](img/nsa.jpg?raw=true)

Yes.. not a L2 networking protocol (j.k.) but a very simple URL tracking tool in order to get notified in realtime when a particular URL is opened

NOTE: this tool was made for educational/studying purposes (must write this here before someone starts doing something un-"gdpr")

# Setup

This application is very simple and requires only Go, under this repo you will find .yml files and a Dockerfile in order to deploy it on a k8s cluster or run it locally (before deploying on k8s change the docker image and registry, on the example you will see the gitlab container registry)

# Quickstart
Run the tracking service locally:
```
# docker build -t trackmepls .

# docker run -d -p 127.0.0.1:80:80 trackmepls

# curl 127.1/?uri=example.org -I
HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: http://example.org
Date: Thu, 23 Jul 2020 23:15:39 GMT
```

In order to recieve Telegram notification you will require to setup those ENV vars and pass them to the docker container (`e.g. docker run -e TGBOT_CHATID -e TGBOT_TOKEN -d -p 127.0.0.1:80:80 trackmepls`):
- TGBOT_CHATID
- TGBOT_TOKEN

# Usage
1) Generate a decoy html page containing a 301 request to our tracking service, e.g. example.org/something_you_want_to_share (this is the url you will send to the enduser containing a hidden redirect):
```
<html>
<head>
<meta http-equiv="refresh" content="0; url=http://mytrackingservice.xyz/?uri=example.org/destination_url">
</head>
</html>
```
2) Once the tracking service is up and running you can test it by browsing the tracking URL:
```
http://mytrackingservice.xyz/?uri=example.org/destination_url
```

3) Everytime a visitor opens the tracking service the owner will get a realtime notification

User opens https://example.org/something_you_want_to_share 

Request is redirected to http://mytrackingservice.xyz/?uri=example.org/destination_url via html

The tracking service will log all the request headers and send a notification to the owner of the URL via Telegram

User is redirected to the intended url example.org/destination_url

Under the current working directory all the requests are logged also inside `tracker.log` file

For any doubt or question feel free to contact me