# User Net Worth Calculator

Welcome to the User Net Worth Calculator! This project helps individuals retrieve and calculate their net worth by programmatically accessing relevant financial accounts and consolidating investment information.

Current sources integrated
- brokers (zerodha, angelone,idirect)
- mutual fund (mfcentral)

In pipeline
- PF data 
- Bank Accounts
- Digital Gold
- US stocks

## Getting Started

Follow these steps to set up and use the User Net Worth Calculator:

### Prerequisites

1. **PostgreSQL Database**: Ensure you have a PostgreSQL database set up, as the application relies on it for data storage.

2. **Database Setup**: Run the `create.sql` script to set up the necessary tables in your PostgreSQL database.

   ```
   psql -U your_username -d your_database -a -f root/pkg/db/sql/create.sql
   ```
3. **Broker Signup**: Ensure you have a TOTP token with your broker. Supported broker are zerodha, angelone, idirect.
4. **MF Central Signup**: Ensure you have a MF Central account to login for data fetch.
   
### Update Credentials
Update your credentials to enable secure communication with the financial institutions. The credentials are stored in /sources/broker/sync.

1. Use API <base>/sources/broker/sync.
2. Update the necessary configuration files with your credentials.
3. Configure Callback URLs if needed with broker/connecting source


### Running the Application
Build the Golang application:

```
go build -o user-net-worth main.go
```
run the application:
```
./user-net-worth
```
Access the application at http://localhost:8080 in postman. (FE integration in pipeline)

### Contribution
Contributions are welcome! If you find any issues or have suggestions, please create a new issue or submit a pull request.

### Access
I intend to keep this project public.
Feel free to use, modify, and distribute the code as needed. If you have any questions or need further assistance, please contact the project maintainers.
