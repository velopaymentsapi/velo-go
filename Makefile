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
	# adjust README.md
	sed -i.bak '1s/.*/# Go client for Velo/' README.md && rm README.md.bak
	sed -i.bak '2s/.*/[![License](https:\/\/img.shields.io\/badge\/License-Apache%202.0-blue.svg)](https:\/\/opensource.org\/licenses\/Apache-2.0) [![npm version](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https:\/\/badge.fury.io\/go\/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go.svg?style=svg)](https:\/\/circleci.com\/gh\/velopaymentsapi\/velo-go)\\/' README.md && rm README.md.bak
	# adjust go.mod
	sed -i.bak 's/GIT_USER_ID\/GIT_REPO_ID/velopaymentsapi\/velo-go/' go.mod
	# remove duplicate definitions from api_payment_audit_service.go
	sed -i.bak '/GetFundingsV1Opts Optional/,/}/d' api_payment_audit_service.go
	sed -i.bak '/GetPaymentsForPayoutOpts Optional/,/}/d' api_payment_audit_service.go
	sed -i.bak '/GetFundingsV1Opts Optional/,/}/d' api_payment_audit_service.go
	sed -i.bak '/GetPaymentsForPayoutV4Opts Optional/,/}/d' api_payment_audit_service.go
	# rename duplicate definitions from model_language_2.go
	sed -i.bak 's/AR Language2/LANGUAGE_2_AR Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/EN Language2/LANGUAGE_2_EN Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/ES Language2/LANGUAGE_2_ES Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/FR Language2/LANGUAGE_2_FR Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/HE Language2/LANGUAGE_2_HE Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/HI Language2/LANGUAGE_2_HI Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/JA Language2/LANGUAGE_2_JA Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/PT Language2/LANGUAGE_2_PT Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/RU Language2/LANGUAGE_2_RU Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/UR Language2/LANGUAGE_2_UR Language2/' model_language_2.go && rm model_language_2.go.bak
	sed -i.bak 's/ZH Language2/LANGUAGE_2_ZH Language2/' model_language_2.go && rm model_language_2.go.bak
	# rename duplicate definitions from model_ofac_status.go
	sed -i.bak 's/PENDING OfacStatus/OFACSTATUS_PENDING OfacStatus/' model_ofac_status.go && rm model_ofac_status.go.bak
	sed -i.bak 's/PASSED OfacStatus/OFACSTATUS_PASSED OfacStatus/' model_ofac_status.go && rm model_ofac_status.go.bak
	sed -i.bak 's/FAILED OfacStatus/OFACSTATUS_FAILED OfacStatus/' model_ofac_status.go && rm model_ofac_status.go.bak
	# rename duplicate definitions from model_ofac_status_2.go
	sed -i.bak 's/PENDING OfacStatus2/OFACSTATUSV2_PENDING OfacStatus2/' model_ofac_status_2.go && rm model_ofac_status_2.go.bak
	sed -i.bak 's/PASSED OfacStatus2/OFACSTATUSV2_PASSED OfacStatus2/' model_ofac_status_2.go && rm model_ofac_status_2.go.bak
	sed -i.bak 's/FAILED OfacStatus2/OFACSTATUSV2_FAILED OfacStatus2/' model_ofac_status_2.go && rm model_ofac_status_2.go.bak
	# rename duplicate definitions from model_onboarded_status_2.go
	sed -i.bak 's/CREATED OnboardedStatus2/ONBOARDSTATUS2_CREATED OnboardedStatus2/' model_onboarded_status_2.go && rm model_onboarded_status_2.go.bak
	sed -i.bak 's/INVITED OnboardedStatus2/ONBOARDSTATUS2_INVITED OnboardedStatus2/' model_onboarded_status_2.go && rm model_onboarded_status_2.go.bak
	sed -i.bak 's/REGISTERED OnboardedStatus2/ONBOARDSTATUS2_REGISTERED OnboardedStatus2/' model_onboarded_status_2.go && rm model_onboarded_status_2.go.bak
	sed -i.bak 's/ONBOARDED OnboardedStatus2/ONBOARDSTATUS2_ONBOARDED OnboardedStatus2/' model_onboarded_status_2.go && rm model_onboarded_status_2.go.bak
	# rename duplicate definitions from model_payment_audit_currency_v4.go
	sed -i.bak 's/USD PaymentAuditCurrencyV4/PAYMENTAUDITCURRENCYV4_USD PaymentAuditCurrencyV4/' model_payment_audit_currency_v4.go && rm model_payment_audit_currency_v4.go.bak
	sed -i.bak 's/GBP PaymentAuditCurrencyV4/PAYMENTAUDITCURRENCYV4_GBP PaymentAuditCurrencyV4/' model_payment_audit_currency_v4.go && rm model_payment_audit_currency_v4.go.bak
	sed -i.bak 's/EUR PaymentAuditCurrencyV4/PAYMENTAUDITCURRENCYV4_EUR PaymentAuditCurrencyV4/' model_payment_audit_currency_v4.go && rm model_payment_audit_currency_v4.go.bak
	# remove duplicate definitions from model_payout_status_v3.go
	sed -i.bak 's/ACCEPTED PayoutStatusV3/PAYOUTSTATUSV3_ACCEPTED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/REJECTED PayoutStatusV3/PAYOUTSTATUSV3_REJECTED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/SUBMITTED PayoutStatusV3/PAYOUTSTATUSV3_SUBMITTED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/QUOTED PayoutStatusV3/PAYOUTSTATUSV3_QUOTED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/INSTRUCTED PayoutStatusV3/PAYOUTSTATUSV3_INSTRUCTED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/COMPLETED PayoutStatusV3/PAYOUTSTATUSV3_COMPLETED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/INCOMPLETE PayoutStatusV3/PAYOUTSTATUSV3_INCOMPLETE PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/CONFIRMED PayoutStatusV3/PAYOUTSTATUSV3_CONFIRMED PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	sed -i.bak 's/WITHDRAWN PayoutStatusV3/PAYOUTSTATUSV3_WITHDRAWN PayoutStatusV3/' model_payout_status_v3.go && rm model_payout_status_v3.go.bak
	# remove duplicate definitions from model_payout_status_v4.go
	sed -i.bak 's/ACCEPTED PayoutStatusV4/PAYOUTSTATUSV4_ACCEPTED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/REJECTED PayoutStatusV4/PAYOUTSTATUSV4_REJECTED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/SUBMITTED PayoutStatusV4/PAYOUTSTATUSV4_SUBMITTED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/QUOTED PayoutStatusV4/PAYOUTSTATUSV4_QUOTED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/INSTRUCTED PayoutStatusV4/PAYOUTSTATUSV4_INSTRUCTED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/COMPLETED PayoutStatusV4/PAYOUTSTATUSV4_COMPLETED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/INCOMPLETE PayoutStatusV4/PAYOUTSTATUSV4_INCOMPLETE PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/CONFIRMED PayoutStatusV4/PAYOUTSTATUSV4_CONFIRMED PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	sed -i.bak 's/WITHDRAWN PayoutStatusV4/PAYOUTSTATUSV4_WITHDRAWN PayoutStatusV4/' model_payout_status_v4.go && rm model_payout_status_v4.go.bak
	# rename duplicate definitions from model_watchlist_status.go
	sed -i.bak 's/NONE WatchlistStatus/WATCHLISTSTATUS_NONE WatchlistStatus/' model_watchlist_status.go && rm model_watchlist_status.go.bak
	sed -i.bak 's/PENDING WatchlistStatus/WATCHLISTSTATUS_PENDING WatchlistStatus/' model_watchlist_status.go && rm model_watchlist_status.go.bak
	sed -i.bak 's/REVIEW WatchlistStatus/WATCHLISTSTATUS_REVIEW WatchlistStatus/' model_watchlist_status.go && rm model_watchlist_status.go.bak
	sed -i.bak 's/PASSED WatchlistStatus/WATCHLISTSTATUS_PASSED WatchlistStatus/' model_watchlist_status.go && rm model_watchlist_status.go.bak
	sed -i.bak 's/FAILED WatchlistStatus/WATCHLISTSTATUS_FAILED WatchlistStatus/' model_watchlist_status.go && rm model_watchlist_status.go.bak
	# rename duplicate definitions from model_user_status.go
	sed -i.bak 's/ENABLED UserStatus/USERSTATUS_ENABLED UserStatus/' model_user_status.go && rm model_user_status.go.bak
	sed -i.bak 's/DISABLED UserStatus/USERSTATUS_DISABLED UserStatus/' model_user_status.go && rm model_user_status.go.bak
	sed -i.bak 's/PENDING UserStatus/USERSTATUS_PENDING UserStatus/' model_user_status.go && rm model_user_status.go.bak
	# rename duplicate definitions from model_user_type_2.go
	sed -i.bak 's/BACKOFFICE UserType2/USERTYPE2_BACKOFFICE UserType2/' model_user_type_2.go && rm model_user_type_2.go.bak
	sed -i.bak 's/PAYOR UserType2/USERTYPE2_PAYOR UserType2/' model_user_type_2.go && rm model_user_type_2.go.bak
	sed -i.bak 's/PAYEE UserType2/USERTYPE2_PAYEE UserType2/' model_user_type_2.go && rm model_user_type_2.go.bak

	- rm *.bak
	- rm ./docs/*.bak

build_client:
	#

client: clean generate trim info build_client

tests:
	go test -race $(go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt -args -key=${KEY} -secret=${SECRET} -payor=${PAYOR}

commit:
	git add --all
	git commit -am 'bump version to v${VERSION}'
	git push --set-upstream origin master

build:
	@echo "Not doing anything just yet"

publish:
	git tag v$(VERSION)
	git push origin tag v$(VERSION)