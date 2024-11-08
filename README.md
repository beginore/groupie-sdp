# Groupie Tracker

## Project Overview
**Groupie Tracker** is a web application developed to display information about musical artists and their concert details, including locations, dates, and interconnections. The application fetches data from a RESTful API and presents it in an interactive and user-friendly format, allowing users to search for artists, view their profiles, and explore their concert locations and dates.

## Features
- **Artist Profiles**: View detailed information about musical artists, including their name, members, first album, and creation date.
- **Concert Locations**: See concert dates and locations for each artist.
- **Search Functionality**: Search for artists, members, albums, or concert locations.
- **Data Visualization**: Information is displayed in an organized, easy-to-navigate format.

## Technology Stack
- **Backend**: Go (Golang) – for server logic, API calls, and data processing.
- **Frontend**: HTML, CSS – for a responsive user interface.
- **Data Format**: JSON – for structured storage and handling of API responses.

## Software Architecture and Design Patterns
### Architecture Pattern
- **MVC (Model-View-Controller)**: 
  - **Model** handles API data retrieval and processing.
  - **View** displays information using HTML templates.
  - **Controller** manages requests, routes, and data flow between Model and View.
- **Justification**: The MVC architecture keeps a clean separation between data processing, UI rendering, and business logic, making the application scalable and maintainable.

### Design Patterns
- **Singleton Pattern**: Manages API client instance, ensuring consistent data retrieval and efficient resource use.
- **Prototype Pattern**: Creates copies of artist data without reinitialization, enhancing performance for displaying multiple instances.
- **Facade Pattern**: Simplifies complex API calls and data handling by consolidating them into a single interface.
- **Strategy Pattern**: Handles different search criteria in a flexible, modular way, allowing for easier search functionality updates.

## Installation
1. **Clone the repository and run the local server:**
   ```bash
   git clone https://github.com/beginore/groupie-sdp.git


    cd groupie-tracker

    cd .\cmd\

    go run .
    ```
