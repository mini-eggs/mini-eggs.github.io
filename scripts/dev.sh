export SPACE_ID=$(cat now.json | jq -r '.env.SPACE_ID')
export ACCESS_TOKEN=$(cat now.json | jq -r '.env.ACCESS_TOKEN')
export HOME_ID=$(cat now.json | jq -r '.env.HOME_ID')
export SIDEBAR_ID=$(cat now.json | jq -r '.env.SIDEBAR_ID')

cat main.go | sed s/package\ handler/package\ main/ > main.go.tmp
mv main.go.tmp main.go
npm start
