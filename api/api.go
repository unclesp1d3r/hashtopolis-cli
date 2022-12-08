/*
Copyright © 2022 UncleSp1d3r <unclespider@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	Action struct {
		Section   string `json:"section"`
		Request   string `json:"request"`
		AccessKey string `json:"accessKey"`
	}
	Response struct {
		Section  string `json:"section"`
		Request  string `json:"request"`
		Response string `json:"response"`
		Message  string `json:"message"`
	}
	ListTasksResponse struct {
		Section  string `json:"section"`
		Request  string `json:"request"`
		Response string `json:"response"`
		Tasks    []struct {
			TaskId      int    `json:"taskId,omitempty"`
			Name        string `json:"name"`
			Type        int    `json:"type"`
			HashlistId  int    `json:"hashlistId"`
			Priority    int    `json:"priority"`
			SupertaskId int    `json:"supertaskId,omitempty"`
		} `json:"tasks"`
	}
	ListHashlistsResponse struct {
		Section   string `json:"section"`
		Request   string `json:"request"`
		Response  string `json:"response"`
		Hashlists []struct {
			HashlistId int    `json:"hashlistId"`
			HashtypeId int    `json:"hashtypeId"`
			Name       string `json:"name"`
			Format     int    `json:"format"`
			HashCount  int    `json:"hashCount"`
		} `json:"hashlists"`
	}
)

func TestConnection(url, key string) error {
	action := Action{
		Section:   "test",
		Request:   "access",
		AccessKey: key,
	}

	body, _ := json.Marshal(action)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		response := Response{}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return err
		}

		if response.Response != "OK" {
			return errors.New(response.Message)
		}

	} else {
		fmt.Println("Get failed with error: ", resp.Status)
	}
	return nil
}

func ListTasks(url, key string) (ListTasksResponse, error) {
	action := Action{
		Section:   "task",
		Request:   "listTasks",
		AccessKey: key,
	}

	body, _ := json.Marshal(action)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return ListTasksResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ListTasksResponse{}, err
		}

		response := ListTasksResponse{}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return ListTasksResponse{}, err
		}

		if response.Response != "OK" {
			return ListTasksResponse{}, errors.New(response.Response)
		}

		return response, nil

	} else {
		fmt.Println("Get failed with error: ", resp.Status)
	}
	return ListTasksResponse{}, nil
}

func ListHashlists(url, key string) (ListHashlistsResponse, error) {
	action := Action{
		Section:   "hashlist",
		Request:   "listHashlists",
		AccessKey: key,
	}

	body, _ := json.Marshal(action)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return ListHashlistsResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ListHashlistsResponse{}, err
		}

		response := ListHashlistsResponse{}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return ListHashlistsResponse{}, err
		}

		if response.Response != "OK" {
			return ListHashlistsResponse{}, errors.New(response.Response)
		}

		return response, nil

	} else {
		fmt.Println("Get failed with error: ", resp.Status)
	}
	return ListHashlistsResponse{}, nil
}
