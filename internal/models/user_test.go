package models

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name           string
		user           User
		validationMode ValidationMode
		wantErr        error
	}{
		{
			name: "valid new user",
			user: User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			validationMode: ValidationCreate,
			wantErr:        nil,
		},
		{
			name: "valid existing user without password",
			user: User{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			validationMode: ValidationUpdate,
			wantErr:        nil,
		},
		{
			name: "blank name",
			user: User{
				Name:     "",
				Email:    "john@example.com",
				Password: "password123",
			},
			validationMode: ValidationCreate,
			wantErr:        ErrNameBlank,
		},
		{
			name: "name with only spaces",
			user: User{
				Name:     "   ",
				Email:    "john@example.com",
				Password: "password123",
			},
			validationMode: ValidationCreate,
			wantErr:        ErrNameBlank,
		},
		{
			name: "blank email",
			user: User{
				Name:     "John Doe",
				Email:    "",
				Password: "password123",
			},
			validationMode: ValidationCreate,
			wantErr:        ErrEmailBlank,
		},
		{
			name: "email with only spaces",
			user: User{
				Name:     "John Doe",
				Email:    "   ",
				Password: "password123",
			},
			validationMode: ValidationCreate,
			wantErr:        ErrEmailBlank,
		},
		{
			name: "invalid email format",
			user: User{
				Name:     "John Doe",
				Email:    "invalid-email",
				Password: "password123",
			},
			validationMode: ValidationCreate,
			wantErr:        ErrInvalidEmailFormat,
		},
		{
			name: "blank password for new user",
			user: User{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			validationMode: ValidationCreate,
			wantErr:        ErrPasswordBlank,
		},
		{
			name: "password with spaces for new user",
			user: User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "   ",
			},
			validationMode: ValidationCreate,
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate(tt.validationMode)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUser_isNameValid(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{
			name:    "valid name",
			user:    User{Name: "John Doe"},
			wantErr: nil,
		},
		{
			name:    "blank name",
			user:    User{Name: ""},
			wantErr: ErrNameBlank,
		},
		{
			name:    "name with only spaces",
			user:    User{Name: "   "},
			wantErr: ErrNameBlank,
		},
		{
			name:    "name with leading/trailing spaces",
			user:    User{Name: "  John Doe  "},
			wantErr: nil,
		},
		{
			name:    "single character name",
			user:    User{Name: "J"},
			wantErr: nil,
		},
		{
			name:    "name with special characters",
			user:    User{Name: "João O'Brien"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.isNameValid()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUser_isEmailValid(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{
			name:    "valid email",
			user:    User{Email: "john@example.com"},
			wantErr: nil,
		},
		{
			name:    "valid email with subdomain",
			user:    User{Email: "john@mail.example.com"},
			wantErr: nil,
		},
		{
			name:    "valid email with plus",
			user:    User{Email: "john+test@example.com"},
			wantErr: nil,
		},
		{
			name:    "valid email with numbers",
			user:    User{Email: "john123@example.com"},
			wantErr: nil,
		},
		{
			name:    "blank email",
			user:    User{Email: ""},
			wantErr: ErrEmailBlank,
		},
		{
			name:    "email with only spaces",
			user:    User{Email: "   "},
			wantErr: ErrEmailBlank,
		},
		{
			name:    "invalid email without @",
			user:    User{Email: "johnexample.com"},
			wantErr: ErrInvalidEmailFormat,
		},
		{
			name:    "invalid email without domain",
			user:    User{Email: "john@"},
			wantErr: ErrInvalidEmailFormat,
		},
		{
			name:    "invalid email without local part",
			user:    User{Email: "@example.com"},
			wantErr: ErrInvalidEmailFormat,
		},
		{
			name:    "invalid email with spaces",
			user:    User{Email: "john doe@example.com"},
			wantErr: ErrInvalidEmailFormat,
		},
		{
			name:    "email with leading/trailing spaces",
			user:    User{Email: "  john@example.com  "},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.isEmailValid()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUser_isNewRegister(t *testing.T) {
	tests := []struct {
		name           string
		validationMode ValidationMode
		want           bool
	}{
		{
			name:           "new register",
			validationMode: ValidationCreate,
			want:           true,
		},
		{
			name:           "existing register",
			validationMode: ValidationUpdate,
			want:           false,
		},
		{
			name:           "invalid mode",
			validationMode: ValidationMode(-1),
			want:           false,
		},
		{
			name:           "update register",
			validationMode: ValidationUpdate,
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			got := u.isNewRegister(tt.validationMode)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_hashPasswd(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		register string
		wantErr  bool
	}{
		{
			name: "hash password for new user",
			user: User{
				Password: "password123",
			},
			register: "new",
			wantErr:  false,
		},
		{
			name: "no hash for existing user",
			user: User{
				Password: "already-hashed-password",
			},
			register: "existing",
			wantErr:  false,
		},
		{
			name: "empty password for new user",
			user: User{
				Password: "",
			},
			register: "new",
			wantErr:  false,
		},
		{
			name: "password with special characters",
			user: User{
				Password: "P@ssw0rd!#$%",
			},
			register: "new",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalPassword := tt.user.Password
			err := tt.user.hashPasswd(tt.register)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				if tt.register == "new" && originalPassword != "" {
					assert.NotEqual(t, originalPassword, tt.user.Password)
					assert.NotEmpty(t, tt.user.Password)
				}

				if tt.register != "new" {
					assert.Equal(t, originalPassword, tt.user.Password)
				}
			}
		})
	}
}

func TestUser_Struct(t *testing.T) {
	t.Run("create user with all fields", func(t *testing.T) {
		id := uuid.New()
		user := User{
			ID:       id,
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		assert.Equal(t, id, user.ID)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "john@example.com", user.Email)
		assert.Equal(t, "password123", user.Password)
		assert.True(t, user.CreatedAt.IsZero())
	})

	t.Run("create minimal user", func(t *testing.T) {
		user := User{
			Name:  "Jane Doe",
			Email: "jane@example.com",
		}

		assert.Empty(t, user.ID)
		assert.Equal(t, "Jane Doe", user.Name)
		assert.Equal(t, "jane@example.com", user.Email)
		assert.Empty(t, user.Password)
	})
}

// TestErrors verifies all custom errors are defined and distinct
func TestErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		msg  string
	}{
		{
			name: "ErrNameBlank",
			err:  ErrNameBlank,
			msg:  "name can not be blank",
		},
		{
			name: "ErrEmailBlank",
			err:  ErrEmailBlank,
			msg:  "email can not be blank",
		},
		{
			name: "ErrPasswordBlank",
			err:  ErrPasswordBlank,
			msg:  "password can not be blank",
		},
		{
			name: "ErrInvalidEmailFormat",
			err:  ErrInvalidEmailFormat,
			msg:  "invalid format",
		},
		{
			name: "ErrEmailAlreadyExist",
			err:  ErrEmailAlreadyExist,
			msg:  "email already used",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.err)
			assert.Equal(t, tt.msg, tt.err.Error())
		})
	}

	// Verify errors are distinct
	t.Run("errors are distinct", func(t *testing.T) {
		assert.NotEqual(t, ErrNameBlank, ErrEmailBlank)
		assert.NotEqual(t, ErrNameBlank, ErrPasswordBlank)
		assert.NotEqual(t, ErrEmailBlank, ErrPasswordBlank)
		assert.NotEqual(t, ErrInvalidEmailFormat, ErrEmailBlank)
		assert.NotEqual(t, ErrEmailAlreadyExist, ErrInvalidEmailFormat)
	})
}

// TestErrorWrapping tests that errors.Is works correctly
func TestErrorWrapping(t *testing.T) {
	t.Run("errors.Is works with custom errors", func(t *testing.T) {
		err := ErrNameBlank
		assert.True(t, errors.Is(err, ErrNameBlank))
		assert.False(t, errors.Is(err, ErrEmailBlank))
	})
}
