package service

import (
	"bloodSystem/entity"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Server      string
	Database    string
	Collection1 string
	Collection2 string
	Collection3 string
	Collection4 string
}

var CollectionUserDetails *mongo.Collection
var CollectionDonorDetails *mongo.Collection
var CollectionBloodDetails *mongo.Collection
var CollectionPatientDetails *mongo.Collection
var ctx = context.TODO()

func (e *Connection) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	CollectionUserDetails = client.Database(e.Database).Collection(e.Collection1)
	CollectionDonorDetails = client.Database(e.Database).Collection(e.Collection2)
	CollectionBloodDetails = client.Database(e.Database).Collection(e.Collection3)
	CollectionPatientDetails = client.Database(e.Database).Collection(e.Collection4)
}

// ===================================userDetails============================================
func (e *Connection) SaveUserDetails(reqBody entity.UserDetailsRequest) (string, error) {
	fmt.Println("save Method:", reqBody)
	saveData, err := SetValueInUserModel(reqBody)
	if err != nil {
		log.Println(err)
		return "", err
	}
	data, err := CollectionUserDetails.InsertOne(ctx, saveData)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to store data")
	}
	fmt.Println(data)
	return "User Saved Successfully", nil
}

func (e *Connection) SearchUsersDetailsById(idStr string) ([]*entity.UserDetails, error) {
	var finalData []*entity.UserDetails

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return finalData, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", id}},
				bson.D{{"active", true}},
			},
		},
	}
	data, err := CollectionUserDetails.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	finalData, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	return finalData, nil
}

func (e *Connection) UpdateUserDetailsById(reqData entity.UserDetailsRequest, idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", id}},
				bson.D{{"active", true}},
			},
		},
	}
	UpdateQuery := bson.D{}
	if reqData.Name != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
	}
	if reqData.Age != 0 {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "age", Value: reqData.Age})
	}
	if reqData.BloodGroup != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "blood_group", Value: reqData.BloodGroup})
	}
	if reqData.AdharCard != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "adhar_card", Value: reqData.AdharCard})
	}
	if reqData.DOB != "" {
		dob, err := convertDate(reqData.DOB)
		if err != nil {
			log.Println(err)
			return updatedDocument, err
		}
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "dob", Value: dob})
	}
	update := bson.D{{"$set", UpdateQuery}}

	r := CollectionUserDetails.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *Connection) DeleteUserDetailsById(idStr string) (string, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{"$set", bson.D{primitive.E{Key: "active", Value: false}}}}
	CollectionUserDetails.FindOneAndUpdate(ctx, filter, update)
	return "Documents Deactivated Successfully", err
}

func convertDbResultIntoUserStruct(fetchDataCursor *mongo.Cursor) ([]*entity.UserDetails, error) {
	var finaldata []*entity.UserDetails
	for fetchDataCursor.Next(ctx) {
		var data entity.UserDetails
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func convertDate(dateStr string) (time.Time, error) {

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Println(err)
		return date, err
	}
	return date, nil
}

func SetValueInUserModel(req entity.UserDetailsRequest) (entity.UserDetails, error) {
	var data entity.UserDetails
	dob, err := convertDate(req.DOB)
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.DOB = dob
	data.Name = req.Name
	data.Age = req.Age
	data.AdharCard = req.AdharCard
	data.BloodGroup = req.BloodGroup
	data.Active = true
	data.Location = req.Location
	data.CreatedDate = time.Now()
	return data, nil
}

// XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
// ==========================================Donor detail======================================
func (e *Connection) SaveDonorDetails(reqBody entity.DonorDetailsRequest) (string, error) {
	saveData, err := SetValueInModel(reqBody)
	if err != nil {
		return "", errors.New("Unable to parse date")
	}
	data, err := CollectionDonorDetails.InsertOne(ctx, saveData)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to store data")
	}
	fmt.Println(data)
	str, err := saveBloodQuantityInBloodDetails(reqBody)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println(str)
	return "Donor Details Saved Successfully", nil
}

func (e *Connection) SearchDonorDetailsById(idStr string) ([]*entity.DonorDetails, error) {
	var finalData []*entity.DonorDetails

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return finalData, err
	}

	data, err := CollectionDonorDetails.Find(ctx, bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	finalData, err = convertDbResultIntoDonorStruct(data)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	return finalData, nil
}

func (e *Connection) UpdateDonorDetailsById(reqData entity.DonorDetailsRequest, idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", id}},
				bson.D{{"active", true}},
			},
		},
	}
	UpdateQuery := bson.D{}
	if reqData.Name != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
	}
	if reqData.Age != 0 {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "age", Value: reqData.Age})
	}
	if reqData.BloodGroup != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "blood_group", Value: reqData.BloodGroup})
	}
	if reqData.AdharCard != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "adhar_card", Value: reqData.AdharCard})
	}
	if reqData.Location != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "location", Value: reqData.Location})
	}
	if reqData.DOB != "" {
		dob, err := convertDate(reqData.DOB)
		if err != nil {
			log.Println(err)
			return updatedDocument, err
		}
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "dob", Value: dob})
	}
	if reqData.DepositDate != "" {
		dd, err := convertDate(reqData.DepositDate)
		if err != nil {
			log.Println(err)
			return updatedDocument, err
		}
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "deposit_date", Value: dd})
	}
	update := bson.D{{"$set", UpdateQuery}}
	r := CollectionDonorDetails.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}

	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *Connection) DeleteDonorDetailsById(idStr string) (string, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{"$set", bson.D{primitive.E{Key: "active", Value: false}}}}
	CollectionDonorDetails.FindOneAndUpdate(ctx, filter, update)
	return "Documents Deactivated Successfully", err
}

func convertDbResultIntoDonorStruct(fetchDataCursor *mongo.Cursor) ([]*entity.DonorDetails, error) {
	var finaldata []*entity.DonorDetails
	for fetchDataCursor.Next(ctx) {
		var data entity.DonorDetails
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func SetValueInModel(req entity.DonorDetailsRequest) (entity.DonorDetails, error) {
	var data entity.DonorDetails
	depositDate, err := convertDate(req.DepositDate)
	if err != nil {
		log.Println(err)
		return data, err
	}
	dob, err := convertDate(req.DOB)
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.DepositDate = depositDate
	data.DOB = dob
	data.Units = req.Units
	data.Name = req.Name
	data.Age = req.Age
	data.AdharCard = req.AdharCard
	data.BloodGroup = req.BloodGroup
	data.Active = true
	data.Location = req.Location
	return data, nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//=======================================Blood Details==========================================

func saveBloodQuantityInBloodDetails(reqBody entity.DonorDetailsRequest) (string, error) {
	var finalData []*entity.BloodDetails
	depositDate, err := convertDate(reqBody.DepositDate)
	if err != nil {
		return "", err
	}
	unitInt, err := convertUnitsStringIntoInt(reqBody.Units)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{primitive.E{Key: "location", Value: reqBody.Location}},
				bson.D{primitive.E{Key: "deposit_date", Value: depositDate}},
			},
		},
	}
	data, err := CollectionBloodDetails.Find(ctx, filter)
	finalData, err = convertDbResultIntoBloodStruct(data)
	if err != nil {
		return "", nil
	}
	if finalData == nil {
		saved, err := createNewEntryIntoBloodDetails(reqBody, unitInt, depositDate)
		if err != nil {
			return "", err
		}
		fmt.Println(saved)
	} else {
		unitDB := finalData[0].Units
		addUnit := unitDB + unitInt
		fmt.Println("Total Units:", addUnit)
		CollectionBloodDetails.FindOneAndUpdate(ctx, filter, bson.D{{"$set", bson.D{{"units", addUnit}}}})
	}
	return "Blood Details Saved Successfully", nil
}

func createNewEntryIntoBloodDetails(reqBody entity.DonorDetailsRequest, unitInt int, depositDate time.Time) (string, error) {
	var bloodDetails entity.BloodDetails

	bloodDetails.Units = unitInt
	bloodDetails.Location = reqBody.Location
	bloodDetails.BloodGroup = reqBody.BloodGroup
	bloodDetails.DepositDate = depositDate
	bloodDetails.CreatedDate = time.Now()
	_, err := CollectionBloodDetails.InsertOne(ctx, bloodDetails)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	return "New entry created in blood details", nil
}

func convertDbResultIntoBloodStruct(fetchDataCursor *mongo.Cursor) ([]*entity.BloodDetails, error) {
	var finaldata []*entity.BloodDetails
	for fetchDataCursor.Next(ctx) {
		var data entity.BloodDetails
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func convertUnitsStringIntoInt(units string) (int, error) {
	unitReplace := strings.ReplaceAll(units, "ml", "")
	unitInt, err := strconv.Atoi(unitReplace)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	return unitInt, nil
}

func deductBloodUnitsFromBloodDetails(units string, location string, bloodDateStr string) (string, error) {
	unitInt, err := convertUnitsStringIntoInt(units)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	bloodDate, err := convertDate(bloodDateStr)
	if err != nil {
		return "", err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"location", location}},
				bson.D{{"deposit_date", bloodDate}},
			},
		},
	}
	fmt.Println(filter)
	data, err := CollectionBloodDetails.Find(ctx, filter)
	finalData, err := convertDbResultIntoBloodStruct(data)
	fmt.Println(finalData)
	if err != nil {
		return "", nil
	}
	if finalData == nil {
		return "", errors.New("Data not present in Blood details according to given location and desposited date")
	}
	unit := finalData[0].Units
	if !(unit >= unitInt) {
		return "", errors.New("Insufficient Blood!")
	}
	addUnit := unit - unitInt
	fmt.Println("Total Units:", addUnit)
	CollectionBloodDetails.FindOneAndUpdate(ctx, filter, bson.D{{"$set", bson.D{{"units", addUnit}}}})
	return "Deduct Successfully", nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

//==========================================Patient Details=====================================

func (e *Connection) ApplyBloodPatientDetails(reqBody entity.PatientDetailsRequest) (string, error) {

	saveData, err := SetValueInPatientModel(reqBody)
	if err != nil {
		log.Println(err)
		return "", err
	}
	deduct, err := deductBloodUnitsFromBloodDetails(reqBody.ApplyUnits, reqBody.Location, reqBody.BloodDate)
	if err != nil {
		return "", err
	}
	fmt.Println(deduct)
	data, err := CollectionPatientDetails.InsertOne(ctx, saveData)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to store data")
	}
	fmt.Println(data)
	return "User Saved Successfully", nil
}

func (e *Connection) SearchAllPendingBloodPatientDetails() ([]*entity.PatientDetails, error) {
	var finalData []*entity.PatientDetails

	data, err := CollectionPatientDetails.Find(ctx, bson.D{primitive.E{Key: "active", Value: true}})
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	finalData, err = convertDbResultIntoPatientStruct(data)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	return finalData, nil
}

func (e *Connection) GivenBloodPatientDetailsById(idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", id}},
				bson.D{{"active", true}},
			},
		},
	}

	UpdateQuery := bson.D{}
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "active", Value: false})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "given_date", Value: time.Now()})

	update := bson.D{{"$set", UpdateQuery}}

	r := CollectionPatientDetails.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *Connection) DeletePendingBloodPatientDetails(idStr string) (string, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{"$set", bson.D{primitive.E{Key: "active", Value: false}}}}
	CollectionPatientDetails.FindOneAndUpdate(ctx, filter, update)
	return "Documents Deactivated Successfully", err
}

func SetValueInPatientModel(req entity.PatientDetailsRequest) (entity.PatientDetails, error) {
	var data entity.PatientDetails
	dob, err := convertDate(req.DOB)
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.DOB = dob
	data.Name = req.Name
	data.Age = req.Age
	data.AdharCard = req.AdharCard
	data.BloodGroup = req.BloodGroup
	data.Active = true
	data.Location = req.Location
	data.CreatedDate = time.Now()
	data.ApplyUnits = req.ApplyUnits
	data.ApplyDate = time.Now()
	return data, nil
}

func convertDbResultIntoPatientStruct(fetchDataCursor *mongo.Cursor) ([]*entity.PatientDetails, error) {
	var finaldata []*entity.PatientDetails
	for fetchDataCursor.Next(ctx) {
		var data entity.PatientDetails
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX