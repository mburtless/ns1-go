generate:
	swagger generate client -f swagger.json --template-dir templates -c ns1 -A ns1 -a ns1
clean:
	rm -rf ns1 models
openapi-to-swagger:
	api-spec-converter --from=openapi_3 --to=swagger_2 --syntax=json ./openapi.yaml > ./swagger.json
