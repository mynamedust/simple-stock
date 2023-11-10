up:
	@docker-compose -f ./devops/docker-compose.yml up --build

down:
	@docker-compose -f ./devops/docker-compose.yml down

clean:
	@docker-compose -f ./devops/docker-compose.yml down -v
	@docker container prune -f
	@docker volume prune -f

.PHONY: up down clean