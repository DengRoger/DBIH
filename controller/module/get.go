package module

import (
	// "fmt"
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "strconv"
	// "github.com/gorilla/mux"
)

type List struct {
	List     string `json:"List"`
	NextPage string `json:"NextPage"`
}

type Head struct {
	ListKey string `json:"ListKey"`
}

var Lists = []List{
	{List: "List1", NextPage: "NextPage1"},
	{List: "List2", NextPage: "NextPage2"},
	{List: "List3", NextPage: "NextPage3"},
}

var Heads = []Head{
	{ListKey: "ListKey1"},
	{ListKey: "ListKey2"},
	{ListKey: "ListKey3"},
}

func GetLists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Lists)
}

// get json from http://localhost:8443/page
func GetHead(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Heads)
}
