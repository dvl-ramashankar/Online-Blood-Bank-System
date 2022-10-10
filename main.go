package main

import (
	"bloodSystem/auth"
	"bloodSystem/entity"
	"bloodSystem/service"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var ser = service.Connection{}
var finalResponse entity.Response

func init() {
	ser.Server = "mongodb://localhost:27017"
	ser.Database = "BloodSystem"
	ser.Collection1 = "UserDetails"
	ser.Collection2 = "DonorDetails"
	ser.Collection3 = "BloodDetails"
	ser.Collection4 = "PatientDetails"
	ser.Collection5 = "LoginDetails"

	ser.Connect()
}

func saveUserDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody entity.UserDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if dataBody.MailId == "" || dataBody.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Please enter mailId or Password")
		return
	}

	if result, err := ser.SaveUserDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchUsersDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.SearchUsersDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func updateUserDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody entity.UserDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateUserDetailsById(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deleteUserDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeleteUserDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func saveDonorDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	var dataBody entity.DonorDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if dataBody.MailId == "" || dataBody.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Please enter mailId or Password")
		return
	}

	if result, err := ser.SaveDonorDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchDonorDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.SearchDonorDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func updateDonorDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody entity.DonorDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateDonorDetailsById(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deleteDonorDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeleteDonorDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func applyBloodPatientDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	var dataBody entity.PatientDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if dataBody.MailId == "" || dataBody.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Please enter mailId or Password")
		return
	}

	if result, err := ser.ApplyBloodPatientDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func givenBloodPatientDetailsById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.GivenBloodPatientDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchAllPendingBloodPatientDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	if result, err := ser.SearchAllPendingBloodPatientDetails(); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deletePendingBloodPatientDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	token := r.Header.Get("tokenid")
	err := validateToken(token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeletePendingBloodPatientDetails(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchFilterBloodDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var bloodDetailsRequest entity.BloodDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&bloodDetailsRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if result, err := ser.SearchFilterBloodDetails(bloodDetailsRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func generateToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var loginDetails entity.LoginDetails
	if err := json.NewDecoder(r.Body).Decode(&loginDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if result, err := ser.GenerateToken(loginDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg}, "error")
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}, err string) {
	if err == "error" {
		finalResponse.Success = "false"
	} else {
		finalResponse.Success = "true"
	}
	finalResponse.SucessCode = fmt.Sprintf("%v", code)
	finalResponse.Response = payload
	response, _ := json.Marshal(finalResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func validateToken(token string) error {
	if token == "" {
		return errors.New("Please Enter Token")
	}
	err := auth.ValidateToken(token)
	if err != nil {
		return errors.New("Either Token Is Invalid Or Expired")
	}
	return err
}

func main() {
	http.HandleFunc("/generate-token", generateToken)
	http.HandleFunc("/save-user-details", saveUserDetails)
	http.HandleFunc("/search-user-details-id/", searchUsersDetailsById)
	http.HandleFunc("/update-user-details-id/", updateUserDetailsById)
	http.HandleFunc("/delete-user-details-id/", deleteUserDetailsById)
	http.HandleFunc("/save-donor-details", saveDonorDetails)
	http.HandleFunc("/search-donor-details-id/", searchDonorDetailsById)
	http.HandleFunc("/update-donor-details-id/", updateDonorDetailsById)
	http.HandleFunc("/delete-donor-details-id/", deleteDonorDetailsById)
	http.HandleFunc("/apply-blood-patient-details", applyBloodPatientDetails)
	http.HandleFunc("/given-blood-patient-details-id/", givenBloodPatientDetailsById)
	http.HandleFunc("/search-all-pending-patient-details/", searchAllPendingBloodPatientDetails)
	http.HandleFunc("/delete-pending-patient-request/", deletePendingBloodPatientDetails)
	http.HandleFunc("/search-filter-blood-details/", searchFilterBloodDetails)
	log.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}
