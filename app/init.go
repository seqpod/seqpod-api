package app

import (
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/seqpod-api/controllers/v0"
	"github.com/otiai10/seqpod-api/filters"
)

func init() {

	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	// defer session.Close()

	mf := filters.InitMongoFilter(session)
	lf := filters.InitLogFilter()
	af := filters.InitializeAuthFilter()
	cf := new(marmoset.ContextFilter)

	unauthorized := marmoset.NewRouter()
	unauthorized.GET("/v0/status", v0.Status)

	authorized := marmoset.NewRouter()
	authorized.GET("/v0/jobs/(?P<id>[0-9a-f]+)", v0.JobGet)
	authorized.POST("/v0/jobs/(?P<id>[0-9a-f]+)/fastq", v0.JobFastqUpload)
	authorized.POST("/v0/jobs/(?P<id>[0-9a-f]+)/ready", v0.JobMarkReady)
	authorized.POST("/v0/jobs/workspace", v0.JobWorkspace)
	authorized.Apply(cf, af, mf)

	root := marmoset.NewRouter()
	root.Apply(lf)
	root.Subrouter(unauthorized)
	root.Subrouter(authorized)
	http.Handle("/", root)
}
