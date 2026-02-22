# 🎓 CS367 Lab 4: REST API Enhancement with Gin and SQLite

This project is a REST API for managing student records, built with **Go** and the **Gin Framework**. It uses **SQLite** as the database and follows a **Clean Layered Architecture** for better maintainability.

## 🚀 Features
- **CRUD Operations:** Complete Create, Read, Update (Challenge 1), and Delete (Challenge 2) functionalities.
- **Input Validation:** Ensures data integrity (e.g., GPA must be between 0.00 and 4.00, required fields cannot be empty) (Challenge 3).
- **Standardized Error Handling:** Returns proper HTTP status codes (400, 404, 500) and clear JSON error messages instead of raw database errors (Challenge 4).
- **Clean Structure:** Separated into Handlers, Services, Repositories, and Models layers (Bonus).

---

## 🛠️ How to Run

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/Kindech/CS367-Lab4-GO-GinAPI.git](https://github.com/Kindech/CS367-Lab4-GO-GinAPI.git)
   cd CS367-Lab4-GO-GinAPI
2. **Install dependencies:**
   ```bash
   go mod tidy
3. **Run the application:**
   ```bash
   go run main.go
   (The server will start on http://localhost:8080.)
   (The students.db SQLite database file will be created automatically upon the first run.)

---

## 📡 API Endpoints
|  Method  |   Endpoint	   |               Description	                 |                Status Codes                |
|----------|---------------|---------------------------------------------|--------------------------------------------|
|   GET	   | /students	   |    Retrieve a list of all students	         |   200 OK                                   |
|   GET	   | /students/:id |	  Retrieve a specific student by ID	       |   200 OK, 404 Not Found                    |
|   POST	 | /students	   |    Create a new student	                   |   201 Created, 400 Bad Request             |
|   PUT	   | /students/:id |	  Update an existing student's information |	 200 OK, 400 Bad Request, 404 Not Found   |
|   DELETE | /students/:id |	  Delete a student by ID	                 |   204 No Content, 404 Not Found            |

**Example Request Body (POST / PUT)**
```json
{
  "id": "6609650160",
  "name": "Kitchaya Kindech",
  "major": "Computer Science",
  "gpa": 3.80
}

