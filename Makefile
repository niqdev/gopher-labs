.DEFAULT_GOAL := help

require-%:
	@ if [ "$(shell command -v ${*} 2> /dev/null)" = "" ]; then \
		echo "[$*] not found"; \
		exit 1; \
	fi

check-param-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Missing parameter: [$*]"; \
		exit 1; \
	fi

# https://pablosanjose.com/the-best-makefile-help
help: ## Show this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##############################

local-up: require-docker ## Starts Localstack
	docker-compose -f local/docker-compose.yml up -d

local-logs: require-docker ## Tails Localstack logs
	docker-compose -f local/docker-compose.yml logs --follow

local-down: require-docker ## Stop Localstack and removes all contents
	docker-compose -f local/docker-compose.yml down -v
	rm -frv ./local/.localstack
