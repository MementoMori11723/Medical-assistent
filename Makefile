run :
	@docker-compose -f config/compose.yml -p med up --build -d

stop :
	@docker-compose -f config/compose.yml -p med down --remove-orphans
