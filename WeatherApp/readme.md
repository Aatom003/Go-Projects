

---
# Weather App

This Weather App provides current weather information based on user-provided city names. It uses the OpenWeatherMap API to fetch weather data.

## Setup

### Prerequisites

- **Go Programming Language**: Ensure Go is installed on your system. If not, install it from [golang.org](https://golang.org/).
- **API Key**: Obtain an API key from [OpenWeatherMap API](https://openweathermap.org/api) (It's free for limited usage).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Aatom003/WeatherApp.git
   ```

2. Navigate to the project directory:

   ```bash
   cd WeatherApp
   ```

3. Set your OpenWeather API key:

   Replace `YOUR_OPENWEATHER_API_KEY` in `main.go` with your actual OpenWeatherMap API key.

4. Run the application:

   ```bash
   go run main.go
   ```

## Usage

1. Access the Weather App by visiting `http://localhost:8083` in your web browser.

2. Enter the name of the city for which you want to get weather information.

3. Click the **Submit** button.

4. The app will display the current weather information for the provided city.

## Folder Structure

- `main.go`: Go code for the web server handling the API requests.
- `static/`: Directory containing static files (HTML, CSS, JavaScript).
  - `index.html`: HTML file with the weather app UI.
  - `styles.css`: CSS file for styling the weather app UI.
  - `script.js`: JavaScript file handling form submission and dynamic content update.

## Modification

- **Customization**: Modify `styles.css` to change the UI appearance.
- **Functionality**: Extend `main.go` to handle additional data or improve error handling.
- **API Integration**: Explore and add more features using the OpenWeatherMap API.

## Dependencies

- **OpenWeatherMap API**: Used to fetch weather data.
- **Weather Icons Library**: The app uses the Weather Icons library for weather icons. You can find it [here](https://cdnjs.cloudflare.com/ajax/libs/weather-icons/2.0.12/css/weather-icons.min.css).

# Acknowledgments

- This Weather App is developed as a learning project and utilizes the OpenWeatherMap API for weather data retrieval. Feel free to use, modify, and enhance it according to your requirements.
---
