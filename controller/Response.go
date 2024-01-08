package controller

//type 

func Response(w *http.ResponseWriter,messageType,message string,code int8){
	w.WriteHeader(code)
	   w.Header().Set("Content-Type", "application/json")
	   resp := make(map[string]string)
	   resp["message"] = message
	   resp["token"]=message
	   jsonResp, err := json.Marshal(resp)
	   if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	   w.Write(jsonResp)
}