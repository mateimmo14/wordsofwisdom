# Words of Wisdom Project Documentation

## Project Setup
To set up the project, follow these steps:
1. **Clone the Repository**  
   Use `git clone https://github.com/mateimmo14/wordsofwisdom.git` to clone the project to your local machine.

2. **Install Dependencies**  
   Navigate to the project directory and run `npm install` to install the necessary dependencies.

3. **Environment Variables**  
   Create a `.env` file in the root of the project and add the required environment variables.
   Refer to the `.env.example` for the list of environment variables needed.

## Architecture
The project follows a modular architecture:
- **src/**: Contains the main source code.
- **tests/**: Contains the unit and integration tests.
- **public/**: Contains static assets like images and styles.

## Environment Variables
The following environment variables are required for the application:
- `DB_CONNECTION`: The database connection string.
- `PORT`: The port on which the application will run.
- `SECRET_KEY`: The secret key for JWT authentication.

## Usage Instructions
1. Run the application using `npm start`.
2. Access the application at `http://localhost:${process.env.PORT}`.
3. Follow the API documentation for endpoints and usage instructions.  

## Contributing
If you want to contribute to this project, feel free to open an issue or submit a pull request. Your contributions are welcome!