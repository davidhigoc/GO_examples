package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", handleRoot)
	server.Handle("POST", "/api", HandleHome)
	server.Listen()
}
