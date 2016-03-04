gom:
	go get github.com/mattn/gom
	gom install

fmt:
	gom exec goimports -w *.go

test:
	gom test ./...

link:
	mkdir -p $(GOPATH)/src/github.com/hico-horiuchi
	ln -s $(CURDIR) $(GOPATH)/src/github.com/hico-horiuchi/go-github-streak

unlink:
	rm $(GOPATH)/src/github.com/hico-horiuchi/go-github-streak
	rmdir $(GOPATH)/src/github.com/hico-horiuchi
