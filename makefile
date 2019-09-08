dev: 
	cat main.go | sed 's/package handler/package main/g' > tmp && \
	rm main.go && \
	mv tmp main.go && \
	gin --all 

deploy: pre-deploy now-deploy post-deploy

pre-deploy: 
	cat main.go | sed 's/package main/package handler/g' > tmp && \
	rm main.go && \
	mv tmp main.go && \
	./build.sh && \
	mkdir tmp && \
	mv html tmp && \
	mv static tmp && \
	mv out/* .

now-deploy: 
	now

post-deploy: 
	rm -rf html static && \
	mv tmp/* . && \
	rm -rf tmp out

yolo: deploy
	now alias evanjon.es && \
	now alias www.evanjon.es

test: 
	go test ./...

setup:
	npm i -g now && \
	go get github.com/codegangsta/gin && \
	go get github.com/tdewolff/minify/cmd/minify

