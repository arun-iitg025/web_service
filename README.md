# Web Service Project

This is a Golang-based web service for fetching popular multiplayer game modes in specific areas. 

## Technologies
- **Programming Language**: Golang
- **Communication Protocol**: Protocol Buffers for model generation and data transfer objects (DTOs)
- **Database**: MongoDB (NoSQL)
- **Caching**: Redis for caching frequently requested data
- **Environment Handling**: Sensitive information is loaded from `.env` file
- **Containerization**: Docker and Docker Compose for consistent deployment
- **Build Management**: Makefile for easy build and run commands

## API Documentation

### Endpoint: `/popular-modes`

- **Method**: GET
- **Description**: Retrieves popular game modes for a given area code.
- **Query Parameters**:
  - `area_code` (string, required): The 3-digit area code identifying the location.
  
- **Example Request**:
  ```http
  GET http://localhost:8080/popular-modes?area_code=123
