\# Internal Transfers System



This is a backend system written in Go that facilitates internal financial transfers between accounts. It exposes HTTP APIs to:



\- Create accounts  

\- Check account balances  

\- Transfer funds between accounts  



PostgreSQL is used for data persistence.



---



\## Features



\- Account creation with initial balance  

\- Real-time balance querying  

\- Secure and consistent transaction processing  

\- Error handling for invalid inputs and transaction failures  



---



\## Tech Stack



\- Go 1.22+  

\- PostgreSQL  

\- Gorilla Mux  

\- curl (for testing HTTP endpoints)  



---



\## Project Structure



```

.

├── db/               # DB connection and migration

├── handlers/         # API handler logic

├── models/           # Account \& Transaction data models

├── main.go           # App entry point

├── go.mod / go.sum   # Dependency management

```



---



\## Setup Instructions



\### 1. Clone this repository



```bash

git clone https://github.com/SamridhiK108/internal-transfers-system.git

cd internal-transfers-system

```



\### 2. Setup PostgreSQL



Ensure you have PostgreSQL running locally and accessible. Then run:



```sql

CREATE DATABASE transferdb;

```



Update the DB connection string in `db/db.go` if you use a different user or password.



\### 3. Run the application



```bash

go run main.go

```



This starts the HTTP server on port 8080.



---



\## API Endpoints



\### Create Account



```bash

curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" -d "{\\"account\_id\\": 123, \\"balance\\": \\"500.75\\"}"

```



\### Check Account Balance



```bash

curl http://localhost:8080/accounts/123

```



\### Transfer Funds



```bash

curl -X POST http://localhost:8080/transactions -H "Content-Type: application/json" -d "{\\"source\_account\_id\\":123,\\"destination\_account\_id\\":456,\\"amount\\":\\"100.25\\"}"

```



---



\## Assumptions



\- All accounts use the same currency  

\- No authentication or authorization is implemented  

\- Transaction amounts and balances are stored as strings for precision  

\- Overdrafts are not allowed  

\- Account IDs must be unique  



---



\## Notes



\- Tables are created automatically when the application starts  

\- All responses are returned in JSON format  

\- The code handles common error scenarios such as insufficient balance or missing accounts  



---



\## Author



\*\*Samridhi Kohli\*\*  

GitHub: \[github.com/SamridhiK108](https://github.com/SamridhiK108)



