package main

import (
	//"fmt"
	"goaccontant/model"
	//"goaccontant/controller"
	"goaccontant/util"
	"goaccontant/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	model.Dbcheck()
	err := model.InitDB()
	if err != nil {
		panic(err)
	}
	err = util.RedidTestClient()
	if err != nil {
		panic(err)
	}
	//p, err := util.GenerateKeys()
	//fmt.Println(p)
	
	//user := model.User{UserName:"Sajad zare", Password:"123456",Email:"SajadZarephysics@gmail.com", Role:"member", Cash:[]model.Cash{{Amount:"100,000 rial", TypeCash:"income"}}}
	//created, err := controller.CreateUser(&user)
	//fmt.Println(created, err)
	/*db, err := model.GetDB()
	if err != nil {
		panic(err)
	}
	userz, err := controller.GetUser("user_name", "Majid Zare")
	fmt.Println(userz, err)
  db.Model(&userz).Association("Cash").Append(model.Cash{Amount:"200,000 rial", TypeCash:"spend"})*/
	/*privateCl := struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}{
		"user1",
		"admin",
	}*/
// 	tok, err := util.GenerateJWTSigned(privateCl)
// 	fmt.Println("Signed :", tok)
// 	errset := util.SetToRedis(privateCl.Name, tok)
// 	if errset != nil {
// 		panic(errset)
// 	}

// 	privateCl2 := struct {
// 		Name string `json:"name"`
// 		Role string `json:"role"`
// 	}{
// 		"user",
// 		"member",
// 	}
// 	tok2, err := util.GenerateJWTEncrypted(privateCl2)
// 	fmt.Println("Encrypted :", tok2)
// 	errset2 := util.SetToRedis(privateCl2.Name, tok2)
// 	if errset2 != nil {
// 		panic(errset2)
// 	}
// 	out2 := struct {
// 		Name string `json:"name"`
// 		Role string `json:"role"`
// 	}{}
// 	out1, erre := util.ParseEncryptedToken(tok2, &out2)
// 	if erre != nil {
// 		panic(erre)
// 	}
// 	fmt.Printf("iss: %s, sub: %s\n", out1.Issuer, out1.Subject)
// 	fmt.Printf("Name: %s, Role: %s\n", out2.Name, out2.Role)
// 	out1, ers := util.ParseJSONWebTokenClaims(tok, &out2)
// 	if ers != nil {
// 		panic(ers)
// 	}
// 	fmt.Printf("iss: %s, sub: %s\n", out1.Issuer, out1.Subject)
// 	fmt.Printf("Name: %s, Role: %s\n", out2.Name, out2.Role)

	server.NewServer()
}
