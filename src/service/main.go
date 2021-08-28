package main

/*
	- main.go should never be modified under this desgin.  It is self
	sufficient at this point and will rely upon the app core structure
	to manage things from here.

	- this design makes use of dependency injection as describe in app.go

	- this design is scaled to grow endpoints indefinitely regardless of
	router within server.go (pkg/server)

	- this design is functional at this point based upon a core configuration
	package (pkg/configurtaion) which can be scaled both horizontally and
	vertically
*/

// ---- main function ----
func main() {
	// let's get a nil app structure for our service
	a := App{}

	// use dependency injection to fill it in
	err := a.Init()
	// if Init fails with error ...
	if err != nil {
		// ... then, stop the train right here.
		panic(err)
	}

	// ... otherwise, all aboard ... let's rock some tracks.
	a.Run()
}