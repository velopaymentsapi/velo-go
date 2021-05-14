.PHONY: tests

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
	mkdir tests
	mv *_test.go tests
	rm -f *.go
	rm -f go.*.tmp
	mv tests/* ./
	rm -Rf tests
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
	# wget https://raw.githubusercontent.com/velopaymentsapi/changelog/main/README.md -O CHANGELOG.md

	echo "package velopayments" >> OneOfPingPaymentStatusChangedPaymentRejectedOrReturnedOnboardingStatusChangedPayableStatusChangedPayeeDetailsChangedDebitStatusChanged.go
	echo "type OneOfPingPaymentStatusChangedPaymentRejectedOrReturnedOnboardingStatusChangedPayableStatusChangedPayeeDetailsChangedDebitStatusChanged struct {" >> OneOfPingPaymentStatusChangedPaymentRejectedOrReturnedOnboardingStatusChangedPayableStatusChangedPayeeDetailsChangedDebitStatusChanged.go
    # echo "	Payload interface{}" >> OneOfPingPaymentStatusChangedPaymentRejectedOrReturnedOnboardingStatusChangedPayableStatusChangedPayeeDetailsChangedDebitStatusChanged.go
	echo "}" >> OneOfPingPaymentStatusChangedPaymentRejectedOrReturnedOnboardingStatusChangedPayableStatusChangedPayeeDetailsChangedDebitStatusChanged.go
	
	# adjust README.md
	sed -i.bak '1s/.*/# Go client for Velo/' README.md && rm README.md.bak
	sed -i.bak '2s/.*/[![License](https:\/\/img.shields.io\/badge\/License-Apache%202.0-blue.svg)](https:\/\/opensource.org\/licenses\/Apache-2.0) [![npm version](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go.svg?style=svg)](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go)\\/' README.md && rm README.md.bak
	# adjust go.mod
	sed -i.bak 's/GIT_USER_ID\/GIT_REPO_ID/velopaymentsapi\/velo-go/' go.mod
	# rename duplicate definitions from api_users.go
	sed -i.bak 's/ApiResendTokenRequest/ApiResendUserTokenRequest/g' api_users.go && rm api_users.go.bak
	
	- rm *.bak
	- rm ./docs/*.bak

build_client:
	#

client: clean generate trim info build_client

tests:
	# test and generate coverage
	# go test -race $(go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt -args -key=${KEY} -secret=${SECRET} -payor=${PAYOR}
	docker build -t=client-go-tests .
	docker run -t -v $(PWD):/usr/src/app -e CGO_ENABLED=0 -e KEY=${KEY} -e SECRET=${SECRET} -e PAYOR=${PAYOR} -e APIURL=${APIURL} -e APITOKEN="" client-go-tests go test -short $(go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt

commit:
	sed -i.bak 's/- Package version: .*/- Package version: ${VERSION}/' README.md && rm README.md.bak
	git add --all
	git commit -am 'bump version to v${VERSION}'
	git push --set-upstream origin master

build:
	@echo "Not doing anything just yet"

publish:
	git tag v$(VERSION)
	git push origin tag v$(VERSION)