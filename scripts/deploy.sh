npm run build 
cat main.go | sed s/package\ main/package\ handler/ > main.go.tmp
mv main.go.tmp main.go
now
