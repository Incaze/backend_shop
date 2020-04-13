package store

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/net/webdav"

	"shop/backend/utils"
)

type Controller struct {
	Repository Repository
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// always pass for now
		next(w, req)
	})
}

/* GET */
func (c *Controller) GetToken(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
	}
	_ = json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

/* GET */
func (c *Controller) Index(w http.ResponseWriter, _ *http.Request) {
	products := c.Repository.GetProducts()
	data, _ := json.Marshal(products)
	w = utils.SetRequestHeaders(w)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
	return
}

/* POST */
func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, utils.OneMb))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddProduct", err)
	}

	if err := json.Unmarshal(body, &product); err != nil {
		w.WriteHeader(webdav.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.AddProduct(product)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

/* PUT */
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, utils.OneMb))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &product); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.UpdateProduct(product) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w = utils.SetRequestHeaders(w)
	w.WriteHeader(http.StatusOK)
	return
}

/* GET */
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	productId, err := strconv.Atoi(id)

	if err != nil {
		log.Fatalln("Error GetProduct", err)
	}

	product := c.Repository.GetProductById(productId)
	data, _ := json.Marshal(product)

	w = utils.SetRequestHeaders(w)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
	}
	return
}

/* DELETE */
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)

	if err != nil {
		log.Fatalln("Error GetProduct", err)
	}

	if err := c.Repository.DeleteProduct(productId); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w = utils.SetRequestHeaders(w)
	w.WriteHeader(http.StatusOK)
	return
}
