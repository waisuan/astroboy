dev: export APP_ENV = dev
dev:
	go run cmd/web/main.go

test-e2e: export APP_ENV = test
test-e2e:
	sudo docker compose -f docker-compose-storage.yml up --detach --wait
	go run cmd/web/main.go &
	sleep 5
	go test -v -count=1 ./e2e || true;
	pkill -f -e main
	sudo docker compose -f docker-compose-storage.yml down

setup: export APP_ENV = test
setup:
	sudo docker compose -f docker-compose-storage.yml up --detach --wait
	go run cmd/web/main.go &

teardown:
	pkill -f -e main
	sudo docker compose -f docker-compose-storage.yml down

#component-tests:
#	sudo docker compose -f docker-compose-kafka.yml up --detach --wait
#	sudo docker compose -f docker-compose-localstack.yml up --detach --wait
#	aws --endpoint-url=http://localhost:4566 sqs create-queue --region=eu-west-1 --queue-name test-q --output table | cat
#	sudo docker compose -f docker-compose-storage.yml up --detach --wait
#	go run cmd/datapipeline/main.go & echo $$! > datapipeline.PID;
#	go run cmd/web/main.go & echo $$! > web.PID;
#	sleep 5
#	go test -v ./...
#	if [ -e datapipeline.PID ]; then \
#		kill -TERM $$(cat datapipeline.PID) || true; \
#	fi;
#	if [ -e web.PID ]; then \
#		kill -TERM $$(cat web.PID) || true; \
#	fi;
#	sudo docker compose -f docker-compose-storage.yml down
#	sudo docker compose -f docker-compose-localstack.yml down
#	sudo docker compose -f docker-compose-kafka.yml down
#
#setup:
#	sudo docker compose -f docker-compose-kafka.yml up --detach --wait
#	sudo docker compose -f docker-compose-localstack.yml up --detach --wait
#	aws --endpoint-url=http://localhost:4566 sqs create-queue --region=eu-west-1 --queue-name test-q --output table | cat
#	sudo docker compose -f docker-compose-storage.yml up --detach --wait
#
#teardown:
#	sudo docker compose -f docker-compose-storage.yml down
#	sudo docker compose -f docker-compose-localstack.yml down
#	sudo docker compose -f docker-compose-kafka.yml down
#	if [ -e datapipeline.PID ]; then \
#		kill -TERM $$(cat datapipeline.PID) || true; \
#	fi;
#	if [ -e web.PID ]; then \
#		kill -TERM $$(cat web.PID) || true; \
#	fi;

dummy:
	go run cmd/jobrunner/main.go & echo $$! > dummy.PID;
	sleep 900
	if [ -e dummy.PID ]; then \
		kill -TERM $$(cat dummy.PID) || true; \
	fi;