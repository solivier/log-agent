package main

func main() {
	err := Ingest("id", 1234, "1234", "1234", "action", "context")
	if err != nil {
		return
	}
}
