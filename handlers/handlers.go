package handlers

import (
	"bitbucket.org/goura/spotmc-ctrl/template-data"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/autoscaling"

	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type AppConfig struct {
	AutoScalingGroupName string
	AWSRegion            string
	BasicUsername        string
	BasicPassword        string
}

func NewAppConfig() *AppConfig {
	config := &AppConfig{
		AutoScalingGroupName: os.Getenv("GFMC_AUTO_SCALING_GROUP"),
		AWSRegion:            os.Getenv("GFMC_AWS_REGION"),
		BasicUsername:        os.Getenv("GFMC_BASIC_USERNAME"),
		BasicPassword:        os.Getenv("GFMC_BASIC_PASSWORD"),
	}
	return config
}

type ServerStatus struct {
	Up      bool
	Started bool
}

func autoScalingClient() *autoscaling.AutoScaling {
	ac := NewAppConfig()

	creds := aws.DetectCreds("", "", "")
	region := ac.AWSRegion

	asCli := autoscaling.New(creds, region, nil)

	return asCli
}

func GetStatus(GroupName string) (*ServerStatus, error) {
	ac := NewAppConfig()

	asCli := autoScalingClient()
	req := &autoscaling.AutoScalingGroupNamesType{
		AutoScalingGroupNames: []string{ac.AutoScalingGroupName},
	}
	res, err := asCli.DescribeAutoScalingGroups(req)
	if err != nil {
		return nil, err
	}

	n := len(res.AutoScalingGroups)
	if n != 1 {
		return nil, fmt.Errorf("len(AutoScalingGroups) was not 1, it was %d", n)
	}

	asg := res.AutoScalingGroups[0]

	ServerIsUp := len(asg.Instances) == 1
	ServerStarted := *asg.DesiredCapacity == 1

	return &ServerStatus{
		Up:      ServerIsUp,
		Started: ServerStarted,
	}, nil
}

func setDesiredCapacity(name string, capacity int) error {
	asCli := autoScalingClient()
	req := &autoscaling.SetDesiredCapacityType{
		AutoScalingGroupName: aws.String(name),
		DesiredCapacity:      aws.Integer(capacity),
	}
	err := asCli.SetDesiredCapacity(req)
	return err
}

func Stop(GroupName string) error {
	ac := NewAppConfig()
	return setDesiredCapacity(ac.AutoScalingGroupName, 0)
}

func Start(GroupName string) error {
	ac := NewAppConfig()
	return setDesiredCapacity(ac.AutoScalingGroupName, 1)
}

func Render(w http.ResponseWriter, Filename string, v interface{}) {
	b, err := data.Asset(Filename)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	t := template.Must(template.New(Filename).Parse(string(b)))
	t.Execute(w, v)
}

func TopPageHandler(w http.ResponseWriter, r *http.Request) {
	ac := NewAppConfig()
	status, err := GetStatus(ac.AutoScalingGroupName)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
	}

	Render(w, "top.html", status)
}

func StartPageHandler(w http.ResponseWriter, r *http.Request) {
	ac := NewAppConfig()
	status, err := GetStatus(ac.AutoScalingGroupName)
	if err != nil {
		w.WriteHeader(500)
	}

	Start(ac.AutoScalingGroupName)

	Render(w, "start.html", status)
}

func StopPageHandler(w http.ResponseWriter, r *http.Request) {
	ac := NewAppConfig()
	status, err := GetStatus(ac.AutoScalingGroupName)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
	}

	Stop(ac.AutoScalingGroupName)

	Render(w, "stop.html", status)
}

func AuthWrapper(handler http.HandlerFunc) (wrapped http.HandlerFunc) {
	ac := NewAppConfig()
	wrapped = func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			log.Print("BasicAuth returned false")
			hdr := w.Header()
			hdr.Set("WWW-Authenticate", `Basic realm="Login Required"`)
			w.WriteHeader(401)
			return
		}
		if username == ac.BasicUsername && password == ac.BasicPassword {
			log.Print("password matched")
			handler(w, r)
			return
		}
		log.Print("password didn't match", ac.BasicUsername)
		w.WriteHeader(403)
	}
	return
}
