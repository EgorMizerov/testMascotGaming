package main

import (
	"github.com/EgorMizerov/testMascotGaming/cmd"
)

type Reply struct {
	Data string
}

//map[string]interface{}{
//"jsonrpc": "2.0",
//"method": "Game.List",
//"id": "1920911592",
//}

var api = "https://api.c27staging.netgamecore.com/v1/"

func main() {
	cmd.Run()
}

//
//func main() {
//	client := NewJSONRPCClient()
//
//	http.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
//		//fmt.Fprintf(writer, "Hello!")
//		query := `{"jsonrpc":"2.0","method":"Game.List","id":"1920911592"}`
//
//		req, err := http.NewRequest("POST", api, strings.NewReader(query))
//		if err != nil {
//			fmt.Fprintf(writer, err.Error())
//		}
//
//		req.Header.Set("Content-Type", "application/json")
//		//req.Header.Set("C")
//		req.Header.Set("Accept", "application/json")
//
//		res, err := client.Do(req)
//		if err != nil {
//			fmt.Fprintf(writer, err.Error())
//		}
//
//		resp, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			fmt.Fprintf(writer, err.Error())
//		}
//
//		fmt.Fprintf(writer, string(resp))
//	})
//
//	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintf(writer, "Get")
//	})
//
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	go func() {
//		log.Fatalln(http.ListenAndServe(":8000", nil))
//		//log.Fatalln(http.ListenAndServeTLS(":443", "./example.com+5.pem", "./example.com+5-key.pem", nil))
//	}()
//
//	<- quit
//}
//
//
//func NewJSONRPCClient() *http.Client {
//	cert, err := tls.LoadX509KeyPair("./ssl/client.crt", "./ssl/client.key")
//	if err != nil {
//		log.Fatalf("server: loadkeys: %s", err)
//	}
//
//	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
//	tlsConfig.BuildNameToCertificate()
//	transport := &http.Transport{TLSClientConfig: tlsConfig}
//	client := &http.Client{Transport: transport}
//
//	return client
//}
