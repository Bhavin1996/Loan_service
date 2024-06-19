# Loan Service API

### This is a simple Loan Service API that manages loans, approvals, investments, and disbursements. The API is built using Go and the Gin framework, and includes email notifications for investors. The project is organized to be modular and maintainable.

## Table of Contents

- Features
- Requirements
- Setup
- Running the Server
- Testing the Application
- API Documentation
- Directory Structure
- License

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
git clone https://github.com/yourusername/loan-service.git

traversed into the cloned respository

2. Install Dependencies
Ensure you have Go modules enabled and then install dependencies:

go mod tidy

3. Configure SMTP Settings
Set up your SMTP configuration. Make sure that values are correct as per your SMTP cred values.


