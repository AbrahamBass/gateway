package main

import (
	"log"
	"net/http"

	"package-gateway/handlers"
	"package-gateway/models"
	"package-gateway/utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	log.Println("Initializing routes...")

	viper.AddConfigPath("./config") //Viper looks here for the files.
	viper.SetConfigType("yaml")     //Sets the format of the config file.
	viper.SetConfigName("default")  // So that Viper loads default.yml.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Warning could not load configuration: %v", err)
	}
	viper.AutomaticEnv() // Merges any overrides set through env vars.

	gatewayConfig := &models.GatewayConfig{}

	err = viper.UnmarshalKey("gateway", gatewayConfig)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	for _, route := range gatewayConfig.Routes {
		proxy, err := utils.NewProxy(route.Target)
		if err != nil {
			panic(err)
		}

		handler := handlers.NewHandler(proxy)

		r.HandleFunc(route.Context+"/{targetPath:.*}", handler)
	}

	log.Fatal(http.ListenAndServe(gatewayConfig.ListenAddr, r))
}
