package function

import(
	// "github.com/dgrijalva/jwt-go"
)



// func CreateToken(userid uint64) (string, error) {
//   var err error
//   atClaims := jwt.MapClaims{}
//   atClaims["authorized"] = true
//   atClaims["user_id"] = userid
//   atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
//   at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
//   token, err := at.SignedString([]byte("secret"))
//   // should have in .env again, but i dont have time
//   if err != nil {
//      return "", err
//   }
//   return token, nil
// }
