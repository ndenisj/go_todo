package api

import (
	"database/sql"
	"testing"

	db "github.com/ndenisj/go_todo/db/sqlc"
	"github.com/ndenisj/go_todo/utils"
	"github.com/stretchr/testify/require"
)

// import (
// 	"database/sql"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/containerd/containerd/pkg/cri/store"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	mockdb "github.com/ndenisj/go_todo/db/mock"
// 	db "github.com/ndenisj/go_todo/db/sqlc"
// 	"github.com/ndenisj/go_todo/utils"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateUserApi(t *testing.T) {
// 	user, password := randomUser(t)

// 	testCases := []struct {
// 		name string
// 		body gin.H
// 		buildStubs func(store *mockdb.MockStore)
// 		checkResponse func(recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"full_name": user.FullName,
// 				"phone": user.Phone,
// 				"username": user.Username,
// 				"password": user.HashedPassword,
// 				"email": user.Email,
// 			},
// 			buildStubs: func(store *mockdb.MockStore){
// 				store.EXPECT().
// 				CreateUser(gomock.Any(), gomock.Any()).
// 				Times(1).
// 				Return(user, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {

// 			},
// 		}
// 	}
// }

func randomUser(t *testing.T) (user db.User, password string) {
	password = utils.RandomString(7)
	hashedPassword, err := utils.HashedPassword(password)
	require.NoError(t, err)

	user = db.User{
		ID:       utils.RandomInt(100, 1000),
		FullName: utils.RandomFullname(),
		Phone: sql.NullString{
			String: utils.RandomPhone(),
			Valid:  true,
		},
		Username:       utils.RandomUsername(),
		HashedPassword: hashedPassword,
		Email:          utils.RandomEmail(),
	}

	return
}
