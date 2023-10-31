# Receipt Processor

## Submission by: Flannery Thompson

### Running the API locally
Clone the project directory to your local machine.

Navigate to `receipt-processor-challenge`.

Run `go run .` in your terminal. 

The API will be running at `localhost:8080`

Make your POST request to `localhost:8080/receipts/process` with a receipt JSON body. You may now use the returned `id` to hit GET `localhost:8080/receipts/{id}/points` to view the total points earned by the receipt.



### Assumptions
I am assuming that all receipts contain only the described fields.

I am assuming that receipts will always have at least one item on them.

I am assuming all requests made to the `localhost:8080/receipts/{id}/points` endpoint will have valid ids. 

All data will be lost on stopping the program.
