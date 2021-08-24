ccy=\033[1;33m
cclb=\033[1;34m
ccnc=\033[0m

.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo "usage: make [target] ..."
	@echo ""
	@echo "targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} \
		| sort | awk 'BEGIN {FS = ":.*?## "}; \
		{printf "${cclb}%-30s${ccnc} %s\n", $$1, $$2}'

.PHONY: run
run: ## Run api
	@echo "${ccy}running API...${ccnc}"
	@DATABASE_URL="host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" go run cmd/api/main.go
	@exit 0
