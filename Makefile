generate:
	swagger generate client -f swagger.json --template-dir templates -c ns1 -A ns1 -a ns1
clean:
	rm -rf ns1/client ns1/models
