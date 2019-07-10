
usage = "\
test"

.PHONY: usage graphql

# Must be the first target!
usage:
	@echo $(usage)


proto:
	cd protorepo && ./build.sh

qs:
	cd question && go build -tags 'netcgo' && ./question -alsologtostderr=true

ans:
	cd answer && go build -tags 'netcgo' && ./answer -alsologtostderr=true

graphql:
	cd graphql && go build -tags 'netcgo' && ./graphql -alsologtostderr=true

build: build-question build-answer build-graphql build-client

build-question:
	cd question && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

build-answer:
	cd answer && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

build-graphql:
	cd graphql &&	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

build-client:
	cd client && npx parcel build index.html

scp: scp-question scp-answer scp-graphql scp-client

scp-question:
	cd question && scp question ubuntu@118.24.201.128:/home/ubuntu/mianke

scp-answer:
	cd answer && scp answer ubuntu@118.24.201.128:/home/ubuntu/mianke

scp-graphql:
	cd graphql && scp graphql ubuntu@118.24.201.128:/home/ubuntu/mianke

scp-client:
	cd client && scp -r dist/* ubuntu@118.24.201.128:/home/ubuntu/mianke/client

sql:
	cd question/db && scp question.sql ubuntu@118.24.201.128:/home/ubuntu/mianke
	cd answer/db && scp answer.sql ubuntu@118.24.201.128:/home/ubuntu/mianke