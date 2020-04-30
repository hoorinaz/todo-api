package errorz

import (
	"fmt"
	"net/http"
)

func WriteHttpError(w http.ResponseWriter, httpStatus int, message ...interface{}) {
	w.WriteHeader(httpStatus)
	fmt.Fprint(w, message...)

}
