integration-test:
	sudo docker compose -f docker-compose-kafka.yml up --detach
	sleep 15
	sudo docker compose -f docker-compose-localstack.yml up --detach
	sleep 15
	aws --endpoint-url=http://localhost:4566 sqs create-queue --region=eu-west-1 --queue-name test-q --output table | cat
	go test -v ./...
	sudo docker compose -f docker-compose-localstack.yml down
	sudo docker compose -f docker-compose-kafka.yml down

setup:
	sudo docker compose -f docker-compose-kafka.yml up --detach
	sleep 15
	sudo docker compose -f docker-compose-localstack.yml up --detach
	sleep 15
	aws --endpoint-url=http://localhost:4566 sqs create-queue --region=eu-west-1 --queue-name test-q --output table | cat

teardown:
	sudo docker compose -f docker-compose-localstack.yml down
	sudo docker compose -f docker-compose-kafka.yml down