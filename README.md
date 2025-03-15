# URL Shortener in Golang 🚀  

This is a simple URL Shortener built using **Golang**, which allows users to generate short URLs and redirect them to their original destinations.  

## Features 🌟  
- Generates an 8-character short URL using **MD5 hashing**  
- Stores URL mappings in an **in-memory database (map)**  
- Supports **creating short URLs** via a POST request  
- Supports **redirecting to the original URL** via a GET request  
- Lightweight and runs on **port 3000**  

## Endpoints 🔗  

### 1️⃣ **Create a Short URL**  
**Endpoint:**  
```http
POST /shortner
```  
**Request Body:**  
```json
{
  "url": "https://example.com/long-url"
}
```  
**Response:**  
```json
{
  "short_url": "5d41402a"
}
```  

### 2️⃣ **Redirect to Original URL**  
**Endpoint:**  
```http
GET /redirect/{short_url}
```  
Example:  
```http
GET /redirect/5d41402a
```  
🔄 Redirects the user to `https://example.com/long-url`  

## How It Works ⚙️  
1. The **`generateShortURL`** function creates an 8-character hash from the original URL.  
2. The **`createURL`** function stores the mapping in an in-memory map.  
3. The **`shortURLHandler`** endpoint generates and returns the short URL.  
4. The **`redirectURL`** endpoint retrieves the original URL and redirects the user.  

## Running the Server ▶️  
1. Clone this repository  
2. Run the server  
```sh
go run main.go
```  
3. The server starts on `http://localhost:3000/`  

## Improvements & Next Steps 🚀  
- Use a **database (Redis/PostgreSQL)** for persistence  
- Implement **custom short URLs**  
- Add **rate limiting and authentication**  

---
Contributions & feedback are welcome! 🛠️🎉
