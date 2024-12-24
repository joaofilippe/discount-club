# discount-eats (or similar)

[![Development Status](https://img.shields.io/badge/status-in%20development-orange)](https://github.com/joaofilippe/discount-club)
[![License](https://img.shields.io/github/license/joaofilippe/discount-club)](https://github.com/joaofilippe/discount-club/blob/main/LICENSE)

An API to connect customers with discounts at restaurants, bars, and snack bars.

## About

This project aims to create a robust and scalable API to connect customers with offers and discounts offered by restaurants, bars, and snack bars. The API will offer features such as:

*   User registration and authentication.
*   Establishment (restaurants, bars, snack bars) registration and management.
*   Publication and management of offers and discounts.
*   Search for establishments by location, cuisine type, etc.
*   Viewing establishment details, including menus and reviews.
*   Integration with other systems (to be defined, such as delivery systems).

The API will be developed using the [Echo](https://echo.labstack.com/) framework, a high-performance web framework for Go.

**Dockerization:** We plan to dockerize the application to facilitate deployment and ensure consistency across different environments (development, testing, and production). This will allow the API to run on any system that supports Docker.

**Name Consideration:** As the focus has shifted to food establishments, I suggest considering renaming the repository to something more appropriate, such as `discount-eats`, `food-discounts`, or similar.

## Technologies

*   [Go](https://go.dev/)
*   [Echo](https://echo.labstack.com/)
*   [Docker](https://www.docker.com/)
*   (Other dependencies will be added as development progresses)

## Project Status

Currently, the project is in the initial development phase. The basic repository structure has been created, and the API implementation with Echo is in progress.

**Next Steps:**

*   [ ] Initial Echo configuration.
*   [ ] Implementation of the first API endpoints (e.g., authentication, establishment registration).
*   [ ] Definition of the data model (including menu information, opening hours, etc.).
*   [ ] Database configuration (to be defined).
*   [ ] Creation of the Dockerfile and docker-compose.yml (when the application is in a more advanced state).
*   [ ] Unit and integration tests.

## How to Run (when available)

After the initial implementation, instructions for running the API locally (with and without Docker) will be added here. It will generally involve the following steps:

**Without Docker:**

1.  Clone the repository: `git clone https://github.com/joaofilippe/discount-club.git`
2.  Navigate to the project directory: `cd discount-club`
3.  Install dependencies: `go mod tidy` (or similar)
4.  Run the API: `go run main.go` (or similar)

**With Docker:**

1.  Clone the repository: `git clone https://github.com/joaofilippe/discount-club.git`
2.  Navigate to the project directory: `cd discount-club`
3.  Build the Docker image: `docker-compose build`
4.  Run the Docker container: `docker-compose up`

## Contributing

Contributions are welcome! If you want to contribute to this project, please follow these steps:

1.  Fork the repository.
2.  Create a branch with your feature: `git checkout -b feature/my-feature`
3.  Commit your changes: `git commit -m 'Adds my feature'`
4.  Push to your branch: `git push origin feature/my-feature`
5.  Open a Pull Request.

## License

This project is licensed under the [MIT](LICENSE) license.

## Contact

## Contato

**João Filippe Rossi Rodrigues**

*   LinkedIn: [https://www.linkedin.com/in/joaofilippe](https://www.linkedin.com/in/joaofilippe)
*   GitHub: [https://github.com/joaofilippe](https://github.com/joaofilippe)
*   Email: joaofilippe@outlook.com (mailto:joaofilippe@outlook.com)
