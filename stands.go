package main

type stand struct {
	ID 	        string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}
func getAllStands() (bool, string, []stand) {

	rows, err := Database.Query(`select id, name, description from stands`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction query)", []stand{}
	}
	defer rows.Close()

	var stands []stand
	for rows.Next() {
		var ID string
		var Name string
		var Description string
		err := rows.Scan(&ID, &Name, &Description)
		if errorOccurred(err, false) {
			return false, "Error creating stands slice", stands
		}
		stands = append(stands, stand{ID, Name, Description})
	}

	return true, "Ok", stands

}
