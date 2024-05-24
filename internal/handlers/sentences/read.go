package sentences

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yunuscvlk/shellix-task/internal/db"
	"github.com/yunuscvlk/shellix-task/internal/utils"
)

func ReadHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	params := make(map[string]string)

	for k, v := range r.URL.Query() {
		params[k] = v[0]
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again (1)"))

		return
	}

	result := db.Read(id, params)

	data, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again (2)"))

		return
	}

	w.Write([]byte(data))
}

func ReadAllHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	params := make(map[string]string)

	for k, v := range r.URL.Query() {
		params[k] = v[0]
	}

	page, err := strconv.Atoi(params["page"])

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again (1)"))

		return
	}

	result := db.ReadAll(page)

	data, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again (2)"))

		return
	}

	w.Write([]byte(data))
}