package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var defaultPort = 9999

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /protected", protectedEndpoint)

	log.Println("Start test client on port: ", defaultPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request){
  authHeader := r.Header.Get("Authorization")
  if authHeader == "" {
	http.Error(w, "Missing JWT", http.StatusUnauthorized)
    return
  }
  token := authHeader[len("Bearer "):]
  
  err := verifyToken(token)
  if err != nil {
	log.Println(err)
    http.Error(w, "Invalid JWT", http.StatusUnauthorized)
    return
  }
  
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Protected resource")
}

func verifyToken(token string) error {
	jwt, err := jwt.Parse(token, func(jwt *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			return "", fmt.Errorf("missing JWT_SECRET_KEY in env vars")
		}
	   	return []byte(secretKey), nil
	})
   
	if err != nil {
	   return err
	}
   
	if !jwt.Valid {
	   return fmt.Errorf("invalid JWT")
	}
   
	return nil
 }
