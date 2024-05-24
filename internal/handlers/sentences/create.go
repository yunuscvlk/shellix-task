package sentences

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yunuscvlk/shellix-task/internal/db"
	"github.com/yunuscvlk/shellix-task/internal/models/sentence"
	"github.com/yunuscvlk/shellix-task/internal/utils"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	var instance sentence.Model

	err := json.NewDecoder(r.Body).Decode(&instance)

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again"))

		return
	}

	db.Create(instance.Sentence)

	w.Write(utils.Result(true, "created"))
}