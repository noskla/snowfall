package main

type room struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StreamURL   string `json:"streamURL"`
}
func getAllRooms() (bool, string, []room) {

	rows, err := Database.Query(`select * from rooms`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction query)", []room{}
	}
	defer rows.Close()

	var rooms []room
	for rows.Next() {
		var ID          string
		var name        string
		var description string
		var streamURL   string
		err := rows.Scan(&ID, &name, &description, &streamURL)
		if errorOccurred(err, false) {
			return false, "Error creating room slice", rooms
		}
		rooms = append(rooms, room{ID, name, description, streamURL})
	}

	return true, "Ok", rooms

}

