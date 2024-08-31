package Serveur

import (
	"fmt"
	"github.com/SpauriRosso/dotlog"
	"strconv"
)

func Filters(creationDate string, endCreationDate string, firstAlbum string, endFirstAlbum string, members []int) []int {
	var (
		res       []int
		cDate, _  = strconv.Atoi(creationDate)
		ecDate, _ = strconv.Atoi(endCreationDate)
		//fAlbum, _  = strconv.Atoi("1990")
		//efAlbum, _ = strconv.Atoi("1992")
		fAlbum, _  = strconv.Atoi(firstAlbum)
		efAlbum, _ = strconv.Atoi(endFirstAlbum)
	)
	getMembers := GetMembersCount(members)
	getfAlbums := GetFirstAlbum(fAlbum, efAlbum)
	getCreation := GetCreationDate(cDate, ecDate)
	res = FilterRes(getMembers, getfAlbums, getCreation)

	for _, v := range res {
		for _, w := range Sug {
			if v == w.Id {
				te := fmt.Sprintf("ID: %v, Name: %v", w.Id, w.Name)
				dotlog.Debug(te)
			}
		}
	}

	return res
}

func FilterRes(rids, aids, cids []int) []int {
	var res []int
	var count int
	idMap := make(map[int]int)

	// Check for empty slices, useful to correctly apply filters by using literal function
	checkTab := func(tab []int) {
		if len(tab) > 0 {
			count++
			for _, id := range tab {
				idMap[id]++
			}
		}
	}
	checkTab(rids) // YEAH I KNOW THIS COULD BE OPTIMIZED BUT JUST STFU!!!
	checkTab(aids)
	checkTab(cids)

	for id, counter := range idMap {
		// Only a matter of time, when did i start peeking at you. Darling i'll be fine, but its only a matter of love
		if counter == count {
			res = append(res, id)
		}
	}
	dotlog.Debug("len(res) = " + strconv.Itoa(len(res)))
	return res
}

func GetCreationDate(cDate int, ecDate int) []int {
	var ids []int
	if (cDate == 1939 && ecDate == 1939) || (cDate == 1939 || ecDate == 1939) {
		return nil
	}
	for _, v := range Sug {
		// DONT BE STUPID!!! cDate >= creationDate && artists.CreationDate <= ecDate
		// This one will pick every occurence of date prior to creation until ecDate
		if v.CreationDate >= cDate && v.CreationDate <= ecDate {
			ids = append(ids, v.Id)
		}
	}

	return ids
}

func GetMembersCount(members []int) []int {
	var ids []int
	for _, v := range Sug {
		for _, w := range members {
			if len(v.Members) == w {
				ids = append(ids, v.Id)
				break
			}
		}
	}
	return ids
}

func getConcertPlayed(location string) []int {
	var ids []int
	return ids
}

func GetFirstAlbum(fAlbum int, efAlbum int) []int {
	var ids []int

	if fAlbum == 1939 && efAlbum == 1939 {
		return nil
	}
	for _, v := range Sug {
		if len(v.FirstAlbum) >= 4 {
			yearStr := v.FirstAlbum[len(v.FirstAlbum)-4:]
			year, err := strconv.Atoi(yearStr)
			//ty := strconv.Itoa(year)
			//dotlog.Debug("Year is " + ty)
			if err == nil && year >= fAlbum && year <= efAlbum {
				ids = append(ids, v.Id)
			}
		}
	}
	return ids
}
