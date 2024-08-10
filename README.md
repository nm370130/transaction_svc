In transaction service using inmemory storage for storing data for now, there are 5 REST apis, Here , You can directly run this on your local machine



Running the Service
To run the service, use the following command:

go run main.go
The service will be running on http://localhost:8080.

Example API Calls

Adding a Transaction:
curl -X PUT -d '{"amount":5000,"type":"cars"}' http://localhost:8080/transactionservice/transaction/10


Adding a Child Transaction:
curl -X PUT -d '{"amount":10000,"type":"shopping","parent_id":10}' http://localhost:8080/transactionservice/transaction/11


Get a Transaction:
curl http://localhost:8080/transactionservice/transaction/10


Get Transactions by Type:
curl http://localhost:8080/transactionservice/types/cars


Get Sum of Linked Transactions:
curl http://localhost:8080/transactionservice/sum/10
