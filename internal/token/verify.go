package token


//func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
//	tokenString := extractToken(c)
//
//	token, err := jwt.Parse(tokenString, jwtKeyFunc)
//	if err != nil {
//		return nil, err
//	}
//
//	return token, nil
//}
//
//func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
//	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
//}
//
//// ParseAndVerifyJWT parses token and verifies authenticity with public key set via user pool
//func ParseAndVerifyJWT(t string) (*jwt.Token, error) {
//	// parse token
//	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		// fetch public JWK
//		keySet, err := GetKeySet()
//		if err != nil {
//			return nil, fmt.Errorf("Public key error: %v", err)
//		}
//		// retrieve kid from token header
//		kid, ok := token.Header["kid"].(string)
//		if !ok {
//			return nil, errors.New("kid header not found")
//		}
//		// Compare the local kid to the public kid lol
//		keys := keySet.LookupKeyID(kid)
//		if len(keys) == 0 {
//			return nil, fmt.Errorf("key %v not found", kid)
//		}
//
//		var raw interface{}
//		return raw, keys[0].Raw(&raw)
//	})
//	// not valid
//	if !token.Valid {
//		return nil, err
//	}
//
//	return token, nil
//}
