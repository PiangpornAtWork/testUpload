package main


import (

   "net/http"


   "github.com/gin-gonic/gin"


   "example.com/test/config"

   "example.com/test/controller"

   "example.com/test/controller/users"

   "example.com/test/middlewares"

)


const PORT = "8000"


func main() {


   // open connection database

   config.ConnectionDB()


   // Generate databases

   config.SetupDatabase()


   r := gin.Default()


   r.Use(CORSMiddleware())


   // Auth Route

   r.POST("/signup", users.SignUp)

   r.POST("/signin", users.SignIn)
   r.GET("/users", users.GetAll)

   router := r.Group("/")

   {

       router.Use(middlewares.Authorizes())


       // User Route

       router.PUT("/user/:id", users.Update)

       

       router.GET("/user/:id", users.Get)

       router.DELETE("/user/:id", users.Delete)


   }


   r.GET("/genders", controller.GetAll)


   r.GET("/", func(c *gin.Context) {

       c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)

   })


   // Run the server


   r.Run("localhost:" + PORT)


}


func CORSMiddleware() gin.HandlerFunc {

   return func(c *gin.Context) {

       c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

       c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

       c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

       c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")


       if c.Request.Method == "OPTIONS" {

           c.AbortWithStatus(204)

           return

       }


       c.Next()

   }

}