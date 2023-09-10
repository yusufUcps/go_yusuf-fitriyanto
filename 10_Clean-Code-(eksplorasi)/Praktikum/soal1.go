package main

//nama atribut diawali huruf besar
//setiap nama variabel digandeng dengan huruf besar kataberikutnya

type User struct {
	ID       int
	Username string // diganti menjadi string karna int tidak tepat dengan username
	Password string // diganti menjadi string karna int tidak tepat dengan password
}

type UserService struct {
	Users []User// variabel tidak deskriptif maka diganti users
}

// GetAllUsers mengembalikan semua pengguna.
func (u UserService) GetAllUsers() []User {
	return u.Users
}

// GetUserByID mengembalikan pengguna dengan ID yang sesuai.
// Jika tidak ditemukan, mengembalikan nil.
func (u UserService) GetUserByID(id int) *User {
	for _, user := range u.Users { // karna r tidak deskriptif maka diganti dengan user
		if id == user.ID {
			return &user
		}
	}

	return nil
}
