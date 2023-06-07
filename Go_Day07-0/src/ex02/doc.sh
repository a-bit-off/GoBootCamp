go install -v golang.org/x/tools/cmd/godoc@latest
godoc -http=:6060
open http://localhost:6060/pkg/ex02/
godoc -url "http://localhost:6060/pkg/ex02/" > doc.html
