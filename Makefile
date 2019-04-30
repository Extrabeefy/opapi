build:
	protoc -I. --go_out=plugins=micro:. \
	  progression-service/proto/progression/progression.proto
	docker build -t progression-service C:\Users\nbrum\OneDrive\Desktop\GoWorkplace\src\github.com\opAPIProgression

run:
	docker run -p 50051:50051 progression-service