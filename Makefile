.PHONY: fmt generate-models

fmt:
	@go fmt
generate-models:
	@oapi-codegen --config=config.yaml ${SWAGGER_PATH}