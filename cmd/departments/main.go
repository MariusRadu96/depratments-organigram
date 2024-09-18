package main

import (
	departmentservice "departments-organigram/internal/core/services/departmentsrv"
	"departments-organigram/internal/core/services/usersrv"
	httphandlers "departments-organigram/internal/handlers/http"
	"departments-organigram/internal/handlers/http/middlewares"
	"departments-organigram/internal/init/config"
	"departments-organigram/internal/init/db"
	"departments-organigram/internal/repositories/departments"
	"departments-organigram/internal/repositories/users"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conf := config.Get()

	dbConn := db.NewMysqlDB(&db.Config{
		Host:     conf.MySQLHost,
		Name:     conf.MySQLDB,
		Port:     conf.MySQLPort,
		User:     conf.MySQLUser,
		Password: conf.MySQLPassword,
	})

	departmentsStore := departments.NewStore(dbConn)
	usersStore := users.NewStore(dbConn)

	departmentsSrv := departmentservice.NewDepartmentSrv(departmentsStore)
	usersSrv := usersrv.NewUsersSrv(usersStore, conf)

	httpHandler := httphandlers.NewHTTPHandler(departmentsSrv, usersSrv)

	http.HandleFunc("/register", httpHandler.Register)
	http.HandleFunc("/login", httpHandler.Login)

	http.HandleFunc("/create-department", middlewares.AuthMiddlewawre(httpHandler.CreateDepartment))
	http.HandleFunc("/update-department", middlewares.AuthMiddlewawre(httpHandler.UpdateDepartment))
	http.HandleFunc("/delete-department/{id}", middlewares.AuthMiddlewawre(httpHandler.DeleteDepartment))
	http.HandleFunc("/departments", middlewares.AuthMiddlewawre(httpHandler.GetAllDepartments))
	http.HandleFunc("/department", middlewares.AuthMiddlewawre(httpHandler.GetDepartment))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}
