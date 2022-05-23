Go module = package management for go 
at the root folder : go mod init 

✅ TIL ✅
2022.05.24
<p> This came from the template : {{index .StringMap "test"}}</p>
.StringMap : From all the data in StringMap
"test" : I'm looking for this - "test" - particular data

🚨 Error 🚨

(1) import "github.com/ayleeee/go-course/pkg/render" is a program, not an importable package
-> It occured as I did not change the package name at the top

for example,

render.go 
I left package name of render.go as main 
so I changed it to 'render' and It worked.


(2) angelalee@Angelas-MacBook-Air hello-world % go run cmd/web/*.go
cmd/web/main.go:5:2: package hello-world/pkg/config is not in GOROOT (/opt/homebrew/Cellar/go/1.18.1/libexec/src/hello-world/pkg/config)
cmd/web/main.go:6:2: package hello-world/pkg/render is not in GOROOT (/opt/homebrew/Cellar/go/1.18.1/libexec/src/hello-world/pkg/render)
-> I added modules in main.go 
-> I think It occured because I didn't add right path

(3) import cycle not allowed 
-> you can not depend on each other
-> seperate datas