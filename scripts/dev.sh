cat main.go | sed s/package\ handler/package\ main/ > main.go.tmp
mv main.go.tmp main.go
npm start
