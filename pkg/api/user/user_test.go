package user_test

import (
	"testing"

	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/user"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/mock"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/mock/mockdb"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type args struct {
		c   echo.Context
		req PD_BookingBR_Evaluation.User
	}
	cases := []struct {
		name     string
		args     args
		wantErr  bool
		wantData PD_BookingBR_Evaluation.User
		udb      *mockdb.User
		rbac     *mock.RBAC
		sec      *mock.Secure
	}{{
		name: "Fail on is lower role",
		rbac: &mock.RBAC{
			AccountCreateFn: func(echo.Context, PD_BookingBR_Evaluation.AccessRole, int, int) error {
				return PD_BookingBR_Evaluation.ErrGeneric
			}},
		wantErr: true,
		args: args{req: PD_BookingBR_Evaluation.User{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "JohnDoe",
			RoleID:    1,
			Password:  "Thranduil8822",
		}},
	},
		{
			name: "Success",
			args: args{req: PD_BookingBR_Evaluation.User{
				FirstName: "John",
				LastName:  "Doe",
				Username:  "JohnDoe",
				RoleID:    1,
				Password:  "Thranduil8822",
			}},
			udb: &mockdb.User{
				CreateFn: func(db orm.DB, u PD_BookingBR_Evaluation.User) (PD_BookingBR_Evaluation.User, error) {
					u.CreatedAt = mock.TestTime(2000)
					u.UpdatedAt = mock.TestTime(2000)
					u.Base.ID = 1
					return u, nil
				},
			},
			rbac: &mock.RBAC{
				AccountCreateFn: func(echo.Context, PD_BookingBR_Evaluation.AccessRole, int, int) error {
					return nil
				}},
			sec: &mock.Secure{
				HashFn: func(string) string {
					return "h4$h3d"
				},
			},
			wantData: PD_BookingBR_Evaluation.User{
				Base: PD_BookingBR_Evaluation.Base{
					ID:        1,
					CreatedAt: mock.TestTime(2000),
					UpdatedAt: mock.TestTime(2000),
				},
				FirstName: "John",
				LastName:  "Doe",
				Username:  "JohnDoe",
				RoleID:    1,
				Password:  "h4$h3d",
			}}}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, tt.sec)
			usr, err := s.Create(tt.args.c, tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantData, usr)
		})
	}
}

func TestView(t *testing.T) {
	type args struct {
		c  echo.Context
		id int
	}
	cases := []struct {
		name     string
		args     args
		wantData PD_BookingBR_Evaluation.User
		wantErr  error
		udb      *mockdb.User
		rbac     *mock.RBAC
	}{
		{
			name: "Fail on RBAC",
			args: args{id: 5},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return PD_BookingBR_Evaluation.ErrGeneric
				}},
			wantErr: PD_BookingBR_Evaluation.ErrGeneric,
		},
		{
			name: "Success",
			args: args{id: 1},
			wantData: PD_BookingBR_Evaluation.User{
				Base: PD_BookingBR_Evaluation.Base{
					ID:        1,
					CreatedAt: mock.TestTime(2000),
					UpdatedAt: mock.TestTime(2000),
				},
				FirstName: "John",
				LastName:  "Doe",
				Username:  "JohnDoe",
			},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
					if id == 1 {
						return PD_BookingBR_Evaluation.User{
							Base: PD_BookingBR_Evaluation.Base{
								ID:        1,
								CreatedAt: mock.TestTime(2000),
								UpdatedAt: mock.TestTime(2000),
							},
							FirstName: "John",
							LastName:  "Doe",
							Username:  "JohnDoe",
						}, nil
					}
					return PD_BookingBR_Evaluation.User{}, nil
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			usr, err := s.View(tt.args.c, tt.args.id)
			assert.Equal(t, tt.wantData, usr)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		c   echo.Context
		pgn PD_BookingBR_Evaluation.Pagination
	}
	cases := []struct {
		name     string
		args     args
		wantData []PD_BookingBR_Evaluation.User
		wantErr  bool
		udb      *mockdb.User
		rbac     *mock.RBAC
	}{
		{
			name: "Fail on query List",
			args: args{c: nil, pgn: PD_BookingBR_Evaluation.Pagination{
				Limit:  100,
				Offset: 200,
			}},
			wantErr: true,
			rbac: &mock.RBAC{
				UserFn: func(c echo.Context) PD_BookingBR_Evaluation.AuthUser {
					return PD_BookingBR_Evaluation.AuthUser{
						ID:         1,
						CompanyID:  2,
						LocationID: 3,
						Role:       PD_BookingBR_Evaluation.UserRole,
					}
				}}},
		{
			name: "Success",
			args: args{c: nil, pgn: PD_BookingBR_Evaluation.Pagination{
				Limit:  100,
				Offset: 200,
			}},
			rbac: &mock.RBAC{
				UserFn: func(c echo.Context) PD_BookingBR_Evaluation.AuthUser {
					return PD_BookingBR_Evaluation.AuthUser{
						ID:         1,
						CompanyID:  2,
						LocationID: 3,
						Role:       PD_BookingBR_Evaluation.AdminRole,
					}
				}},
			udb: &mockdb.User{
				ListFn: func(orm.DB, *PD_BookingBR_Evaluation.ListQuery, PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.User, error) {
					return []PD_BookingBR_Evaluation.User{
						{
							Base: PD_BookingBR_Evaluation.Base{
								ID:        1,
								CreatedAt: mock.TestTime(1999),
								UpdatedAt: mock.TestTime(2000),
							},
							FirstName: "John",
							LastName:  "Doe",
							Email:     "johndoe@gmail.com",
							Username:  "johndoe",
						},
						{
							Base: PD_BookingBR_Evaluation.Base{
								ID:        2,
								CreatedAt: mock.TestTime(2001),
								UpdatedAt: mock.TestTime(2002),
							},
							FirstName: "Hunter",
							LastName:  "Logan",
							Email:     "logan@aol.com",
							Username:  "hunterlogan",
						},
					}, nil
				}},
			wantData: []PD_BookingBR_Evaluation.User{
				{
					Base: PD_BookingBR_Evaluation.Base{
						ID:        1,
						CreatedAt: mock.TestTime(1999),
						UpdatedAt: mock.TestTime(2000),
					},
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@gmail.com",
					Username:  "johndoe",
				},
				{
					Base: PD_BookingBR_Evaluation.Base{
						ID:        2,
						CreatedAt: mock.TestTime(2001),
						UpdatedAt: mock.TestTime(2002),
					},
					FirstName: "Hunter",
					LastName:  "Logan",
					Email:     "logan@aol.com",
					Username:  "hunterlogan",
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			usrs, err := s.List(tt.args.c, tt.args.pgn)
			assert.Equal(t, tt.wantData, usrs)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}

}

func TestDelete(t *testing.T) {
	type args struct {
		c  echo.Context
		id int
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		udb     *mockdb.User
		rbac    *mock.RBAC
	}{
		{
			name:    "Fail on ViewUser",
			args:    args{id: 1},
			wantErr: PD_BookingBR_Evaluation.ErrGeneric,
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
					if id != 1 {
						return PD_BookingBR_Evaluation.User{}, nil
					}
					return PD_BookingBR_Evaluation.User{}, PD_BookingBR_Evaluation.ErrGeneric
				},
			},
		},
		{
			name: "Fail on RBAC",
			args: args{id: 1},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
					return PD_BookingBR_Evaluation.User{
						Base: PD_BookingBR_Evaluation.Base{
							ID:        id,
							CreatedAt: mock.TestTime(1999),
							UpdatedAt: mock.TestTime(2000),
						},
						FirstName: "John",
						LastName:  "Doe",
						Role: &PD_BookingBR_Evaluation.Role{
							AccessLevel: PD_BookingBR_Evaluation.UserRole,
						},
					}, nil
				},
			},
			rbac: &mock.RBAC{
				IsLowerRoleFn: func(echo.Context, PD_BookingBR_Evaluation.AccessRole) error {
					return PD_BookingBR_Evaluation.ErrGeneric
				}},
			wantErr: PD_BookingBR_Evaluation.ErrGeneric,
		},
		{
			name: "Success",
			args: args{id: 1},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
					return PD_BookingBR_Evaluation.User{
						Base: PD_BookingBR_Evaluation.Base{
							ID:        id,
							CreatedAt: mock.TestTime(1999),
							UpdatedAt: mock.TestTime(2000),
						},
						FirstName: "John",
						LastName:  "Doe",
						Role: &PD_BookingBR_Evaluation.Role{
							AccessLevel: PD_BookingBR_Evaluation.AdminRole,
							ID:          2,
							Name:        "Admin",
						},
					}, nil
				},
				DeleteFn: func(db orm.DB, usr PD_BookingBR_Evaluation.User) error {
					return nil
				},
			},
			rbac: &mock.RBAC{
				IsLowerRoleFn: func(echo.Context, PD_BookingBR_Evaluation.AccessRole) error {
					return nil
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			err := s.Delete(tt.args.c, tt.args.id)
			if err != tt.wantErr {
				t.Errorf("Expected error %v, received %v", tt.wantErr, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		c   echo.Context
		upd user.Update
	}
	cases := []struct {
		name     string
		args     args
		wantData PD_BookingBR_Evaluation.User
		wantErr  error
		udb      *mockdb.User
		rbac     *mock.RBAC
	}{
		{
			name: "Fail on RBAC",
			args: args{upd: user.Update{
				ID: 1,
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return PD_BookingBR_Evaluation.ErrGeneric
				}},
			wantErr: PD_BookingBR_Evaluation.ErrGeneric,
		},
		{
			name: "Fail on Update",
			args: args{upd: user.Update{
				ID: 1,
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			wantErr: PD_BookingBR_Evaluation.ErrGeneric,
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
					return PD_BookingBR_Evaluation.User{
						Base: PD_BookingBR_Evaluation.Base{
							ID:        1,
							CreatedAt: mock.TestTime(1990),
							UpdatedAt: mock.TestTime(1991),
						},
						CompanyID:  1,
						LocationID: 2,
						RoleID:     3,
						FirstName:  "John",
						LastName:   "Doe",
						Mobile:     "123456",
						Phone:      "234567",
						Address:    "Work Address",
						Email:      "golang@go.org",
					}, nil
				},
				UpdateFn: func(db orm.DB, usr PD_BookingBR_Evaluation.User) error {
					return PD_BookingBR_Evaluation.ErrGeneric
				},
			},
		},
		{
			name: "Success",
			args: args{upd: user.Update{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Mobile:    "123456",
				Phone:     "234567",
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			wantData: PD_BookingBR_Evaluation.User{
				Base: PD_BookingBR_Evaluation.Base{
					ID:        1,
					CreatedAt: mock.TestTime(1990),
					UpdatedAt: mock.TestTime(2000),
				},
				CompanyID:  1,
				LocationID: 2,
				RoleID:     3,
				FirstName:  "John",
				LastName:   "Doe",
				Mobile:     "123456",
				Phone:      "234567",
				Address:    "Work Address",
				Email:      "golang@go.org",
			},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
					return PD_BookingBR_Evaluation.User{
						Base: PD_BookingBR_Evaluation.Base{
							ID:        1,
							CreatedAt: mock.TestTime(1990),
							UpdatedAt: mock.TestTime(2000),
						},
						CompanyID:  1,
						LocationID: 2,
						RoleID:     3,
						FirstName:  "John",
						LastName:   "Doe",
						Mobile:     "123456",
						Phone:      "234567",
						Address:    "Work Address",
						Email:      "golang@go.org",
					}, nil
				},
				UpdateFn: func(db orm.DB, usr PD_BookingBR_Evaluation.User) error {
					usr.UpdatedAt = mock.TestTime(2000)
					return nil
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			usr, err := s.Update(tt.args.c, tt.args.upd)
			assert.Equal(t, tt.wantData, usr)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestInitialize(t *testing.T) {
	u := user.Initialize(nil, nil, nil)
	if u == nil {
		t.Error("User service not initialized")
	}
}
