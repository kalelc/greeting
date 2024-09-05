build:
	docker build  --platform linux/arm64 -f Dockerfile -t greeting .
tag:
	docker tag greeting:latest kalelc/greeting:latest
push:
	docker push kalelc/greeting
