# .PHONY: test
# test:
	# docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	# docker-compose -f docker-compose.test.yml down --volumes 

test: test-db-down test-db-up
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

test-db-up:
	docker-compose -f docker-compose.test.yml up -d --build postgres

test-db-down:
	docker-compose -f docker-compose.test.yml down --volumes
