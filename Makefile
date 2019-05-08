SPEC_VERSION=0.0.1
SWAGGER_DIR=./swagger/
OPENAPI_PATH=$(SWAGGER_DIR)openapi-$(SPEC_VERSION).yaml
SWAGGER_PATH=$(SWAGGER_DIR)swagger-$(SPEC_VERSION).yaml

generate:
	swagger generate client -f $(SWAGGER_PATH) --template-dir templates -c ns1 -A ns1 -a ns1
clean:
	rm -rf ns1 models
openapi-to-swagger:
	api-spec-converter --from=openapi_3 --to=swagger_2 --syntax=json $(OPENAPI_PATH) > $(SWAGGER_PATH)
