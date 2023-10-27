package types

// The SignUp struct is used to bind the request body form to the struct.
type SignUp struct {
	Username  string `json:"username" form:"username" binding:"required,min=4,max=32"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required,min=8,max=32"`
	FirstName string `json:"first_name" form:"first_name" binding:"required,min=4,max=32"`
	LastName  string `json:"last_name" form:"last_name" binding:"required,min=4,max=32"`
	Age       int    `json:"age" form:"age" binding:"required,min=1"`
}

// The Login struct is used to bind the request body form to the struct.
type Login struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
