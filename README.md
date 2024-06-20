# Loan Service API

### This is a simple Loan Service API that manages loans, approvals, investments, and disbursements. The API is built using Go and the Gin framework, and includes email notifications for investors. The project is organized to be modular and maintainable.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Setup](#setup)
- [Running the Server](#running-the-server)
- [Testing the Application](#testing-the-application)


## Features

- Create, approve, invest in, and disburse loans.
- Validations for loan creation and investment.
- Email notifications to investors when the loan is fully invested.

## Requirements

- Go 1.20 or later
- Git
- A running SMTP server or SMTP service credentials for email notifications.

## Setup

1. Clone the Repository -
git clone https://github.com/Bhavin1996/Loan_service.git

traverse into the cloned respository

2. Install Dependencies
- Ensure you have Go modules enabled and then install dependencies via below command
- go mod tidy

3. Configure SMTP Settings to Set up your SMTP configuration.
#### Make sure that values are correct as per your SMTP credential values.
Example values like below :
- SMTP_SERVER=smtp.example.com
- SMTP_PORT=587
- SMTP_USER=your-email@example.com
- SMTP_PASSWORD=your-email-password
- FROM_EMAIL=your-email@example.com

## Running the server

- go run main.go to start the main server

## Testing the Application

Unit tests are provided for the core functionalities. 
- To run the tests use command - go test ./...
- Ensure to follow the comments as the test result might not come to PASS
- Make sure that some values are setup in prior before running the tests
  
Above command is to run all test in the project


