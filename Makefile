help:
	@echo ""
	@echo "\033[0;34m    Velo Payments - Go Client (\033[1;34mvelopaymentsapi/velo-go\033[0;34m) \033[0m"
	@echo ""
	@echo "    To dynamically generate the Go client based on a spec issue the following command."
	@echo "    You can specify the spec by replacing the url parameter"
	@echo ""
	@echo "\033[92m    make WORKING_SPEC=https://raw.githubusercontent.com/velopaymentsapi/VeloOpenApi/master/spec/openapi.yaml client \033[0m"
	@echo ""

version:
	@docker run -i --rm mikefarah/yq sh -c "apk -q add curl && curl -s $$WORKING_SPEC -o /tmp/oa3.yaml;  yq r /tmp/oa3.yaml info.version" 2>&1

oa3config:
	sed -i.bak 's/"packageVersion": ".*"/"packageVersion": "${VERSION}"/g' oa3-config.json && rm oa3-config.json.bak

clean:
	rm -Rf api
	rm -Rf docs
	rm -f *.go
	rm -f README.md
	rm -Rf .openapi-generator
	rm -f .openapi-generator-ignore
	rm -f .travis.yml
	rm -f git_push.sh

generate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i $$WORKING_SPEC \
		-g go \
		-o /local \
		-c /local/oa3-config.json

trim:
	rm -Rf .openapi-generator
	rm .openapi-generator-ignore
	rm .travis.yml
	rm git_push.sh
	rm -Rf api

info:
	sed -i.bak '2s/.*/[![License](https:\/\/img.shields.io\/badge\/License-Apache%202.0-blue.svg)](https:\/\/opensource.org\/licenses\/Apache-2.0) [![npm version](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go.svg?style=svg)](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go)\\/' README.md && rm README.md.bak
	sed -i.bak 's/GIT_USER_ID\/GIT_REPO_ID/velopaymentsapi\/velo-go/' go.mod && rm go.mod.bak

build_client:
	#

client: clean generate trim info build_client

tests:
	# language: go

	# install:
	# - go get -d -v .

	# script:
	# - go build -v ./

commit:
	git add --all
	git commit -am 'bump version to ${VERSION}'
	git push --set-upstream origin master

build:
	@echo "Not doing anything just yet"

publish:
	git tag $(VERSION)
	git push origin tag $(VERSION)