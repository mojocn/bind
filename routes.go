package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
)
var done chan(struct{})

func getIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"v": version,
	})
}
func postIndex(c *gin.Context) {
	log.Info("Is this exist?")
	if isJob {
		jobs := "one"
		log.WithFields(log.Fields{
			"job": "mp42dream",
		}).Info("added new job to queue")
		c.HTML(200, "jobs.html", gin.H{
			jobs: jobs,
		})
	} else {
		dream(c)
		c.HTML(200, "jobs.html", gin.H{})
	}
}
func getAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"variableName": "value",
	})
}

func getContact(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"variableName": "value",
	})
}

func getDownloads(c *gin.Context) {
	open.Run(basePath)
	open.Run(basePath + "/videos")
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
func getJobs(c *gin.Context) {
	c.HTML(200, "jobs.html", gin.H{
		"jobs": "the job you just made is awesome!",
		"est": "should be done in ten zetaseconds",
	})
}
func getCode(c *gin.Context) {
	c.HTML(200, "code.html", gin.H{
		"variableName": "value",
	})
}
func getDonate(c *gin.Context) {
	c.HTML(200, "donate.html", gin.H{
		"variableName": "value",
	})
}
