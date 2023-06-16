{{define "newPost"}}
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Login page</title>
</head>
<body>
<form id="newPostForm" name="newPostForm" action="/newPost" method="post" class="mt-4">
<label for="header">Заголовок</label><br>
<input type="header" id="header" name="header" autocomplete="on"><br>
<label for="content">Содержание</label><br>
<input type="content" id="content" name="content" autocomplete="on"><br>
<br>
<button type="submit" name="submitBtn">Войти</button>
</form>
{{if . }}
<div>
{{.Message}}
</div>
{{end}}
<br>
<a href="/newPost">Зарегистрироваться</a>
</body>
</html>
{{end}}