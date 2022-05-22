Go module = package management for go 
at the root folder : go mod init 

🚨 Error 🚨

(1) import "github.com/ayleeee/go-course/pkg/render" is a program, not an importable package
-> It occured as I did not change the package name at the top

for example,

render.go 
I left package name of render.go as main 
so I changed it to 'render' and It worked.
