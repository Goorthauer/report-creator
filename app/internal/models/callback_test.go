package models

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
	var (
		server *httptest.Server
	)
	params := map[string]interface{}{"callback": Callback{
		Start:  "127.0.0.1/start",
		Finish: "127.0.0.1/finish",
		Status: "127.0.0.1/status",
		Failed: "127.0.0.1/failed",
		Body:   "",
	}}
	callback := NewCallback(params)
	tests := map[string]struct {
		volumeName string
		err        error
	}{
		"TestOne":   {"5", fmt.Errorf("this is err ")},
		"TestTwo":   {"15", fmt.Errorf("this is err ")},
		"TestThree": {"25", fmt.Errorf("this is err ")},
		"TestFour":  {"35", fmt.Errorf("this is err ")},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			response := `{"metadata":{"annotations":{"vsm.openebs.io/targetportals":"10.98.65.136:3260","vsm.openebs.io/cluster-i    ps":"10.98.65.136","openebs.io/jiva-iqn":"iqn.2016-09.com.openebs.jiva:vol","deployment.kubernetes.io/revision":"1","openebs.io/storage-pool"    :"default","vsm.openebs.io/replica-count":"1","openebs.io/jiva-controller-status":"Running","openebs.io/volume-monitor":"false","openebs.io/r    eplica-container-status":"Running","openebs.io/jiva-controller-cluster-ip":"10.98.65.136","openebs.io/jiva-replica-status":"Running","vsm.ope    nebs.io/iqn":"iqn.2016-09.com.openebs.jiva:vol","openebs.io/capacity":"2G","openebs.io/jiva-controller-ips":"10.36.0.6","openebs.io/jiva-repl    ica-ips":"10.36.0.7","vsm.openebs.io/replica-status":"Running","vsm.openebs.io/controller-status":"Running","openebs.io/controller-container-    status":"Running","vsm.openebs.io/replica-ips":"10.36.0.7","openebs.io/jiva-target-portal":"10.98.65.136:3260","openebs.io/volume-type":"jiva    ","openebs.io/jiva-replica-count":"1","vsm.openebs.io/volume-size":"2G","vsm.openebs.io/controller-ips":"10.36.0.6"},"creationTimestamp":null    ,"labels":{},"name":"vol"},"status":{"Message":"","Phase":"Running","Reason":""}}`
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, response)
			}))
			callback.SendErr(tt.err)
			callback.SetBody(name)
			callback.SendFinish()
			callback.SendStart()
			callback.SendStatus(tt.volumeName)
			defer server.Close()
		})
	}
}
