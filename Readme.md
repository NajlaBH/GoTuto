
### GoLang scripts
>>> Some Go scripts from tutorials for tests and illustrations.


## - Requirement
### 1.Docker image: 
<code>docker pull najlabioinfo/gorevel:latest</code>


### 2.Quick access:
- Go help:
<code>docker run -it najlabioinfo/gorevel:latest sh -c 'go help'</code>

- Go Version:
<code>docker run -it najlabioinfo/gorevel:latest sh -c 'go version'</code>

- Go run example: Hello world
<code>docker run -it najlabioinfo/gorevel:latest sh -c 'go run /home/NajlaBH/GoTuto/Hello/helloworld.go'</code>


### 3.Run Go scripts
 - NB: 
<br> *  You have to copy GoTuto directory <code>cp -R /home/NajlaBH/GoTuto /go/src/ </code> to have all scripts available in this image under go path: 
<code>/go/src</code>
<br> *  You can create the script build with the following command:<br>
<code> go build scriptname.go </code>


## - References
### 1. <a href="https://golangbot.com/learn-golang-series">https://golangbot.com/learn-golang-series</a>.
### 2. Run your go script in the web: <a href="https://play.golang.org/">https://play.golang.org/</a>.



## - Scripts Lists
### 1. Hello
<code>go run /go/src/GoTuto/Hello/helloworld.go</code>

### 2. TimeLocation
<code>go run /go/src/GoTuto/TimeLoc/timelocation.go</code>



## - License
MIT
