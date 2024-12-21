package utils

import (
	"net/http"
	"strconv"
)

func GetPaginator(r *http.Request) (int, int, int) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil || limit < 1 {
		limit = 10
	}

	return page, (page - 1) * limit, limit

}
