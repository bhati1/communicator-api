package controller

import (
	"comm-api/connector"
	"comm-api/models"
	"comm-api/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// func SendMessage(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Inside GetMessage")
// 	w.Header().Set("Content-Type", "application/json")
// 	// loc, _ := time.LoadLocation("Asia/Calcutta")
// 	message := models.Message{
// 		MessageBody: "hi tushar",
// 		PhoneNumber: "8383891601",
// 		TimeStamp:   time.Now().String(),
// 	}

// 	json.NewEncoder(w).Encode(message)

// }

func SendMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside Send Message")

	w.Header().Set("Content-Type", "application/json")

	var msg models.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {

		json.NewEncoder(w).Encode("Incorrect Input")
		return

	}

	if msg.MessageBody == "" {
		json.NewEncoder(w).Encode("Message given by u is empty")
		return
	}

	// loc, _ := time.LoadLocation("Asia/Calcutta")
	msg.TimeStamp = time.Now().String()

	if msg.PhoneNumber == "" {
		json.NewEncoder(w).Encode("Phone Number is not given by you")
		return
	}

	if err := connector.InsertOne(msg); err != nil {
		json.NewEncoder(w).Encode("Internal server Error")

		return
	}

	fmt.Println("Message received is -", msg.MessageBody)
	json.NewEncoder(w).Encode("Message added successsfully")

}

func GetMessageById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside function get msg by id")

	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	var msg models.Message

	msgId := params["id"]

	if msgId == "" {

		json.NewEncoder(w).Encode("Pls provide a msg id..")

	}

	msg, err := connector.GetOne(msgId)

	if err != nil {
		json.NewEncoder(w).Encode("Invalid id")
		// return
	}

	json.NewEncoder(w).Encode(msg)
}

func GetAllMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside Function Get all messages")

	w.Header().Set("content-type", "application/json")

	var messages []models.Message

	messages, err := connector.GetAll()

	if err != nil {
		json.NewEncoder(w).Encode("Not able to fetch all messages..")
	}

	json.NewEncoder(w).Encode(messages)

}

func DeleteOneMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside function delete one message")

	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	msgId := params["id"]

	if msgId == "" {

		json.NewEncoder(w).Encode("Pls provide a msg id..")

	}

	err := connector.DeleteOne(msgId)

	if err != nil {
		json.NewEncoder(w).Encode("Invalid id")
		// return
	}

	json.NewEncoder(w).Encode("Successfully deleted message by id " + msgId)

}

func DeleteAllMessages(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside function delete All message")

	w.Header().Set("content-type", "application/json")

	err := connector.DeleteAll()

	if err != nil {
		json.NewEncoder(w).Encode("Not able to delete all the messages")
	}

	json.NewEncoder(w).Encode("Successfully deleted all messages")

}

func UpdateMessageById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside function update msg by id")

	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	msgId := params["id"]

	msgbody := params["msg"]

	if msgId == "" {

		json.NewEncoder(w).Encode("Pls provide a msg id..")
		return

	}

	if msgbody == "" {

		json.NewEncoder(w).Encode("Pls provide msg to update")
		return

	}

	err := connector.UpdateOne(msgId, msgbody)

	if err != nil {
		json.NewEncoder(w).Encode("Problem occurred while updating a msg_id : " + msgId)
		return
	}

	json.NewEncoder(w).Encode("Successfully update msg_id : " + msgId + "with new message : " + msgbody)

}

func DownloadMsgAsFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside Send File function")

	w.Header().Set("content-type", "application/json")

	// to set file name
	file_name := "allmessages" + time.Now().String()
	attachment := fmt.Sprintf("attachment; filename=%s.json", file_name)

	w.Header().Set("Content-Disposition", attachment)

	content, err := service.Downloadallmsg()
	if err != nil {
		json.NewEncoder(w).Encode("Failed to download all messages")
	}

	http.ServeContent(w, r, file_name, time.Now(), strings.NewReader(content))

}

func DownloadVideo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside Download video function")

	w.Header().Set("content-type", "video/mp4")

	http.ServeFile(w, r, "/Users/ankitbhati/Downloads/video_to_test.mp4")
}
