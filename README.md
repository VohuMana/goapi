# goapi
Generate the Go structs, used to serialize/deserialize JSON schemas, without having to type a single line of Go code.

## Installing
1. Clone this repo
2. Run `go build`
3. Done.

## Usage
1. Create a text file with the JSON schema you want to generate the Go structs for.
2. Run `goapi -inputfile <YOUR JSON SCHEMA FILE> -filename <OUTPUT GO FILE> -packagename <NAME OF THE PACKAGE YOU WANT IN THE FILE e.g. main> -usekeynames`
3. Done!

## Flags
`inputfile` - This is the input JSON schema you want to generate Go structs for.
`filename` - The output filename, this is a file that will be created. Examples include `myapi.go`, `WeatherApi.go`, etc.
`packagename` - The name of the package the new file will be a part of. Examples include `main`, `api`, `webapi`, etc.
`usekeynames` - This tells the program to name structs based off the key names in the JSON schema, if this is not specified the program will use randomly generated strings for type names.  The random strings will need to be replaced before compiling but a simple find and replace should work.

## Issues
File any issues or feature requests to this repo and I'll get to them when I can.
