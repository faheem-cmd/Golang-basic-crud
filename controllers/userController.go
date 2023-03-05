package controllers

import (
	"net/http"
	"os"
	"time"

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//Get email from eq body
	var body struct{
		Email string
		Password string 
	}
	if c.Bind(&body) !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to request",
		})
		return
	}
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10,)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to hash password",
		})
		return
	}
	 
	//Create the user
	 user := models.User{Email: body.Email,Password: string(hash)}
	 result := initializers.DB.Create(&user)

	 if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to create user",
		})
		return
	 }

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"status":201,
		"data":"User creates successfully",

	})
}


func Login(c *gin.Context) {
	//Get email from eq body
	var body struct{
		Email string
		Password string 
	}
	if c.Bind(&body) !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to request",
		})
		return
	}
	
	var user models.User
	initializers.DB.First(&user, "email =?",body.Email)

	if user.ID == 0{
			c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid creds",
		})
		return 
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid creds",
		})
		return 
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userID":user.ID,
		"expire":time.Now().Add(time.Hour * 24 * 30).Unix(), 
	})

	tokenString ,err:= token.SignedString([]byte(os.Getenv("SECRET")))
	num:= "1233"
	tokendata := map[string]string{"access_token": tokenString,"number":num }
	  details := []gin.H{
            {"name": "joe", "age": 28},
            {"name": "philip", "age": 18},
        }
		  metaData := []string{"1", "faheem", "hello"}
		  number := []int{1,2,3,4}


	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid creds",
		})
		return 
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(" ",tokenString,3600 * 24 * 30,"","",false,true )
	
	c.JSON(http.StatusOK, gin.H{
		"status":200,
		"token":tokendata,
		"data":details,
		"meta":metaData,
		"number":number,
	})

}

func Valid(c *gin.Context) {
	 userID := c.MustGet("userID")
	 var data models.User
	 initializers.DB.Find(&data,userID)
		 c.JSON(http.StatusOK, gin.H{"userID": data})
}