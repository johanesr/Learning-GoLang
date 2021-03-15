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

func JsonExample (w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse {
		Ok: true,
		Message: "You have successfully written a json response!",
	}

	//res, err := json.Marshal(resp)
	res,err := json.MarshalIndent(resp, "", "  ")
	log.Println(res)

	if err!=nil {log.Fatal(err)}

	w.Header().Set("Content-Type", "application/json")
	log.Fatal(w.Write(res))
}
