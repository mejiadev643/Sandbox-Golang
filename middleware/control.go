package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Haz lo que quieras aquí antes de pasar la solicitud al siguiente middleware o controlador
		fmt.Println("Request received: ", r.Method, r.URL.Path)
		token := GetToken(r.Header.Get("Authorization"))
		if token != "" {
			fmt.Println("Token: ", token)
		}
		payload, _ := GetPayload(token)
		fmt.Println(payload)
		//obtener un permiso del token y recorrerlo
		permiso := payload["permisions"]

		//recorrer con un bucle for para comparar los permisos
		for _, v := range permiso.([]interface{}) {
			//fmt.Println(v)
			if v == "read" {
				fmt.Println("Tiene permiso de lectura")
				next.ServeHTTP(w, r)
				return
			}
		}
		//fmt.Println(permiso)
		// Llama al siguiente middleware o controlador

		//next.ServeHTTP(w, r)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))

		// Haz lo que quieras aquí después de que la solicitud haya sido manejada( noe s necesariio en este caso)
		fmt.Println("Request handled: ", r.Method, r.URL.Path)
	})
}

func GetToken(token string) string {
	authHeader := token
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

func GetPayload(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
