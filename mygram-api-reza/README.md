# My Gram Social-Media API 🖥️

My Gram API is an API that allows users to create, read, update, and delete photos.
Built with Go and PostgreSQL. This is a final project of Hacktiv8's Golang Scalable Web Service Bootcamp.

## Run the Application ▶️

This application requires Go and PostgreSQL to run. Make sure you have them installed on your machine before running the
application.
Please change the application configuration in `src/.env` to match your environment.

```bash
# Move to the workspace directory
cd workspace

# Clone the repository
git clone LINK GITHUB

# Move to the application directory
cd mygram-api-reza

# Install the go.mod depedencies
go get -u

# Build the application
go build -o .build/mygram-api-reza

# Run the application
.build/mygram-api-reza
```

## Endpoint-endpoint yang ada

Berikut adalah daftar endpoint-endpoint yang disediakan dalam aplikasi MyGram:

#### 1. User

```
Register [POST]: Mendaftarkan user baru.
Login [POST]: Proses masuk user.
Update User [PUT]: Proses edit user.
Delete User [Delete]: Proses penghapusan user.
```

#### 2. Photo

```
GetAll [GET]: Mendapatkan daftar semua foto.
GetOne [GET]: Mendapatkan satu foto berdasarkan ID.
CreatePhoto [POST]: Membuat foto baru.
UpdatePhoto [PUT]: Memperbarui informasi foto.
DeletePhoto [DELETE]: Menghapus foto.
```

#### 3. Comment

```
GetAll [GET]: Mendapatkan daftar semua komentar.
GetOne [GET]: Mendapatkan satu komentar berdasarkan ID.
CreateComment [POST]: Membuat komentar baru.
UpdateComment [PUT]: Memperbarui komentar.
DeleteComment [DELETE]: Menghapus komentar.
```

#### 3. Social Media

```
GetAll [GET]: Mendapatkan daftar semua media sosial.
GetOne [GET]: Mendapatkan satu media sosial berdasarkan ID.
CreateSocialMedia [POST]: Membuat media sosial baru.
UpdateSocialMedia [PUT]: Memperbarui informasi media sosial.
DeleteSocialMedia [DELETE]: Menghapus media sosial.
```

Dengan menyediakan endpoint-endpoint ini, MyGram memungkinkan pengguna

## API Documentation 📑

Access the API documentation at `http://localhost:3000/swagger/index.html` when the application is running.

## Tools ⚒️

- [Go](https://golang.org/) (Programming Language)
- [Gin](https://gin-gonic.com/) (Web Framework)
- [GORM](https://gorm.io/) (ORM)
- [PostgreSQL](https://www.postgresql.org/) (Database)
- [Swagger](https://swagger.io/) (API Documentation)

## Disclaimer ⚠️

The purpose of this project is to learn how to build a scalable web service using Go at Hacktiv8's Go Scalable Web Service Bootcamp.
Any plagiarism for this project is strictly prohibited and not my responsibility as the author of this repository.
