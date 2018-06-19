package route

import (
	"fmt"
	"net/http"
)

/**
 * route-web
 * @return {[type]} [description]
 */
func Route(rw, re) {
	http.HandleFunc("/", func() {
		fmt.Fprint(rw, "/") //客服端
	})

	http.HandleFunc("/index", func() {
		fmt.Fprint(rw, "/index") //客服端
	})
}
