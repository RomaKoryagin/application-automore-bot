docker_migrate:
	docker pull migrate/migrate
	docker run -v $(PWD)/migrations:/var/www migrate/migrate:latest create -ext sql -dir /var/www migration
