package utils

const (
	IndexTemplateName = `<html>
	<head>
	<title></title>
	</head>
	<body>
	<form action="/email-verifier/api/1.0/emails/verify" method="post">
	Email:<input type="email" name="email">
	<input type="submit" value="Submit">
	</form>
	</body>
	</html>`
)
