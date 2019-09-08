mkdir -p out

# css
mkdir -p out/static
mkdir -p out/static/css
for FILE in static/css/* 
do 
  minify "$FILE" > "out/$FILE"
done

# img
mkdir -p out/static/img
cp -R static/img/* out/static/img

# html
mkdir -p out/html
for FILE in html/*
do 
  minify "$FILE" > "out/$FILE"
done


