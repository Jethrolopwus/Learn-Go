
package storage

import "os"

type Storage interface {
    Save(data string) error
    Load() (string, error)
}

// FileStorage implements Storage using the file system
type FileStorage struct {
    FilePath string
}

// Save writes data to a file
func (fs FileStorage) Save(data string) error {
    return os.WriteFile(fs.FilePath, []byte(data), 0644)
}

// Load reads data from a file
func (fs FileStorage) Load() (string, error) {
    data, err := os.ReadFile(fs.FilePath)
    return string(data), err
}

// DatabaseStorage implements Storage using a database
type DatabaseStorage struct {
    ConnectionString string
}

// Save writes data to a database
func (ds DatabaseStorage) Save(data string) error {
    // Implementation to save to database
    return nil
}

// Load reads data from a database
func (ds DatabaseStorage) Load() (string, error) {
    // Implementation to load from database
    return "data from database", nil
}

// SaveUserData is a helper function that works with any Storage implementation
func SaveUserData(storage Storage, userData string) error {
    return storage.Save(userData)
}