# Currency Converter API

## Overview
This is a simple currency conversion API built using **Go** and the **HTTP framework**. It fetches real-time exchange rates from the ExchangeRate-API and allows users to convert currency values.

## Features
- Fetches live exchange rates from [ExchangeRate-API](https://www.exchangerate-api.com/).
- Converts currency based on user input.
- Uses **Gin** for routing and JSON response handling.
- Lightweight and efficient.

## Requirements
- Go (version 1.18+ recommended)
- An active internet connection
- ExchangeRate-API key (replace with your own key in the `API_URL` constant)

## Installation & Setup
1. Clone this repository:
   ```sh
   git clone https://github.com/your-repo/currency-converter-api.git
   cd currency-converter-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. The server will start at `http://localhost:8080`.

## API Endpoints
### Convert Currency
#### Request:
```http
GET /convert?base=USD&target=EUR&amount=100
```
#### Query Parameters:
- `base` - Base currency code (e.g., USD)
- `target` - Target currency code (e.g., EUR)
- `amount` - Amount to convert

#### Response:
```json
{
  "base": "USD",
  "target": "EUR",
  "rate": 0.85,
  "message": "1 USD = 0.85 EUR"
}
```

## Author
Developed by **SAURABH SINGH PATEL**. Contributions are welcome!

