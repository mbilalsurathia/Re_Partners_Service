This project provides a solution for efficiently fulfilling customer orders using predefined pack sizes.
It ensures that orders are met using the minimum number of packs while also minimizing the total number of items sent.
The goal is to optimize order fulfillment while adhering to specific constraints.

Features
Dynamic Pack Sizes: Users can input custom pack sizes for their orders.
Minimum Pack Utilization: The algorithm calculates the optimal combination of packs to minimize both the number of packs used and the total items sent.
Order Fulfillment: It provides feedback on whether an order can be fulfilled based on the available pack sizes.
Requirements
Go (version 1.15 or later) for running the backend calculations.
HTML and JavaScript for the frontend input of pack sizes.
Usage
Running the Application
Clone the repository:
```
git clone https://github.com/your-username/pack-size-calculation.git
cd pack-size-calculation
```

Compile and run the Go application:
```
go run main.go
```

With Docker
Ensure you have Docker installed and running on your machine.

Build the Docker image:
```
docker build -t pack-size-calculation .
```
Run the Docker container:
```
docker run -p 8080:8080 pack-size-calculation
```
Access the application by navigating to http://localhost:8080 in your web browser.

Open index.html in a web browser to access the frontend interface.

Using the Makefile
If you prefer using a Makefile for easier commands, you can do the following:

Run the following command to build the Go application:
```
make build
```
To Build with Docker
```
make docker-build
```
To Run with Docker
```
make docker-run
```


Submitting Pack Sizes
To update pack sizes:

Enter the desired pack sizes in the input fields provided.
Click the "Submit" button. The application will alert you with the updated pack sizes or notify you if no valid sizes were entered.
Sample Inputs
Example orders can be tested within the Go application:

Order for 1 item
Order for 250 items
Order for 251 items
Order for 501 items
Order for 12001 items
Expected Output
The application will provide output in the following format for each order:

```
Order for X items: 
- 1x250
- 1x500
- 2x5000, 1x2000, 1x250
```
