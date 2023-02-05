/*
Copyright © 2023 Melvin Wiens <EMAIL ADDRESS>
*/
package cmd

import "encoding/json"

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)

}
