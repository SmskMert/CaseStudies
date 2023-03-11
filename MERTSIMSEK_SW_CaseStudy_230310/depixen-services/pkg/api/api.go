package api

import (
   
    "encoding/json"
    "depixen-services/pkg/db/models"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "github.com/go-pg/pg/v10"
)

//start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB) *chi.Mux {
    //get the router
    r := chi.NewRouter()
    //add middleware
    //store DB to use it later
    r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

    r.Route("/cards", func(r chi.Router) {
        r.Post("/", createCard)
        r.Get("/", getCards)
    })

    r.Route("/card", func(r chi.Router) {
        r.Get("/", getTheLastCard)
        r.Get("/{id}", getCardById)
    })
    
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("up and running"))
    })

    corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	}).Handler

	// Wrap router with the CORS handler
	http.ListenAndServe(":8080", corsHandler(r))
    return r
}

// -- Responses

type CreateCardRequest struct {
    Title string `json:"title"`
	Description string `json:"description"`
	Imageuri string `json:"imageuri"`
	Createddate string `json:"createddate"`
}

type CardResponse struct {
    Success bool            `json:"success"`
    Error   string          `json:"error"`
    Card *models.Card `json:"card"`
}

type CardsResponse struct {
	Success  bool              `json:"success"`
	Error    string            `json:"error"`
	Cards []*models.Card `json:"cards"`
}

func handleErr(w http.ResponseWriter, err error) {
    res := &CardResponse{
        Success: false,
        Error:   err.Error(),
        Card: nil,
    }
    err = json.NewEncoder(w).Encode(res)
    //if there's an error with encoding handle it
    if err != nil {
        log.Printf("error sending response %v\n", err)
    }
    //return a bad request and exist the function
    w.WriteHeader(http.StatusBadRequest)
}

func handleDBFromContextErr(w http.ResponseWriter) {
    res := &CardResponse{
        Success: false,
        Error:   "could not get the DB from context",
        Card: nil,
    }
    err := json.NewEncoder(w).Encode(res)
    //if there's an error with encoding handle it
    if err != nil {
        log.Printf("error sending response %v\n", err)
    }
    //return a bad request and exist the function
    w.WriteHeader(http.StatusBadRequest)
}

func getTheLastCard(w http.ResponseWriter, r *http.Request) {
    //get db from ctx
    pgdb, ok := r.Context().Value("DB").(*pg.DB)
    if !ok {
    handleDBFromContextErr(w)
        return
    }
    //call models package to access the database and return the last card
    card, err := models.GetTheLastCard(pgdb)
    if err != nil {
        handleErr(w, err)
        return
    }
    //positive response
    res := &CardResponse{
        Success:  true,
        Error:    "",
        Card: card,
    }
    //encode the positive response to json and send it back
    err = json.NewEncoder(w).Encode(res)
    if err != nil {
        log.Printf("error encoding card: %v\n", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func getCardById(w http.ResponseWriter, r *http.Request){
    //get the id from the URL parameter
    cardId := chi.URLParam(r, "id")

    //get db from ctx
    pgdb, ok := r.Context().Value("DB").(*pg.DB)
    if !ok {
        handleDBFromContextErr(w)
        return
    }
    //get the card from the DB
    card, err := models.GetCardById(pgdb,cardId)
    if err != nil {
        handleErr(w,err)
        return
    }
    //positive response
    res := &CardResponse{
        Success: true,
        Error: "",
        Card: card,
    }
    //encode the positive response to json and send it back
    err = json.NewEncoder(w).Encode(res)
    if err != nil {
        log.Printf("error encoding card: %v\n", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
}


func createCard(w http.ResponseWriter, r *http.Request) {
    //get the request body and decode it
    req := &CreateCardRequest{}
    err := json.NewDecoder(r.Body).Decode(req)
    //if there's an error with decoding the information
    //send a response with an error
    if err != nil {
       handleErr(w, err)
       return
    }
    //get the db from context
    pgdb, ok := r.Context().Value("DB").(*pg.DB)
    //if no connection to db, handle the error
    //and send an adequate response
    if !ok {
        handleDBFromContextErr(w)
        return
    }
    //if can get the db then
    card, err := models.CreateCard(pgdb, &models.Card{
        Title: req.Title,
        Description:  req.Description,
        Imageuri:  req.Imageuri,
        Createddate: req.Createddate ,
    })
    if err != nil {
        handleErr(w, err)
        return
    }
    //everything is good
    //return a positive response
    res := &CardResponse{
        Success: true,
        Error:   "",
        Card: card,
    }
    err = json.NewEncoder(w).Encode(res)
    if err != nil {
        log.Printf("error encoding after creating card %v\n", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func getCards(w http.ResponseWriter, r *http.Request) {
    //get db from ctx
    pgdb, ok := r.Context().Value("DB").(*pg.DB)
    if !ok {
       handleDBFromContextErr(w)
        return
    }
    //call models package to access the database and return the cards
    cards, err := models.GetCards(pgdb)
    if err != nil {
       handleErr(w, err)
        return
    }
    //positive response
    res := &CardsResponse{
        Success:  true,
        Error:    "",
        Cards: cards,
    }
    //encode the positive response to json and send it back
    err = json.NewEncoder(w).Encode(res)
    if err != nil {
        log.Printf("error encoding comments: %v\n", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
}