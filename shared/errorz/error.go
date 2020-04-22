package errorz

import (
	"fmt"
	"net/http"
)

func WriteHttpError(w http.ResponseWriter, httpStatus int, message ...interface{}) {
	fmt.Fprint(w, message...)
	w.WriteHeader(httpStatus)
}
