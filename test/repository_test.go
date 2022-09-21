package test

//// db connection
//func dbMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//
//	return db, mock
//}
//
//// Successfull Case
//func TestRepository_InsertUser(t *testing.T) {
//	sqlDB, mock := dbMock(t)
//	defer sqlDB.Close()
//
//	// request
//	req := request.InsertUsers{
//		FirstName: "derry",
//		LastName:  "renaldy",
//		Password:  "password",
//		Email:     "email@gmail.com",
//	}
//
//	// Expected sql.DB Begin to be called
//	prep := mock.ExpectPrepare("INSERT INTO")
//	prep.ExpectExec().WithArgs("derry", "renaldy", "password", "email@gmail.com").WillReturnResult(sqlmock.NewResult(1, 1))
//
//	// Execute Method
//	res, err := sqlDB.
//}
