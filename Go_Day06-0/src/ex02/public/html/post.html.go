<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Places</title>
<meta name="description" content="">
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
	<ul>
		{{ range .Posts }}
			<li>
				<div>{{ .Created }}</div>
				<div>{{ .Header  }}</div>
				<div>{{ .Content }}</div>
			</li>
		{{ end }}
	</ul>

	{{ if ne .Page 1 }}
		<a href="/?page={{.Buttons.Previous}}">Previous</a>
	{{ end }}

	{{ if ne .Page .Buttons.Last}}
		<a href="/?page={{.Buttons.Next}}">Next</a>
	{{ end }}

	<a href="/?page={{.Buttons.Last}}">Last</a>
</body>
</html>

