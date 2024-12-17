# CRUD API

This project is a simple CRUD (Create, Read, Update, Delete) API built using Go, Gin, and MongoDB. It allows you to manage tasks with basic operations.

## Prerequisites

- Go 1.23.4 or later
- MongoDB
- Git

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/said7388/Golang-Mongodb-crud.git
   cd crud-api
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**

   Create a `.env` file in the root directory with the following content:

   ```plaintext
   PORT=8080
   MONGODB_URI=mongodb://localhost:27017
   DB_NAME=crudDB
   ```

   Adjust the `MONGODB_URI` and `DB_NAME` as needed for your MongoDB setup.

4. **Run the application:**

   ```bash
   go run main.go
   ```

   The server will start on the port specified in the `.env` file (default is 8080).

## API Endpoints

### Task Management

- **Get All Tasks**

  - **Endpoint:** `GET /tasks/`
  - **Description:** Retrieves a list of all tasks.
  - **Response:** JSON array of tasks.

- **Get Task by ID**

  - **Endpoint:** `GET /tasks/:id`
  - **Description:** Retrieves a task by its ID.
  - **Response:** JSON object of the task.

- **Create Task**

  - **Endpoint:** `POST /tasks/`
  - **Description:** Creates a new task.
  - **Request Body:**
    ```json
    {
      "title": "Task Title",
      "content": "Task Content"
    }
    ```
  - **Response:** JSON object of the created task.

- **Update Task**

  - **Endpoint:** `PUT /tasks/:id`
  - **Description:** Updates an existing task by its ID.
  - **Request Body:**
    ```json
    {
      "title": "Updated Title",
      "content": "Updated Content"
    }
    ```
  - **Response:** Success message.

- **Delete Task**

  - **Endpoint:** `DELETE /tasks/:id`
  - **Description:** Deletes a task by its ID.
  - **Response:** Success message.

## Project Structure

- `main.go`: Entry point of the application.
- `db/mongo.go`: MongoDB connection setup.
- `controllers/task_controller.go`: Handlers for task-related API endpoints.
- `repository/task_repository.go`: Database operations for tasks.
- `models/task.go`: Task model definition.
- `routes/task_routes.go`: API route definitions.

## License

This project is licensed under the MIT License.
