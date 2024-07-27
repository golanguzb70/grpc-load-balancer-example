copy_proto:
	rm -rf ./api-gateway/protos/* && cp -r ./proto/* ./api-gateway/protos/
	rm -rf ./server/protos/* && cp -r ./proto/* ./server/protos/

start: 
	cd dev-ops && docker-compose up --build -d 

stop:
	cd dev-ops && docker-compose down