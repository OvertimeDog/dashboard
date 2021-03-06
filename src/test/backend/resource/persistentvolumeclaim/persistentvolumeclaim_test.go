package persistentvolumeclaim


import (
"reflect"
"testing"

"github.com/kubernetes/dashboard/src/app/backend/resource/common"
"k8s.io/kubernetes/pkg/api"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
)

func TestGetPersistentVolumeClaimList(t *testing.T) {
	cases := []struct {
		persistentVolumeClaims []api.PersistentVolumeClaim
		expected          	*PersistentVolumeClaimList
	}{
		{nil, &PersistentVolumeClaimList{Items: []PersistentVolumeClaim{}}},
		{
			[]api.PersistentVolumeClaim{{
				ObjectMeta: api.ObjectMeta{Name: "foo"},
				Spec: api.PersistentVolumeClaimSpec{VolumeName: "my-volume"},
				Status: api.PersistentVolumeClaimStatus{ Phase: api.ClaimBound},
				},
			},
			&PersistentVolumeClaimList{
				ListMeta: common.ListMeta{TotalItems: 1},
				Items: []PersistentVolumeClaim{{
					TypeMeta:   common.TypeMeta{Kind: "persistentvolumeclaim"},
					ObjectMeta: common.ObjectMeta{Name: "foo"},
					Status: "Bound",
					Volume: "my-volume",
				}},
			},
		},
	}
	for _, c := range cases {
		actual := getPersistentVolumeClaimList(c.persistentVolumeClaims, dataselect.NoDataSelect)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("getPersistentVolumeClaimList(%#v) == \n%#v\nexpected \n%#v\n",
				c.persistentVolumeClaims, actual, c.expected)
		}
	}
}


