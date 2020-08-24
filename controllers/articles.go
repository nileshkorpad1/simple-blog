package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/nileshkorpad1/simple-blog/config"
	"github.com/nileshkorpad1/simple-blog/models"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/matryer/respond.v1"
)

// connect db
var collection = config.ConnectDB()

// CreateArticle godoc
// @Summary Create a new Article
// @Description Create a new Article with the input paylod
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body models.Article true "Create Article"
// @Success 200 {object} models.Article
// @Router /api/v1/articles [post]
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
func CreateArticle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var article models.Article

	// we decode our body request params
	json.NewDecoder(request.Body).Decode(&article)

	//validation for empty fields

	validate := validator.New()

	err := validate.Struct(article)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//cancel to prevent memory leakage
	defer cancel()

	// insert our book model.
	result, err := collection.InsertOne(ctx, article)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	data := map[string]interface{}{"data": result, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)

}

//GetArticle : Get article by id

// GetArticle godoc
// @Summary Get details for a given articleId
// @Description Get details of article corresponding to the input articleId
// @Tags articles
// @Accept  json
// @Produce  json
// @Param articleId path int true "ID of the article"
// @Success 200 {object} models.Article
// @Router /api/v1/articles/{id} [get]
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
func GetArticle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	//request params
	params := mux.Vars(request)
	//convert id to usable mongodb object id
	id, errorID := primitive.ObjectIDFromHex(params["id"])
	if errorID != nil {
		data := map[string]interface{}{"data": nil, "message": errorID.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	var article models.Article

	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//cancel to prevent memory leakage
	defer cancel()

	//query the model
	err := collection.FindOne(ctx, models.Article{ID: id}).Decode(&article)

	//handle error
	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	// handle success data
	data := map[string]interface{}{"data": article, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}

//GetArticles : Get all articles

// GetArticles godoc
// @Summary Get details of all Articles
// @Description Get details of all articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Article
// @Router /api/v1/articles [get]
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
func GetArticles(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var articles models.Articles

	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//cancel to prevent memory leakage
	defer cancel()

	//query data
	cursor, err := collection.Find(ctx, bson.M{})

	//handle error
	if err != nil {

		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	defer cursor.Close(ctx)

	//loop through person
	for cursor.Next(ctx) {
		var article models.Article
		cursor.Decode(&article)
		articles = append(articles, article)
	}

	//handle error
	if err := cursor.Err(); err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	//handle success
	data := map[string]interface{}{"data": articles, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}
