package file

// LoadString loads a file into a string buf
func LoadString(path string) (*string, error) {
	f, err := LoadBytes(path)
	if err != nil {
		return nil, err
	}

	loaded := string(f)

	return &loaded, nil
}
