package handlers

import (
	"net/http"
	"time"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/dto"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/entity"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/08-apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
	"github.com/goccy/go-json"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(UserDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: UserDB,
	}
}

// CreateUser godoc
// @Summary      Create user
// @Description  Creates a new user with email, name and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserInput  true  "User request"
// @Success      201
// @Failure      500  {object}  Error
// @Router			 /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetJWT godoc
// @Summary      Get JWT
// @Description  Get JWT token for user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.GetJWTInput  true  "User credentials"
// @Success      200  {object}  dto.GetJwtOutput
// @Failure      404 	{object}  Error
// @Failure      500  {object}  Error
// @Router			 /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJwtOutput{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)

}
