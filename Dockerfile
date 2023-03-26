from alpine:3.15 as front-build
run apk add nodejs make npm
copy . .
run make frontend/out

from golang:1.19.3-alpine as appbuild
run apk add make
copy . /usr/src/app
workdir /usr/src/app
copy --from=front-build frontend/out frontend/out
run make build

from alpine:edge
copy --from=appbuild /usr/src/app/science-cup /bin/science-cup
expose 8080
cmd science-cup -nobrowser -b 0.0.0.0
