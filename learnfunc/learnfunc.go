package learnfunc

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func PrintHello(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello World!")

	if(err!=nil) {
		log.Println(err)
		return
	}

	fmt.Println("Bytes for printing is:", n)

}

func Addition(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, adding(2,2));
}

func adding (x,y int) string {
	return strconv.Itoa(x+y);
}

func Division(w http.ResponseWriter, r *http.Request) {
	res,err := dividing(2,0);
	if(err!=nil) {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("RESULT: %f", res))
}

func dividing (x,y float32) (float32, error) {
	if(y==0) {
		err := errors.New("Cannot Divide by 0");
		return 0,err;
	}

	result := x/y;
	return result, nil;
}

type jsonResponse struct {
	Ok			bool		`json:"ok"`
	Message	string	`json:"message"`
}

var res jsonResponse

func JsonExample (w http.ResponseWriter, r *http.Request) {
	//Unmarshal
	myJson := `
		[
			{
				"ok": true,
				"message": "First Data"
			},
			{
				"ok": false,
				"message": "Second Data"
			}
		]`

	var unmarshalled []jsonResponse

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {log.Println("Error unmarshalling json", err)}
	log.Printf("unmarshalled: %v", unmarshalled)


	//Marshall
	resp := jsonResponse {
		Ok: true,
		Message: "You have successfully written a json response!",
	}

	//res, err := json.Marshal(resp)
	res,err := json.MarshalIndent(resp, "", "  ")
	log.Println(string(res))

	if err!=nil {log.Fatal(err)}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	_,_ = w.Write(res)
}

func JsonGetExample(w http.ResponseWriter, r *http.Request) {
	res.Ok = true
	res.Message = "test"
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(res)
}

func JsonPostExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var unmarshalled jsonResponse
	err := json.NewDecoder(r.Body).Decode(&unmarshalled)
	if err!=nil {log.Println(err)}
	log.Println(unmarshalled)
	json.NewEncoder(w).Encode(unmarshalled)
}