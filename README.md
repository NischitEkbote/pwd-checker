SETUP & RUN (USING DOCKER)

1. Clone the Repository

   git clone https://github.com/NischitEkbote/pwd-checker.git
   cd pwd-checker

   This downloads the project and navigates into the project directory.

2. Build the Docker Image

   docker build -t pwd-checker .

   This command:
   - Reads the Dockerfile in the current directory
   - Builds the Go binary inside a Docker container
   - Creates a Docker image named "pwd-checker"

3. Run the Password Checker

   docker run pwd-checker --pass "Hello123!"

   This starts a container from the pwd-checker image,
   passes the password as a command-line argument,
   and displays the password strength result.
