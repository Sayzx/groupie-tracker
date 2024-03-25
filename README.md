# Groupie Tracker Project Overview

Groupie Tracker was a 1.5-month project challenging us to leverage multiple APIs and tackle various request issues using JavaScript. We explored SQL through additional features and learned to handle APIs with Go.

## Technologies

### Front-End

- **HTML**: For page structure
- **CSS**: For page styling
- **JS**: For dynamic pages and effects

### Back-End

- **Go**: For data management and API calls
- **SQL**: For storing data such as comments

## Features

- **Multi-Part API Integration**: Managing four distinct APIs for details on artists, their concert venues and dates, and the links between these elements.
- **Intuitive Data Display**: Presenting artist data through various visualizations like blocks, cards, tables, and charts.
- **Advanced Filtering**: Users can filter artists by several criteria, such as creation date, first album date, member count, and concert locations, with range and multiple selection options.
- **Concert Geolocation**: Showing concert locations on a map by converting addresses to geographic coordinates.
- **Dynamic Search Bar**: A search function to easily find artists, band members, and venues, with autocomplete suggestions and indication of the suggestion type.

## Additional Features

- **Discord Login**: Site login via Discord for an enhanced user experience.
- **Comment System**: Allowing users to leave comments on artist pages, with data stored in SQL.
- **Artist Music Playback**: Integration of features for playing music directly from artist pages.
- **Light/Dark Mode**: Enabling users to switch between light and dark themes.

## Getting Started

To get the Groupie Tracker project up and running, ensure you have a Go environment set up on your machine. Follow these steps to install and launch the project:

### Prerequisites

- Go installed on your machine.
- Access to a terminal or command prompt.

### Installation and Execution

1. **Clone the repository** with the following command:
   ```bash
   git clone https://github.com/Sayzx/groupie-tracker/
   cd groupie-tracker
   go get
   go run cmd/main.go
   ```

### Custom Database Configuration (Optional)

To use your own SQL server with Groupie Tracker (this is optional and only recommended if you desire), please follow the steps below:

### Prepare the Database
First, prepare your database by executing the SQL script located in `internal/db/sql.sql`. This script establishes the necessary structure to store the project's data.

### Modify Connection Settings
Next, modify the database connection settings in the configuration file found at `internal/db/sqlgo`. At line 13, you will encounter the connection parameters such as username, password, server address, and database name.

Replace the default settings with those corresponding to your SQL database environment. An example of these modifications is shown below:

```sql
sql.Open("mysql", "user:password@tcp(IP:3306)/database")
```
👏 Visit http://localhost:8080 to explore the project. Enjoy!

### Authors : 

- [@sayzx](https://www.github.com/sayzx) [![wakatime](https://wakatime.com/badge/user/018d13a0-dea5-424f-9eef-3afdc71ebf87/project/018dacdc-1ebd-4c10-a5d7-2c183e8952c0.svg)](https://wakatime.com/badge/user/018d13a0-dea5-424f-9eef-3afdc71ebf87/project/018dacdc-1ebd-4c10-a5d7-2c183e8952c0)
- [@nicolasgouy](https://www.github.com/gonicolas12)

## Demo





![App](https://media.discordapp.net/attachments/1012749489402023956/1221833480799785060/image.png?ex=6614041a&is=66018f1a&hm=ccbaec249252d1e8c35092b38294b5a1a11881895aa69268bf5be4206ed860d3&=&format=webp&quality=lossless&width=1305&height=662)



