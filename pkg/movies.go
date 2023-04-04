package whatever

import (
	"encoding/json"
	"net/http"

	"github.com/jmulhern/whatever/pkg/kind"
)

var (
	mcu = kind.Group{
		Unique: "mcu",
		Name:   "Marvel Cinematic Universe",
		Movies: []kind.Movie{
			{Title: "Iron Man", Released: "May 2, 2008", Unique: "iron-man-1"},
			{Title: "The Incredible Hulk", Released: "June 13, 2008", Unique: "hulk-1"},
			{Title: "Iron Man 2", Released: "May 7, 2010", Unique: "iron-man-2"},
			{Title: "Thor", Released: "May 6, 2011", Unique: "thor-1"},
			{Title: "Captain America: The First Avenger", Released: "July 22, 2011", Unique: "captain-america-1"},
			{Title: "The Avengers", Released: "May 4, 2012", Unique: "avengers-1"},
			{Title: "Iron Man 3", Released: "May 3, 2013", Unique: "iron-man-3"},
			{Title: "Thor: The Dark World", Released: "November 8, 2013", Unique: "thor-2"},
			{Title: "Captain America: The Winter Soldier", Released: "April 4, 2014", Unique: "captain-america-2"},
			{Title: "Guardians of the Galaxy", Released: "August 1, 2014", Unique: "gotg-1"},
			{Title: "Avengers: Age of Ultron", Released: "May 1, 2015", Unique: "avengers-2"},
			{Title: "Ant-Man", Released: "July 17, 2015", Unique: "ant-man-1"},
			{Title: "Captain America: Civil War", Released: "May 6, 2016", Unique: "captain-america-3"},
			{Title: "Doctor Strange", Released: "November 4, 2016", Unique: "dr-strange-1"},
			{Title: "Guardians of the Galaxy Vol. 2", Released: "May 5, 2017", Unique: "gotg-2"},
			{Title: "Spider-Man: Homecoming", Released: "July 7, 2017", Unique: "spiderman-1"},
			{Title: "Thor: Ragnarok", Released: "November 3, 2017", Unique: "thor-3"},
			{Title: "Black Panther", Released: "February 16, 2018", Unique: "black-panther-1"},
			{Title: "Avengers: Infinity War", Released: "April 27, 2018", Unique: "avengers-3"},
			{Title: "Ant-Man and the Wasp", Released: "July 6, 2018", Unique: "ant-man-2"},
			{Title: "Captain Marvel", Released: "March 8, 2019", Unique: "captain-marvel-1"},
			{Title: "Avengers: Endgame", Released: "April 26, 2019", Unique: "avengers-4"},
			{Title: "Spider-Man: Far From Home", Released: "July 2, 2019", Unique: "spiderman-2"},
			{Title: "Black Widow", Released: "July 9, 2021", Unique: "black-widow"},
			{Title: "Shang-Chi and the Legend of the Ten Rings", Released: "September 3, 2021", Unique: "shang-chi-1"},
			{Title: "Eternals", Released: "November 5, 2021", Unique: "eternals-1"},
			{Title: "Spider-Man: No Way Home", Released: "December 17, 2021", Unique: "spiderman-3"},
			{Title: "Doctor Strange in the Multiverse of Madness", Released: "May 6, 2022", Unique: "dr-strange-2"},
			{Title: "Thor: Love and Thunder", Released: "July 8, 2022", Unique: "thor-4"},
			{Title: "Black Panther: Wakanda Forever", Released: "November 11, 2022", Unique: "black-panther-2"},
			{Title: "Ant-Man and the Wasp: Quantumania", Released: "February 17, 2023", Unique: "ant-man-3"},
			{Title: "Guardians of the Galaxy Vol. 3", Released: "May 5, 2023", Unique: "gotg-3"},
			{Title: "The Marvels", Released: "November 10, 2023", Unique: "marvels"},
			{Title: "Captain America: New World Order", Released: "May 3, 2024", Unique: "captain-america-4"},
			{Title: "Thunderbolts", Released: "July 26, 2024", Unique: "thunderbolts"},
			{Title: "Blade", Released: "September 6, 2024", Unique: "blade-1"},
			{Title: "Deadpool 3", Released: "November 8, 2024", Unique: "deadpool-3"},
			{Title: "Fantastic Four", Released: "February 14, 2025", Unique: "fantastic-four-1"},
			{Title: "Avengers: The Kang Dynasty", Released: "May 2, 2025", Unique: "avengers-5"},
			{Title: "Avengers: Secret Wars", Released: "May 1, 2026", Unique: "avengers-6"},
		},
	}

	starWars = kind.Group{
		Unique: "star-wars",
		Name:   "Star Wars",
		Movies: []kind.Movie{
			{Title: "Episode IV – A New Hope", Released: "May 25, 1977", Unique: "star-wars-4"},
			{Title: "Episode V – The Empire Strikes Back", Released: "May 21, 1980", Unique: "star-wars-5"},
			{Title: "Episode VI – Return of the Jedi", Released: "May 25, 1983", Unique: "star-wars-6"},
			{Title: "Episode I – The Phantom Menace", Released: "May 19, 1999", Unique: "star-wars-1"},
			{Title: "Episode II – Attack of the Clones", Released: "May 16, 2002", Unique: "star-wars-2"},
			{Title: "Episode III – Revenge of the Sith", Released: "May 19, 2005", Unique: "star-wars-3"},
			{Title: "Episode VII – The Force Awakens", Released: "December 15, 2015", Unique: "star-wars-7"},
			{Title: "Episode VIII – The Last Jedi", Released: "December 15, 2017", Unique: "star-wars-8"},
			{Title: "Episode IX – The Rise of Skywalker", Released: "December 20, 2019", Unique: "star-wars-9"},
		},
	}
)

func (h Handler) GetMovies(w http.ResponseWriter, _ *http.Request) {
	raw, _ := json.Marshal([]kind.Group{
		mcu,
		starWars,
	})
	_, _ = w.Write(raw)
}
