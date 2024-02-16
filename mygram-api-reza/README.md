# My Gram Social-Media API 🖥️

My Gram API is an API that allows users to create, read, update, and delete photos.
Built with Go and PostgreSQL. This is a final project of Hacktiv8's Golang Scalable Web Service Bootcamp.

# Deployed at :
- [RAILWAY](https://mygram-api-reza-production.up.railway.app/)

# API Documentation 📑

Access the API documentation at `https://mygram-api-reza-production.up.railway.app/swagger/index.html` when the application is running.

# POSTMAN:
- [Collection](https://github.com/rezazulf/bootcampbflpit/blob/main/mygram-api-reza/Hacktiv8%20Final%20Project%20GO%20MyGram.postman_collection.json) (The Collection)
- [Postman Workspace](https://www.postman.com/rezazulf/workspace/hacktiv8/collection/12473257-91467b6c-5f71-4a35-859b-44ca50000004?action=share&creator=12473257&active-environment=12473257-be5307ab-1540-437a-a02a-6ad5ae4fb837) (The workspace)

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
UpdateUserData [PUT]: Proses edit user.
DeleteUser [Delete]: Proses penghapusan user.
```

#### 2. Photo

```
GetAllPhotos [GET]: Mendapatkan daftar semua foto.
PostPhoto [POST]: Membuat foto baru.
UpdatePhoto [PUT]: Memperbarui informasi foto.
DeletePhoto [DELETE]: Menghapus foto.
```

#### 3. Comment

```
GetAllComments [GET]: Mendapatkan daftar semua komentar.
PostComment [POST]: Membuat komentar baru.
UpdateComment [PUT]: Memperbarui komentar.
DeleteComment [DELETE]: Menghapus komentar.
```

#### 3. Social Media

```
GetAllSocialMedia [GET]: Mendapatkan daftar semua media sosial.
AddSocialMedia [POST]: Membuat media sosial baru.
EditSocialMediaData [PUT]: Memperbarui informasi media sosial.
DeleteSocialMedia [DELETE]: Menghapus media sosial.
```

## Tools ⚒️

- [Go](https://golang.org/) (Programming Language)
- [Gin](https://gin-gonic.com/) (Web Framework)
- [GORM](https://gorm.io/) (ORM)
- [PostgreSQL](https://www.postgresql.org/) (Database)
- [Swagger](https://swagger.io/) (API Documentation)

## Disclaimer ⚠️

The purpose of this project is to learn how to build a scalable web service using Go at Hacktiv8's Go Scalable Web Service Bootcamp.
Any plagiarism for this project is strictly prohibited and not my responsibility as the author of this repository.
