# Go Web App

This is a simple web application written in Go using the `net/http` package. It serves HTTP requests and provides a basic form for collecting feedback on essential tools and skills for cloud engineers. The application also serves static files, including HTML, CSS, and images.

## Features

- **Home Page**: Displays a list of essential tools and skills for cloud engineers and a form for user feedback.
- **Static Files**: Serves HTML, CSS, and images.
- **Form Handling**: Processes form submissions and displays a thank you message.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Running the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-web-app.git
   cd go-web-app
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Open your browser and visit `http://localhost:8080` to access the application.

### Project Structure

```
go-web-app/
│
├── static/
│   ├── css/
│   │   └── styles.css
│   ├── images/
│   │   └── golang-website(1).png
│   │   └── golang-website(2).png
│   └── index.html
│
└── main.go
```

- **`static/index.html`**: Contains the HTML for the feedback form and a list of cloud engineering tools and skills.
- **`static/css/styles.css`**: Contains the CSS for styling the page.
- **`static/images/golang-website.png`**: An image of the website.
- **`main.go`**: The main Go application file that sets up the server and handles requests.


## Looks like this

![Website](static/images/golang-website(1).png)![Website](static/images/golang-website(2).png)