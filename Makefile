run :
	@docker-compose -f compose.yml build
	@docker-compose -f compose.yml -p med up --build -d

stop :
	@docker-compose -f compose.yml -p med down --remove-orphans
