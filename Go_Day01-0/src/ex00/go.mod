module main

replace DBReader => ./DBReader

replace MyJson => ./DBReader/MyJson

replace MyXml => ./DBReader/MyXml

go 1.19

require (
	DBReader v0.0.0-00010101000000-000000000000 // indirect
	MyJson v0.0.0-00010101000000-000000000000 // indirect
	MyXml v0.0.0-00010101000000-000000000000 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
)
