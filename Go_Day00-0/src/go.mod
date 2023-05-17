module main

replace parser => ./parser

replace mean => ./metrics/mean

replace median => ./metrics/median

replace mode => ./metrics/mode

replace sd => ./metrics/sd

go 1.19

require (
	mean v0.0.0-00010101000000-000000000000 // indirect
	median v0.0.0-00010101000000-000000000000 // indirect
	mode v0.0.0-00010101000000-000000000000 // indirect
	parser v0.0.0-00010101000000-000000000000 // indirect
	sd v0.0.0-00010101000000-000000000000 // indirect
)
