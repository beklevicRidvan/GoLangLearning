package users

import (
	"errors"
	"net/mail"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	testManager := NewManager()

	testFirstName := "Test"
	testLastName := "Userman"
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Fatalf("error parsing test email address: %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}
	if len(testManager.users) != 1 {
		t.Errorf("Bad test manager user count, wanted %d, got: %d", 1, len(testManager.users))

		if len(testManager.users) < 1 {
			t.Fatal()
		}
	}

	expectedUser := User{
		FirstName: testFirstName,
		LastName:  testLastName,
		Email:     *testEmail,
	}

	foundUser := testManager.users[0]

	if !reflect.DeepEqual(expectedUser, foundUser) {
		t.Errorf("added user data is not correct\nwanted: %+v\n got: %+v\n", expectedUser, foundUser)
	}

}

func TestAddUserInvalidEmail(t *testing.T) {
	testManager := NewManager()

	testFirstName := "Test"
	testLastName := "UserMan"
	testEmail := "foobar"

	err := testManager.AddUser(testFirstName, testLastName, testEmail)
	if err == nil {
		t.Error("Noo error returned for invalid email")
	} else {
		exceptedErr := "invalid email: foobar"
		if err.Error() != exceptedErr {
			t.Errorf("bad error text,wanted: %s, got %s", exceptedErr, err)
		}
	}
	if len(testManager.users) > 0 {
		t.Errorf("bad test manager user count, wanted %d, got %d", 0, len(testManager.users))
	}

}

func TestAddUserInvalidName(t *testing.T) {
	testManager := NewManager()

	testFirstName := ""
	testLastName := "Userman"
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Fatalf("error parsing test email address: %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err == nil {
		t.Error("Noo error returned for invalid name")
	} else {
		exceptedErr := "invalid first name\n"
		if err.Error() != exceptedErr {
			t.Errorf("bad error text,wanted: %s, got %s", exceptedErr, err)
		}
	}
	if len(testManager.users) > 0 {
		t.Errorf("bad test manager user count, wanted %d, got %d", 0, len(testManager.users))
	}

}

func TestGetUserByName(t *testing.T) {
	testManager := NewManager()

	err := testManager.AddUser("foo", "bar", "f.bar@example.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}
	err = testManager.AddUser("bar", "baz", "bbaz@example.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}
	err = testManager.AddUser("foo", "baz", "fbaz@example.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}
	err = testManager.AddUser("baz", "foo", "bazf@example.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}

	tests := map[string]struct {
		first         string
		last          string
		expected      *User
		expectedError error
	}{
		"simple lookup": {
			first:         "foo",
			last:          "bar",
			expected:      &testManager.users[0],
			expectedError: ErrNoResultsFound,
		},
		"last element lookup": {
			first:         "baz",
			last:          "foo",
			expected:      &testManager.users[3],
			expectedError: ErrNoResultsFound,
		},
		"no match lookup": {
			first:         "quux",
			last:          "quuz",
			expected:      nil,
			expectedError: ErrNoResultsFound,
		},
		"empty first name": {
			first:         "",
			last:          "baz",
			expected:      nil,
			expectedError: ErrNoResultsFound,
		},
		"empty last name": {
			first:         "foo",
			last:          "",
			expected:      nil,
			expectedError: ErrNoResultsFound,
		},
	}

	for name, test := range tests {
		result, err := testManager.GetUserByName(test.first, test.last)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("%s: invalid result\n got: %+v\n wanted: %+v\n", name, result, test.expected)
		}
		if !errors.Is(err, test.expectedError) {
			t.Errorf("%s:invalid error result\ngot: %v\nwanted:%v", name, result, test.expectedError)
		}
	}

}

func TestAddUserDuplicateName(t *testing.T) {
	testManager := NewManager()

	testFirstName := "Test"
	testLastName := "Userman"
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Fatalf("error parsing test email address: %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}
	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())

	if err == nil {
		t.Error("Noo error returned for duplicate user")
	} else {
		exceptedErr := "user with this name already exists"
		if err.Error() != exceptedErr {
			t.Errorf("bad error text,wanted: %s, got %s", exceptedErr, err)
		}
	}
	if len(testManager.users) != 1 {
		t.Errorf("bad test manager user count, wanted %d, got %d", 0, len(testManager.users))
	}
}
