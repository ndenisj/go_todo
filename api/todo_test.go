package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/ndenisj/go_todo/db/mock"
	db "github.com/ndenisj/go_todo/db/sqlc"
	"github.com/ndenisj/go_todo/utils"
	"github.com/stretchr/testify/require"
)

func TestGetTodoAPI(t *testing.T) {
	todo := randomTodo()

	testCases := []struct {
		name          string
		todoId        int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			todoId: todo.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Times(1).
					Return(todo, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTodo(t, recorder.Body, successResponse("successful", todo))
			},
		},
		{
			name:   "NotFound",
			todoId: todo.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Times(1).
					Return(db.Todo{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:   "InternalError",
			todoId: todo.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Times(1).
					Return(db.Todo{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:   "InvalidId",
			todoId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// create new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test HTTP server and send account request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/todos/%d", tc.todoId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func randomTodo() db.Todo {
	return db.Todo{
		ID:      utils.RandomInt(1, 1000),
		Owner:   utils.RandomOwner(),
		Title:   utils.RandomTitle(),
		Content: utils.RandomContent(),
	}
}

func requireBodyMatchTodo(t *testing.T, body *bytes.Buffer, todo interface{}) {

	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTodo db.Todo
	err = json.Unmarshal(data, &gotTodo)

	require.NoError(t, err)
	require.Equal(t, todo, successResponse("successful", gotTodo))
}
