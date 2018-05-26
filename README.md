# This is an experiment with go lang for testing asynchronous providers calls

## To develop on this repo:

1. Just for your first time 😜
    1. git clone git@github.com:mailofjordi/music-global-score.git
    2. cd music-global-score
    3. docker run --rm -v $(pwd):/go/src/github.com/mailofjordi/music-global-score -w /go/src/github.com/mailofjordi/music-global-score instrumentisto/dep ensure -vendor-only
4. docker-compose up --build
5. You can http request GET localhost:1323/score/pepa (you can do it from your browser)
6. ctrl + C to stop the docker container