package main

import (
	"fmt"
	"goaccontant/model"
	"goaccontant/util"
	"net/http"

	"github.com/gin-gonic/gin"
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
	rout := gin.Default()

	rout.GET("/", func(c *gin.Context) {

		// gin.H is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//rout.Run(":9000")
	//user := model.User{Name:"Zahra zare", Age:"26", Email:"Zahrazareephysics@gmail.com", Role:"member", MemberNumber:"4311229569"}
	//created, err := controller.CreateUser(&user)
	//controller.GetUser( "5080075066")
	//fmt.Println(created, err)
	privateCl := struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}{
		"user1",
		"admin",
	}
	tok, err := util.GenerateJWTSigned(privateCl)
	fmt.Println("Signed :", tok)
	errset := util.SetToRedis(privateCl.Name, tok)
	if errset != nil {
		panic(errset)
	}

	privateCl2 := struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}{
		"user",
		"member",
	}
	tok2, err := util.GenerateJWTEncrypted(privateCl2)
	fmt.Println("Encrypted :", tok2)
	errset2 := util.SetToRedis(privateCl2.Name, tok2)
	if errset2 != nil {
		panic(errset2)
	}
	out2 := struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}{}
	out1, erre := util.ParseEncryptedToken(tok2, &out2)
	if erre != nil {
		panic(erre)
	}
	fmt.Printf("iss: %s, sub: %s\n", out1.Issuer, out1.Subject)
	fmt.Printf("Name: %s, Role: %s\n", out2.Name, out2.Role)
	out1, ers := util.ParseJSONWebTokenClaims(tok, &out2)
	if ers != nil {
		panic(ers)
	}
	fmt.Printf("iss: %s, sub: %s\n", out1.Issuer, out1.Subject)
	fmt.Printf("Name: %s, Role: %s\n", out2.Name, out2.Role)

}
