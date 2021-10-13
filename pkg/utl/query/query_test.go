package query_test

import (
	"testing"

	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"

	"github.com/stretchr/testify/assert"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/query"
)

func TestList(t *testing.T) {
	type args struct {
		user PD_BookingBR_Evaluation.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *PD_BookingBR_Evaluation.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: PD_BookingBR_Evaluation.AuthUser{
				Role: PD_BookingBR_Evaluation.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: PD_BookingBR_Evaluation.AuthUser{
				Role:      PD_BookingBR_Evaluation.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &PD_BookingBR_Evaluation.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: PD_BookingBR_Evaluation.AuthUser{
				Role:       PD_BookingBR_Evaluation.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &PD_BookingBR_Evaluation.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: PD_BookingBR_Evaluation.AuthUser{
				Role: PD_BookingBR_Evaluation.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
