.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help section
	@echo ""
	@echo "\033[0;34m    Velo Payments - Go Client (\033[1;34mvelopaymentsapi/velo-go\033[0;34m) \033[0m"
	@echo ""
	@echo "    To dynamically generate the Go client based on a spec issue the following command."
	@echo "    You can specify the spec by replacing the url parameter"
	@echo ""
	@echo "\033[92m    make client WORKING_SPEC=https://raw.githubusercontent.com/velopaymentsapi/VeloOpenApi/master/spec/openapi.yaml \033[0m"
	@echo ""
	@echo "     *** Available Targets ***"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "    \033[36m%-38s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

version: ## Parse (via docker) spec file passed in WORKING_SPEC variable to print the version
	@docker run -i --rm mikefarah/yq:3 sh -c "apk -q add curl && curl -s $$WORKING_SPEC -o /tmp/oa3.yaml;  yq r /tmp/oa3.yaml info.version" 2>&1

oa3config: ## Set version on the openapi generator config to value of the VERSION variable
	sed -i.bak 's/"packageVersion": ".*"/"packageVersion": "${VERSION}"/g' oa3-config.json && rm oa3-config.json.bak

clean: ## Remove files & directories that are auto created by generator cli
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

generate: ## Run (via docker) generator cli against a spec file passed in WORKING_SPEC variable to create files
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i $$WORKING_SPEC \
		-g go \
		-o /local \
		--global-property=skipFormModel=false \
		-c /local/oa3-config.json

trim: ## Remove unused files that are auto geneated
	- rm -Rf .openapi-generator
	- rm .openapi-generator-ignore
	- rm .travis.yml
	- rm git_push.sh
	- rm -Rf api

info: ## Update the auto generated README.md with Velo info
	# create interface for instance of oneOf
	echo "package velopayments" >> model_oneOf.go
	echo "type OneOfPingPaymentStatusChangedPaymentRejectedOrReturnedOnboardingStatusChangedPayableStatusChangedPayeeDetailsChangedDebitStatusChanged interface {" >> model_oneOf.go
	echo "}" >> model_oneOf.go
	
	# adjust README.md
	sed -i.bak '1s/.*/# Go client for Velo/' README.md && rm README.md.bak
	sed -i.bak '2s/.*/[![License](https:\/\/img.shields.io\/badge\/License-Apache%202.0-blue.svg)](https:\/\/opensource.org\/licenses\/Apache-2.0) [![npm version](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go.svg?style=svg)](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go)\\/' README.md && rm README.md.bak
	
	# adjust go.mod
	sed -i.bak 's/GIT_USER_ID\/GIT_REPO_ID/velopaymentsapi\/velo-go/' go.mod
	
	# rename duplicate definitions from api_users.go
	sed -i.bak 's/ApiResendTokenRequest/ApiResendUserTokenRequest/g' api_users.go && rm api_users.go.bak
	
	- rm *.bak
	- rm ./docs/*.bak

build_client: ## Post generate client processing (optional per sdk)
	#

client: clean generate trim info build_client ## Generate sdk via cli

.PHONY: tests
tests: ## Run (via docker) tests using supplied variables KEY, SECRET, PAYOR, APIURL
	docker build -t=velo-sdk-go-tests .
	docker run -t -v $(PWD):/usr/src/app -e CGO_ENABLED=0 -e KEY=${KEY} -e SECRET=${SECRET} -e PAYOR=${PAYOR} -e APIURL=${APIURL} -e APITOKEN="" velo-sdk-go-tests go test -short $(go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt

commit: ## Commit & Push generated client to git repo
	sed -i.bak 's/- Package version: .*/- Package version: ${VERSION}/' README.md && rm README.md.bak
	git add --all
	git commit -am 'bump version to v${VERSION}'
	git push --set-upstream origin master

build: ## Build compiled package (optional per sdk)
	@echo "Not doing anything just yet"

publish: ## Tag & Push git ref, (optional per sdk) publish to 3rd party registry
	git tag v$(VERSION)
	git push origin tag v$(VERSION)