generate-api:
	python ./../../spec/remove_extra_tags.py sandboxes
	openapi-python-client generate --output-path khulnasoft/api/api --overwrite --path ../../spec/openapi_generated.yml
	rm -rf khulnasoft/api/client
	mv khulnasoft/api/api/khulnasoft_api_client khulnasoft/api/client
	rm -rf khulnasoft/api/api

init:
	pip install openapi-python-client
