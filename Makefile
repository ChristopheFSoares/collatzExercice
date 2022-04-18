default: build

build:
	docker-compose build

run: build
	docker-compose up -d
test:
	go fmt ./...
	go test -vet all ./...

getURL:
	minikube service collatz-app-service --url

applyDeploy:
	kubectl apply -f k8s-deployment.yml

deploy:
		eval $(minikube docker-env)
		docker build -t collatz-app .
		kubectl apply -f k8s-deployment.yml

runScript:
	 bash command.sh $(n)