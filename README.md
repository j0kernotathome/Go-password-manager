# Password manager server on golang

for starting server you need starting postgres container by this command
docker-compose up -d
after start 2 service:
go run UserService/internal/cmd/main.go
go run PasswordService/internal/cmd/main.go

# Endpoints
this http server have 5 endpoints:
/get_user/:login
/create_user
/create_api_key
/add_password
/get_passwords
