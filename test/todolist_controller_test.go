package test

//
//import (
//	"PZN_RESTfulAPI/app"
//	"PZN_RESTfulAPI/controller"
//	"PZN_RESTfulAPI/helper"
//	"PZN_RESTfulAPI/middleware"
//	"PZN_RESTfulAPI/repository"
//	"PZN_RESTfulAPI/service"
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"github.com/go-playground/assert/v2"
//	"github.com/go-playground/validator/v10"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"strconv"
//	"strings"
//	"testing"
//	"time"
//)
//
//func setupTestDB() *sql.DB {
//	db, err := sql.Open("mysql", "eep:1903@tcp(localhost:3306)/G0_RESTfulAPI_test")
//	helper.PanicIfError(err)
//
//	db.SetMaxIdleConns(5)
//	db.SetMaxOpenConns(20)
//	db.SetConnMaxLifetime(60 * time.Minute)
//	db.SetConnMaxIdleTime(10 * time.Minute)
//
//	return db
//}
//
//func setupRouter(db *sql.DB) http.Handler {
//	validate := validator.New()
//	todolistRepository := repository.NewTodolistRepository()
//	todolistService := service.NewTodolistService(todolistRepository, db, validate)
//	todolistRepositoryController := controller.NewTodolistController(todolistService)
//	router := app.NewRouter(todolistRepositoryController)
//
//	return middleware.NewAuthMiddleware(router)
//}
//
//func truncateCategory(db *sql.DB) {
//	db.Exec("TRUNCATE category")
//}
//
//func TestCreateCategorySuccess(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//	router := setupRouter(db)
//
//	requestBody := strings.NewReader(`{"title" : "MENGELEARN HAYABUSA", "description" : "MENGELEARN IN YT"}`)
//	request := httptest.NewRequest(http.MethodPost, "http://localhost:1234/api/categories", requestBody)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 200, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 200, int(responseBody["status"].(float64)))
//	assert.Equal(t, "OK", responseBody["message"])
//	assert.Equal(t, "MENGELEARN HAYABUSA", "MENGELEARN IN YT", responseBody["data"].(map[string]interface{})["title", "description"])
//}
//
//func TestCreateCategoryFailed(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//	router := setupRouter(db)
//
//	requestBody := strings.NewReader(`{"name" : ""}`)
//	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 400, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 400, int(responseBody["code"].(float64)))
//	assert.Equal(t, "BAD REQUEST", responseBody["status"])
//}
//
//func TestUpdateCategorySuccess(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//
//	tx, _ := db.Begin()
//	categoryRepository := repository.NewCategoryRepository()
//	category := categoryRepository.Save(context.Background(), tx, domain.Category{
//		Name: "Gadget",
//	})
//	tx.Commit()
//
//	router := setupRouter(db)
//
//	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
//	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 200, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 200, int(responseBody["code"].(float64)))
//	assert.Equal(t, "OK", responseBody["status"])
//	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
//	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
//}
//
//func TestUpdateCategoryFailed(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//
//	tx, _ := db.Begin()
//	categoryRepository := repository.NewCategoryRepository()
//	category := categoryRepository.Save(context.Background(), tx, domain.Category{
//		Name: "Gadget",
//	})
//	tx.Commit()
//
//	router := setupRouter(db)
//
//	requestBody := strings.NewReader(`{"name" : ""}`)
//	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 400, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 400, int(responseBody["code"].(float64)))
//	assert.Equal(t, "BAD REQUEST", responseBody["status"])
//}
//
//func TestGetCategorySuccess(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//
//	tx, _ := db.Begin()
//	categoryRepository := repository.NewCategoryRepository()
//	category := categoryRepository.Save(context.Background(), tx, domain.Category{
//		Name: "Gadget",
//	})
//	tx.Commit()
//
//	router := setupRouter(db)
//
//	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 200, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 200, int(responseBody["code"].(float64)))
//	assert.Equal(t, "OK", responseBody["status"])
//	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
//	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
//}
//
//func TestGetCategoryFailed(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//	router := setupRouter(db)
//
//	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 404, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 404, int(responseBody["code"].(float64)))
//	assert.Equal(t, "NOT FOUND", responseBody["status"])
//}
//
//func TestDeleteCategorySuccess(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//
//	tx, _ := db.Begin()
//	categoryRepository := repository.NewCategoryRepository()
//	category := categoryRepository.Save(context.Background(), tx, domain.Category{
//		Name: "Gadget",
//	})
//	tx.Commit()
//
//	router := setupRouter(db)
//
//	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 200, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 200, int(responseBody["code"].(float64)))
//	assert.Equal(t, "OK", responseBody["status"])
//}
//
//func TestDeleteCategoryFailed(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//	router := setupRouter(db)
//
//	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
//	request.Header.Add("Content-Type", "application/json")
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 404, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 404, int(responseBody["code"].(float64)))
//	assert.Equal(t, "NOT FOUND", responseBody["status"])
//}
//
//func TestListCategoriesSuccess(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//
//	tx, _ := db.Begin()
//	categoryRepository := repository.NewCategoryRepository()
//	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
//		Name: "Gadget",
//	})
//	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
//		Name: "Computer",
//	})
//	tx.Commit()
//
//	router := setupRouter(db)
//
//	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
//	request.Header.Add("X-API-Key", "RAHASIA")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 200, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 200, int(responseBody["code"].(float64)))
//	assert.Equal(t, "OK", responseBody["status"])
//
//	fmt.Println(responseBody)
//
//	var categories = responseBody["data"].([]interface{})
//
//	categoryResponse1 := categories[0].(map[string]interface{})
//	categoryResponse2 := categories[1].(map[string]interface{})
//
//	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
//	assert.Equal(t, category1.Name, categoryResponse1["name"])
//
//	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
//	assert.Equal(t, category2.Name, categoryResponse2["name"])
//}
//
//func TestUnauthorized(t *testing.T) {
//	db := setupTestDB()
//	truncateCategory(db)
//	router := setupRouter(db)
//
//	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
//	request.Header.Add("X-API-Key", "SALAH")
//
//	recorder := httptest.NewRecorder()
//
//	router.ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	assert.Equal(t, 401, response.StatusCode)
//
//	body, _ := io.ReadAll(response.Body)
//	var responseBody map[string]interface{}
//	json.Unmarshal(body, &responseBody)
//
//	assert.Equal(t, 401, int(responseBody["code"].(float64)))
//	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
//}
//
