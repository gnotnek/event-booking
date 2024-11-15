# Event Booking System

## Overview
This project is an Event Booking System built with Golang. It allows users to browse, book, and manage event reservations. Its written using onion infrastructure

## Features
- User registration and authentication
- Event browsing and searching
- Booking and cancellation of events
- Admin panel for event management

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/gnotnek/event-booking.git
    ```
2. Navigate to the project directory:
    ```sh
    cd event-booking
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```
3. Make env file:
    ```
    DATABASE_HOST=
    DATABASE_PORT=
    DATABASE_USER=
    DATABASE_PASSWORD=
    DATABASE_NAME=

    JWT_SECRET_KEY=
    ```
4. Run the application:
    ```sh
    go run . api
    ```

## Usage
- Access the application at `http://localhost:8080`
- Register a new user or log in with existing credentials
- Browse available events and make bookings

## Contributing
Contributions are welcome! Please fork the repository and create a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For any inquiries, please contact [your email](mailto:youremail@example.com).

## Building the Application
To build the application, follow these steps:

1. Ensure you have Golang installed on your machine.
2. Navigate to the project directory:
    ```sh
    cd event-booking
    ```
3. Build the application:
    ```sh
    go build -o event-booking
    ```
4. The executable file `event-booking` will be created in the project directory. You can run it using:
    ```sh
    ./event-booking
    ```

## Building and Running with Docker
To build and run the application using Docker, follow these steps:

1. Build the Docker image:
    ```sh
    docker build -t event-booking .
    ```
2. Run the Docker container:
    ```sh
    docker run -p 8080:8080 event-booking
    ```