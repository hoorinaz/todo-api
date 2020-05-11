package web

import (

	"github.com/hoorinaz/TodoList/pkg/user/processor"
	"github.com/hoorinaz/TodoList/pkg/user/store"

)

func RegisterUserService () UserWebService{

userStore := store.NewUserStore()
UserProcessor:= processor.NewUserProcessor(userStore)
webService:= NewUserWebService(UserProcessor)

	return webService
}





