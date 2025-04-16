package jwt

import "testing"

func TestJWTCreate(t *testing.T) {
	const email = "test@test.ru"
	jwtService := NewJWT("/2+XnmJGz1j3ehIVI/5P9k1l+CghrE3DcS7rnt+qar5w=")
	token, err := jwtService.Create(JWTData{Email: email})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not eqal %s", data.Email, email)
	}
}
