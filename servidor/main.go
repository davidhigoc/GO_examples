package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", handleRoot)
	server.Handle("POST", "/api", HandleHome)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Listen()
}
