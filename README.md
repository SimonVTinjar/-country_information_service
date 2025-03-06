# Country Information Service

## Introduction
The **Country Information Service** is a RESTful web service built in Go that provides detailed information about countries. This includes general details like population, capital, languages, neighboring countries, and a list of major cities. The service integrates with two external APIs:

- **REST Countries API** (http://129.241.150.113:8080/v3.1/)
- **CountriesNow API** (http://129.241.150.113:3500/api/v0.1/)

This documentation serves as a **user guide** explaining how to interact with the API and what responses to expect.

## Deployment on Render
The service is deployed on **Render**, making it accessible online. You can access the live API at:
```
https://your-render-app-url.onrender.com
```
(Replace `your-render-app-url` with the actual deployment URL.)

## API Endpoints and Usage
### 1. Get Country Information
- **Endpoint:** `/countryinfo/v1/info/{country_code}`
- **Method:** `GET`
- **Description:** Retrieves detailed information about a country based on its ISO 2-code.
- **Example Request:**
  ```sh
  curl -X GET "https://country-information-service-pack.onrender.com/countryinfo/v1/info/NO"
  ```
- **Example Response:**
  ```json
  {
    "name": "Norway",
    "continent": "Europe",
    "population": 5379475,
    "languages": {
      "nno": "Norwegian Nynorsk",
      "nob": "Norwegian Bokmål",
      "smi": "Sami"
    },
    "borders": ["FIN", "SWE", "RUS"],
    "flag": "https://flagcdn.com/w320/no.png",
    "capital": "Oslo",
    "cities": ["Abelvaer",
    "Adalsbruk",
    "Adland",
    "Agotnes",
    "Agskardet",
    "Aker","..."]
  }
  ```

### 2. Get Historical Population Data
- **Endpoint:** `/countryinfo/v1/population/{country_code}`
- **Method:** `GET`
- **Description:** Fetches historical population data for a given country.
- **Example Request:**
  ```sh
  curl -X GET "https://country-information-service-pack.onrender.com/countryinfo/v1/population/NO"
  ```

### 3. Check API Status
- **Endpoint:** `/countryinfo/v1/status`
- **Method:** `GET`
- **Description:** Returns the API status to ensure the service is operational.

## How to Run Locally
### Prerequisites
- Install **Go (1.18 or later)**.
- Ensure an active internet connection for API calls.

### Setup Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/SimonVTinjar/-Country_Information_Service.git
   ```
2. Navigate to the project folder:
   ```sh
   cd -Country_Information_Service
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the application:
   ```sh
   go run main.go
   ```
5. Use a tool like **Postman** or `curl` to test the API endpoints.

## How the Service Works
The service fetches data from two different APIs:
1. **REST Countries API** provides general country details like name, population, capital, flag, and neighboring countries.
2. **CountriesNow API** is used to fetch a **list of major cities** in the country.

These APIs are combined to provide a comprehensive response to the client.

## Error Handling
If an invalid country code is provided, the service returns an error message:
```json
{
  "error": "Invalid country code provided. Please check the ISO code."
}
```

## Hosting on Render
This service is hosted on **Render**, which automatically redeploys changes from GitHub. To update your deployed API:
1. Push new changes to GitHub:
   ```sh
   git add .
   git commit -m "Updated API functionality"
   git push origin main
   ```
2. Render will automatically detect and redeploy the changes.




