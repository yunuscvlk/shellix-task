package sentences

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yunuscvlk/shellix-task/internal/db"
	"github.com/yunuscvlk/shellix-task/internal/utils"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request, db *db.DB) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Println(err)
		
		w.Write(utils.Result(false, "failed, try again"))

		return
	}

	db.Delete(id)

	w.Write(utils.Result(true, "deleted"))
}