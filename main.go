// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that takes a stream of bytes and looks for the bytes
// “elvis” and when they are found, replace them with “Elvis”. The code
// cannot assume that there are any line feeds or other delimiters in the
// stream and the code must assume that the stream is of any arbitrary length.
// The solution cannot meaningfully buffer to the end of the stream and
// then process the replacement.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)



func main() {
	mux:=new (http.ServeMux)
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_,err:=writer.Write([]byte("11111"));
		fmt.Println(err);
	})
	s := &http.Server{
		Addr:           ":3001",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}


//type MuxHandle struct {
//
//}
//
//
//func(this *MuxHandle) ServeHTTP(rep http.ResponseWriter, req *http.Request){
//	var err error
//	//var num int64
//	fmt.Println(req.URL)
//	_,err=rep.Write([]byte("11111"))
//	fmt.Println(err)
//	http.HandleFunc()
//}



