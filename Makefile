
usage = "\
test"

.PHONY: usage graphql

# Must be the first target!
usage:
	@echo $(usage)


qs:
	cd question && go build -tags 'netcgo' && ./question -alsologtostderr=true

graphql:
	cd graphql && go build -tags 'netcgo' && ./graphql -alsologtostderr=true