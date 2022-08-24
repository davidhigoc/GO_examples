package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", handleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()
}
