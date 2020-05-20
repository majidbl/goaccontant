package server

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"goaccontant/model"
	"goaccontant/controller"
	"goaccontant/util"
	"reflect"
	"strconv"
	
)


type Clim struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		Authorized  bool   `json:"authorized"`
	}

func NewServer(){
  router := gin.Default()
	//new template engine
	//router.Use(authMiddelware)
	router.HTMLRender = ginview.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", getIndex)
	router.GET("/register", getRegister)
	router.POST("/register", postRegister)
	router.GET("/login", getLogin)
	router.POST("/login",postLogin)
	router.GET("/page", authMiddelware ,getPage)
	router.GET("/ping", getPing)
	router.POST("/ping", postPing)
	router.GET("/cash/:name", getCashUser)
	router.GET("/usercash", authMiddelware, getUserCash)
	router.POST("/usercash", authMiddelware, postUserCash)
	router.GET("/delete/:id", authMiddelware, deleteCash)
	router.Run(":4321")
	
}


func getIndex(ctx *gin.Context){
  session := sessions.Default(ctx)
  session.Set("name","majid")
  successmsg := session.Flashes("success")
  dangermsg := session.Flashes("danger")
  infomsg := session.Flashes("info")
  session.Delete("success")
  session.Delete("danger")
  session.Delete("info")
  session.Save()
  //fmt.Println(session)
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
			"SuccessMsg": successmsg,
			"DangerMsg": dangermsg,
			"InfoMsg": infomsg,
		})
}

func getRegister(ctx *gin.Context){
		//render with master
		//session := sessions.Default(ctx)
		//name := session.Get("name")
		//fmt.Println(name)
		ctx.HTML(http.StatusOK, "register", gin.H{})
}

func getLogin(ctx *gin.Context){
  session := sessions.Default(ctx)
  //session.Set("name","majid")
  successmsg := session.Flashes("success")
  dangermsg := session.Flashes("danger")
  infomsg := session.Flashes("info")
  fmt.Println("info session added",session.Flashes("info"))
  session.Delete("success")
  session.Delete("danger")
  session.Delete("info")
  session.Save()
  //fmt.Println(session)
		//render with master
		ctx.HTML(http.StatusOK, "login", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
			"SuccessMsg": successmsg,
			"DangerMsg": dangermsg,
			"InfoMsg": infomsg,
		})
  //session := sessions.Default(ctx)
  /*cookie, err := ctx.Cookie("gin_cookie")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Cookie value: %s \n", cookie)
  fmt.Println(reflect.TypeOf(cookie))*/
}

func postLogin(ctx *gin.Context){
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
	  session := sessions.Default(ctx)
		//fmt.Println("email get from form:. ",email)
		//fmt.Println("password get from form:  ",password)
		user, errg := controller.GetUser("email", email)
		session.Delete("info")
		session.Delete("danger")
		session.Delete("success")
		if errg != nil {
		  panic(errg)
		  session.AddFlash("User with this Email Not exist","danger")
		  session.AddFlash("Please Register to access This Page ", "info")
		  ctx.Redirect(http.StatusFound, "/")
		  }
		  //fmt.Println(user)
		if email == "majidzarephysics@gmail.com" && password == "123456" {
		    //fmt.Println("check user exist")
		    cl := Clim{Email:email, Name: user.UserName ,Role:user.Role, Authorized: true}
		    token, errt := util.GenerateJWTSigned(&cl)
		    if errt != nil {
		      panic(errt)
		      
		    }
		    //fmt.Println("token GenerateJWTSigned",  token)
		    ctx.SetCookie("gin_cookie", token, 60*60*24, "/", "localhost", false, true)
		    session.Set("Role", user.Role)
		    session.Set(fmt.Sprintf("Authenticated_%s", user.UserName), "true")
		    session.Set("username", user.UserName)
		    session.AddFlash("You are logged successfully","success")
		    session.Save()
		    ctx.Redirect(http.StatusFound, "/")//gin.H{"MsgsSuccess":session.Flashes("success"),}
		    
		  }else{
		    session.AddFlash("This User Does not Privilege To login", "info")
		    ctx.Redirect(http.StatusOK, "/")
		    
		  }
		}

func postRegister(ctx *gin.Context){
		email := ctx.PostForm("email")
		name := ctx.PostForm("name")
		role := ctx.PostForm("role")
		password := ctx.PostForm("password")
		user := model.User{UserName:name,Password:password, Email:email, Role:role}
		session := sessions.Default(ctx)
		session.Delete("info")
		session.Delete("danger")
		session.Delete("success")
		//session.Clear()
		_, u := controller.UniqueCheck(email)
		fmt.Println(u)
		if !u {
		  controller.CreateUser(&user)
		  session.AddFlash("User With This Email Registered Before","danger")
		  session.AddFlash("User Creating going failed","danger")
		  session.Save()
		  ctx.Redirect(http.StatusFound, "/")//gin.H{"MsgsDanger":session.Flashes("danger"),})
		} else{
		//fmt.Println(created)
		session.AddFlash("User Created Successfully ","success")
		session.Save()
		ctx.Redirect(http.StatusFound, "/")
		}
		//fmt.Println(session.Flashes("success"))
		//.Println(session.Flashes("danger"))
}

func getPage(ctx *gin.Context){
		//render only file, must full name with extension
		cookie, err := ctx.Cookie("gin_cookie")
		var customClime Clim
		defaultClime, err := util.ParseJSONWebTokenClaims(cookie, &customClime)
		fmt.Println(err)
		fmt.Println(customClime.Email)
		fmt.Println(defaultClime.Issuer)
		session := sessions.Default(ctx)
    session.Set("name","majid")
    successmsg := session.Flashes("success")
    dangermsg := session.Flashes("danger")
    infomsg := session.Flashes("info")
		fmt.Println(session.Get("username"))
    session.Delete("success")
    session.Delete("danger")
    session.Delete("info")
    session.Save()
    //fmt.Println(session)
		//render with master
		ctx.HTML(http.StatusOK, "page.html", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
			"SuccessMsg": successmsg,
			"DangerMsg": dangermsg,
			"InfoMsg": infomsg,
		})
}
func getPing(ctx *gin.Context){
		//render only file, must full name with extension
		users, err := controller.GetAllUser()
		fmt.Println(users[0].UserName)
		if err != nil {
		  panic(err)
		}
		//session := sessions.Default(ctx)
		//fmt.Println(session.Get("username"))
		ctx.JSON(http.StatusOK, users)

}
func postPing(ctx *gin.Context){
		//render only file, must full name with extension
		var mapJson Clim
		b := ctx.Request.Body
		err := json.NewDecoder(b).Decode(&mapJson)
		if err != nil {
		  panic(err)
		}
		ctx.JSON(http.StatusOK, b)

}

// getCashUser represent cash for each user
func getCashUser(ctx *gin.Context){
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)


}
func loguot(ctx *gin.Context){
  session := sessions.Default(ctx)
  session.Clear()
  ctx.Redirect(http.StatusFound, "/")
}
func authMiddelware(ctx *gin.Context){
  fmt.Println("authMiddelware is running......")
  session := sessions.Default(ctx)
  //session.Delete("info")
	//session.Delete("danger")
	//session.Delete("success")
	//session.Clear()
  cookie, _ := ctx.Cookie("gin_cookie")
  var customClime Clim
  defaultClime, _ := util.ParseJSONWebTokenClaims(cookie, &customClime)
  if customClime.Authorized{
    ctx.Next()
  }else {
    session.AddFlash("You must be logged", "info")
    session.Save()
    //fmt.Println("info session added",session.Flashes("info"))
    ctx.Redirect(http.StatusFound, "login")
  }
  fmt.Println(defaultClime)
  /*var customClime Clim
  defaultClime, err := util.ParseJSONWebTokenClaims(cookie, &customClime)
	session.Save()
  fmt.Println(err)
  fmt.Println(customClime)
  fmt.Println(defaultClime)*/
  //fmt.Printf("Cookie value: %s \n", cookie)
  fmt.Println(reflect.TypeOf(cookie))
  //ctx.Next()
}

func getUserCash(ctx *gin.Context){
  cookie, _ := ctx.Cookie("gin_cookie")
  var customClime Clim
  util.ParseJSONWebTokenClaims(cookie, &customClime)
  session := sessions.Default(ctx)
  session.Set("name","majid")
  successmsg := session.Flashes("success")
  dangermsg := session.Flashes("danger")
  infomsg := session.Flashes("info")
  session.Delete("success")
  session.Delete("danger")
  session.Delete("info")
  session.Save()
  //defaultClime, err := util.ParseJSONWebTokenClaims(cookie, &customClime)
  util.ParseJSONWebTokenClaims(cookie, &customClime)
  cashs, _ := controller.GetCash("user_user_name",customClime.Name)
  var amountIncome float64
  var amountSpend float64
  amountSpend = 0.0
  amountIncome = 0.0
  for _, record := range cashs {
    if record.TypeCash == "spend" {
      rmcomma, _ := util.RemoveComma(record.Amount)
      ras , _ := strconv.ParseFloat(rmcomma,4)
      amountSpend = amountSpend + float64(ras)
    }
    if record.TypeCash == "income" {
      rmcomma, _ := util.RemoveComma(record.Amount)
      rai, _ := strconv.ParseFloat(rmcomma,4)
      amountIncome = amountIncome + float64(rai)
      
    }
  }
    amountIncomeaddc, _, _ := util.FormatAmount(strconv.FormatFloat(amountIncome,'f',4,64))
    amountSpendaddc, _, _ := util.FormatAmount(strconv.FormatFloat(amountSpend,'f',4,64))
    walletf := amountIncome - amountSpend
    wallet, _, _ := util.FormatAmount(strconv.FormatFloat(walletf,'f',4,64))
  fmt.Println("cash registered for majid zare" , customClime.Name, cashs)
  ctx.HTML(http.StatusOK, "cash", gin.H{
    "cashs": cashs,
    "AmountSpend":amountSpendaddc,
    "AmountIncome":amountIncomeaddc,
    "Wallet":wallet,
    "DangerMsg":dangermsg,
    "InfoMsg":infomsg,
    "SuccessMsg":successmsg,
  })
  //fmt.Printf("Cookie value: %s \n", cookie)
  
}

func postUserCash(ctx *gin.Context){
  cookie, _ := ctx.Cookie("gin_cookie")
  amount := ctx.PostForm("amount")
	typeCash := ctx.PostForm("typecash")
  var customClime Clim
  //defaultClime, err := util.ParseJSONWebTokenClaims(cookie, &customClime)
  util.ParseJSONWebTokenClaims(cookie, &customClime)
  session := sessions.Default(ctx)
  session.Delete("success")
  session.Delete("danger")
  session.Delete("info")
  session.Save()
	if !util.IsIntFloat(amount){
	  session.AddFlash("Invalid value for Amount e.g 123456 or 1.25644", "danger")
    session.Save()
	  ctx.Redirect(http.StatusFound, "/usercash")
	}else{
	
	a ,_ ,_ := util.FormatAmount(amount)
  added, _ := controller.AddCash(customClime.Name, a , typeCash)
  if added == false {
    session.AddFlash(" Cash doesn't added unfortunately", "danger")
    fmt.Println("Cash does not added unfortunately")
  }
  
  session.AddFlash(" Cash added Successfully", "success")
  
  session.Save()
  ctx.Redirect(http.StatusFound, "/usercash")
  //fmt.Printf("Cookie value: %s \n", cookie)
	}
}
func deleteCash(ctx *gin.Context){
  id := ctx.Param("id")
  
  fmt.Println(id)
  controller.DeleteCash("id",id)
  ctx.Redirect(http.StatusFound,"/usercash")
}