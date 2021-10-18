package helpers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestGenerateRoute(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "FirstTest", want: "/report/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRoute(); got != tt.want {
				t.Errorf("GenerateRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateReportLink(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "FirstTest", want: "/report/test.xlsx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateReportLink("test.xlsx"); got != tt.want {
				t.Errorf("GenerateReportLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImplode(t *testing.T) {
	type args struct {
		glue string
		ci   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "FirstTest", args: struct {
			glue string
			ci   []interface{}
		}{glue: ", ", ci: []interface{}{"one", "two", "three"}}, want: "one, two, three"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Implode(tt.args.glue, tt.args.ci); got != tt.want {
				t.Errorf("Implode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceSlice(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{name: "firstTest", args: struct{ slice interface{} }{slice: []interface{}{"one", "two", "three"}}, want: []interface{}{"one", "two", "three"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToMap(t *testing.T) {
	type args struct {
		ci []interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "FirstTest", args: struct{ ci []interface{} }{ci: []interface{}{"one", "two", "three"}}, want: []string{"one", "two", "three"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToMap(tt.args.ci); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SendPostMessage(t *testing.T) {
	SendPostMessage("127.0.0.1", url.Values{"path": {"test send post messages"}})
}

func TestSendPostMessage(t *testing.T) {
	var (
		server *httptest.Server
	)
	volume := url.Values{"path": {"path test one"}}
	tests := map[string]struct {
		volumeName string
		err        error
	}{
		"TestOne":   {"127.0.0.1", nil},
		"TestTwo":   {"localhost", nil},
		"TestThree": {"ya.ru", nil},
		"TestFour":  {"routeam.ru", nil},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			response := `{"metadata":{"annotations":{"vsm.openebs.io/targetportals":"10.98.65.136:3260","vsm.openebs.io/cluster-i    ps":"10.98.65.136","openebs.io/jiva-iqn":"iqn.2016-09.com.openebs.jiva:vol","deployment.kubernetes.io/revision":"1","openebs.io/storage-pool"    :"default","vsm.openebs.io/replica-count":"1","openebs.io/jiva-controller-status":"Running","openebs.io/volume-monitor":"false","openebs.io/r    eplica-container-status":"Running","openebs.io/jiva-controller-cluster-ip":"10.98.65.136","openebs.io/jiva-replica-status":"Running","vsm.ope    nebs.io/iqn":"iqn.2016-09.com.openebs.jiva:vol","openebs.io/capacity":"2G","openebs.io/jiva-controller-ips":"10.36.0.6","openebs.io/jiva-repl    ica-ips":"10.36.0.7","vsm.openebs.io/replica-status":"Running","vsm.openebs.io/controller-status":"Running","openebs.io/controller-container-    status":"Running","vsm.openebs.io/replica-ips":"10.36.0.7","openebs.io/jiva-target-portal":"10.98.65.136:3260","openebs.io/volume-type":"jiva    ","openebs.io/jiva-replica-count":"1","vsm.openebs.io/volume-size":"2G","vsm.openebs.io/controller-ips":"10.36.0.6"},"creationTimestamp":null    ,"labels":{},"name":"vol"},"status":{"Message":"","Phase":"Running","Reason":""}}`
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, response)
			}))
			SendPostMessage(tt.volumeName, volume)
			defer server.Close()
		})
	}
}
