run-frontend:
	cd frontend && npm run dev

run-backend:
	cd backend && go run cmd/main.go

docker-build-backend:
	cd backend && docker build -t backend .

docker-run-backend:
	cd backend && docker run -p 8080:8080 backend 

docker-build-frontend:
	cd frontend && docker build -t frontend .

docker-run-frontend:
	cd frontend && docker run -p 3000:80 frontend

run-project:
	docker-compose up
