package cfapi

import (
	"encoding/json"
	"fmt"
	"github.com/bmerchant22/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CodeforcesClient struct {
	client http.Client
}

func (cfClient *CodeforcesClient) RecentActions(maxCount int) ([]models.RecentAction, error) {
	api := "https://codeforces.com/api/recentActions?maxCount=" + strconv.Itoa(maxCount)

	resp, err := cfClient.client.Get(api)
	if err != nil {
		log.Printf("Error occured while calling cf api: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error occurred while reading the resp body")
		return nil, err
	}

	wrapper := struct {
		Status string
		Result []models.RecentAction
	}{}

	if err = json.Unmarshal(data, &wrapper); err != nil {
		log.Printf("Error while unmarshalling data from cfapi : %v", err)
	}

	//fmt.Println(string(data))

	return wrapper.Result, err
}
