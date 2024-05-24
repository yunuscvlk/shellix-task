package sentences

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yunuscvlk/shellix-task/internal/db"
	"github.com/yunuscvlk/shellix-task/internal/models/sentence"
	"github.com/yunuscvlk/shellix-task/internal/utils"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	var instance sentence.Model

	err := json.NewDecoder(r.Body).Decode(&instance)

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again (1)"))

		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again (2)"))

		return
	}

	db.Update(id, instance.Sentence)
	
	w.Write(utils.Result(true, "updated"))
}