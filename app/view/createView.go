package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.blackeye.id/Aldi.Rismawan/centrotil/db/umongo"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	PIPELINE_JSON_DIR = "./app/view/"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := gotenv.Load(); err != nil {
		log.Println(err)
	}
}

func main() {
	//* -------------------------- CREATE VIEW FROM JSON ------------------------- */
	listPipeline, err := os.ReadDir(PIPELINE_JSON_DIR)
	if err != nil {
		log.Println(err)
		return
	}

	for _, pipeline := range listPipeline {
		if !strings.Contains(pipeline.Name(), ".json") {
			fmt.Println("Skip", pipeline.Name())
			continue
		}

		listPart := []string{}
		collectionName := ""
		fileWithoutExtension := strings.TrimSuffix(filepath.Base(pipeline.Name()), filepath.Ext(pipeline.Name()))
		viewName := strings.ReplaceAll(fileWithoutExtension, "v_v_", "v_")
		for i, part := range strings.Split(fileWithoutExtension, "_") {
			takePartCount := 2
			if strings.Contains(fileWithoutExtension, "v_v_") {
				takePartCount = 3
			}
			if i < takePartCount {
				listPart = append(listPart, part)
			}
		}
		collectionName = strings.Replace(strings.Join(listPart, "_"), "v_", "", 1)

		pipelineRaw, err := os.ReadFile(PIPELINE_JSON_DIR + pipeline.Name())
		if err != nil {
			log.Println(err)
			return
		}

		pipeline := []bson.M{}
		if err := json.Unmarshal(pipelineRaw, &pipeline); err != nil {
			fmt.Printf("pipelineRaw: %s\n", string(pipelineRaw))
			log.Println(err)
			return
		}

		_, _ = collectionName, viewName
		// client := db.NewMongoDbUtilLocal(collectionName)
		client := umongo.NewMongoDbUtilUseEnv(collectionName)
		if err := client.CreateViewIfNotExists(viewName, pipeline); err != nil {
			return
		}
	}
}
