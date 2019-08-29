dev: 
	cat main.go | sed 's/package handler/package main/g' > tmp && \
	rm main.go && \
	mv tmp main.go && \
	gin --all # go get github.com/codegangsta/gin

deploy: 
	cat main.go | sed 's/package main/package handler/g' > tmp && \
	rm main.go && \
	mv tmp main.go && \
	now # npm i -g now

test: 
	go test ./...

