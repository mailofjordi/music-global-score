#To develop on this repo:

0. You need go installed on your nachine (dockerfile is WIP)
1. git clone git@github.com:mailofjordi/music-global-score.git
2. cd music-global-score
3. docker run --rm -v $(pwd):/go/src/github.com/mailofjordi/music-global-score -w /go/src/github.com/mailofjordi/music-global-score instrumentisto/dep ensure -vendor-only
4. cd internal/app
5. go build
6. ./app
7. You can http request GET localhost:1323/score/pepa (you can do it from your browser)
