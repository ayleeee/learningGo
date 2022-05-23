Go module = package management for go <br>
at the root folder : go mod init 

✅ TIL ✅ 

<b>2022.05.24</b>  
For example, "This came from the template : {{index .StringMap "test"}}"  
<b>.StringMap</b> : From all the data in StringMap <br>
<b>"test"</b> : I'm looking for this - "test" - particular data

🚨 Error 🚨

<b> (1) import "github.com/ayleeee/go-course/pkg/render" is a program, not an importable package </b> <br>
-> It occured as I did not change the package name at the top<br>

for example,

render.go <br>
I left package name of render.go as main <br>
so I changed it to 'render' and It worked. <br>


<b>(2) angelalee@Angelas-MacBook-Air hello-world % go run cmd/web/*.go </b><br>
cmd/web/main.go:5:2: package hello-world/pkg/config is not in GOROOT (/opt/homebrew/Cellar/go/1.18.1/libexec/src/hello-world/pkg/config)<br>
cmd/web/main.go:6:2: package hello-world/pkg/render is not in GOROOT (/opt/homebrew/Cellar/go/1.18.1/libexec/src/hello-world/pkg/render)<br>
-> I added modules in main.go <br>
-> I think It occured because I didn't add right path<br>

<b>(3) import cycle not allowed</b> <br>
-> you can not depend on each other<br>
-> seperate datas<br>
