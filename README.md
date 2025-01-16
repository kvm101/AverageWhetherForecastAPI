# AverageWheatherForecastAPI
## version 0.0.1-alfa (15.01.25)

**The goal of the project is to make an API for the average weather indicator. Thus, reducing the completely false forecast to zero.**

## Techologies:
- **Golang**
- **Postgres**
- **Redis (Cache)**

## Endpoints:
## Weather Forecast API Endpoints
1. **Get Current Weather**
   - **Endpoint**: `GET /weather/current`
   - **Description**: Get the current weather for a given city.
   - **Example Request**:
     ```http
     GET /weather/current?q=London
     ```
   - **Response**:
     ```json
     {
       "temperature": "15°C",
       "humidity": "75%",
       "condition": "Cloudy"
     }
     ```
2. **Get 7-Day Forecast**
   - **Endpoint**: `GET /weather/forecast`
   - **Description**: Get a 7-day weather forecast for the city.
   - **Example Request**:
     ```http
     GET /weather/forecast?q=London
     ```
   - **Response**:
     ```json
     [
       {
         "day": "2025-01-16",
         "temperature": "10°C",
         "condition": "Sunny"
       },
       {
         "day": "2025-01-17",
         "temperature": "12°C",
         "condition": "Cloudy"
       }
     ]
     ```

3. **Get 10-Day Forecast**
   - **Endpoint**: `GET /weather/forecast-10`
   - **Description**: Get a 10-day weather forecast for the city.
   - **Example Request**:
     ```http
     GET /weather/forecast-10?q=London
     ```
   - **Response**:
     ```json
     [
       {
         "day": "2025-01-16",
         "temperature": "10°C",
         "condition": "Sunny"
       },
       {
         "day": "2025-01-17",
         "temperature": "12°C",
         "condition": "Cloudy"
       }
     ]
     ```

4. **Get Weather Alerts**
   - **Endpoint**: `GET /weather/alerts`
   - **Description**: Get warnings about extreme weather conditions.
   - **Example Request**:
     ```http
     GET /weather/alerts?q=London
     ```
   - **Response**:
     ```json
     {
       "alerts": [
         {
           "event": "Heavy Rain",
           "description": "Expected heavy rainfall in the next 24 hours.",
           "severity": "High"
         }
       ]
     }
     ```

4. **Get Weather History**
   - **Endpoint**: `GET /weather/history`
   - **Description**: Get historical weather data for a given period.
   - **Example Request**:
     ```http
     GET /weather/history?q=London&start=2025-01-01&end=2025-01-10
     ```
   - **Response**:
     ```json
     [
       {
         "date": "2025-01-01",
         "temperature": "5°C",
         "condition": "Rain"
       },
       {
         "date": "2025-01-02",
         "temperature": "7°C",
         "condition": "Clear"
       }
     ]
     ```

## User API Endpoints
5. **Set User Preferences**
   - **Endpoint**: `POST /user/preferences`
   - **Description**: Set user preferences (e.g., selecting temperature units).
   - **Example Request**:
     ```http
     POST /user/preferences
     Content-Type: application/json
     {
       "unit": "metric",
       "language": "en"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Preferences updated successfully."
     }
     ```

6. **Get User Preferences**
   - **Endpoint**: `GET /user/preferences`
   - **Description**: Get the current user settings.
   - **Example Request**:
     ```http
     GET /user/preferences
     ```
   - **Response**:
     ```json
     {
       "unit": "metric",
       "language": "en"
     }
     ```


7. **Get User History**
   - **Endpoint**: `GET /user/history`
   - **Description**: Get the user's query history (e.g., cities viewed or weather).
   - **Example Request**:
     ```http
     GET /user/history
     ```
   - **Response**:
     ```json
     {
       "history": [
         { "city": "London", "date": "2025-01-10" },
         { "city": "Paris", "date": "2025-01-09" }
       ]
     }
     ```

## Authentication API Endpoints

8. **User Login**
   - **Endpoint**: `POST /auth/login`
   - **Description**: User authentication using a login and password.
   - **Example Request**:
     ```http
     POST /auth/login
     Content-Type: application/json
     {
       "username": "user1",
       "password": "password123"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Login successful",
       "token": "JWT_TOKEN"
     }
     ```

9. **User Registration**
   - **Endpoint**: `POST /auth/registration`
   - **Description**: New user registration.
   - **Example Request**:
     ```http
     POST /auth/registration
     Content-Type: application/json
     {
       "username": "newuser",
       "password": "newpassword123",
       "email": "newuser@example.com"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Registration successful"
     }
     ```

## Versions Info:
**version 0.0.1-alfa (15.01.25):**
- Added the basic idea of what the API should look like. `15.01.25`
- Added Endpoints `16.01.25`
- Added structs and redis functions `17.01.25`
