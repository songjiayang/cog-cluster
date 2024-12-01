build-agent:
	docker build -f agent/Dockerfile -t cog-agent:0.1.0 .
build-api:
	docker build -f api/Dockerfile -t cog-api:0.1.0 .
start:
	docker-compose up -d 
stop:
	docker-compose down
test:
	curl http://localhost:8000/predictions -X POST \
		-H 'Content-Type: application/json' \
		-d '{"input": {"image": "https://gist.githubusercontent.com/bfirsh/3c2115692682ae260932a67d93fd94a8/raw/56b19f53f7643bb6c0b822c410c366c3a6244de2/mystery.jpg"}}'