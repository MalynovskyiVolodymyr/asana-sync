package services

import (
	"asanaSync/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GetTodoByID Traction tools
func GetTodoByID() {
	url := "https://traction.tools/Todo/Pad/2216892"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer aYxZvEt0fX0U1WmMmejKSgyyjoyh5ekwDo4m6UtRzxpYxd10bK5mlonFzR2ww33N_5JgCuFQxHdPolhOy33ISYTJgexcrsZe2IDRav8U9hEJxVWH_UKl1nsIMYbSLIEaiVmNM6_MuL0K9BWCane6r4DA0NVG44SKUEphRaE_RPfFj7bkBoJzIhydt2KIEXVQOiCocD7VBYcE1DioLjIsDIDAEFZiFmW9_gyD4uiH8gRse7og8XSp1E5N-upraPLuSRgZmwhrDSVqOCPus7NoI_eShyz72mntsnQCDdmnBDPcipTwOFLrQFCacU5pl-JrQ7gNUKq9htht0LqpZ3pDQE1esugM08rpxgypATDCTh4_UjI8QgbWkEZnY2Nu_Q-CuZb_Cwc50gD4IrJg6m58F7ocpOlVCUXbki3-c8vjSKl4SfVVoSD8lORzKqy2Ygo8eZWs3SBsSKeGajrbOLhDD3vPOHUO8gIS7-h8WwJWFwfCMt4-3BfN-hxjO-C5jhRw")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// GetTasksByMitting Traction tools
func GetTasksByMitting(mittingID int64) []models.TractionTaskModel {
	var result []models.TractionTaskModel
	var filteredResult []models.TractionTaskModel

	url := "https://traction.tools/api/v1/l10/" + strconv.FormatInt(mittingID, 10) + "/todos"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer aYxZvEt0fX0U1WmMmejKSgyyjoyh5ekwDo4m6UtRzxpYxd10bK5mlonFzR2ww33N_5JgCuFQxHdPolhOy33ISYTJgexcrsZe2IDRav8U9hEJxVWH_UKl1nsIMYbSLIEaiVmNM6_MuL0K9BWCane6r4DA0NVG44SKUEphRaE_RPfFj7bkBoJzIhydt2KIEXVQOiCocD7VBYcE1DioLjIsDIDAEFZiFmW9_gyD4uiH8gRse7og8XSp1E5N-upraPLuSRgZmwhrDSVqOCPus7NoI_eShyz72mntsnQCDdmnBDPcipTwOFLrQFCacU5pl-JrQ7gNUKq9htht0LqpZ3pDQE1esugM08rpxgypATDCTh4_UjI8QgbWkEZnY2Nu_Q-CuZb_Cwc50gD4IrJg6m58F7ocpOlVCUXbki3-c8vjSKl4SfVVoSD8lORzKqy2Ygo8eZWs3SBsSKeGajrbOLhDD3vPOHUO8gIS7-h8WwJWFwfCMt4-3BfN-hxjO-C5jhRw")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	err = json.Unmarshal(body, &result)

	for _, v := range result {
		if !v.Complete {
			filteredResult = append(filteredResult, v)
		}
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return filteredResult
}

// GetMyWieviable Traction tools
func GetMyWieviable() {
	url := "https://traction.tools/api/v1/users/mineviewable"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer CkxcN_xuzeVuMrUj1YfEzV7SZgLSf2TT9oFuEtuPK9TLYhoelQd6i2SO2EbS4SxiXh3t7_fTNjPSD8AYo0JnSrql6zeTGQ7CRmDGN5OcKoVmm0qNPJCiUX-_qxyTYJBFRJaitO-DAiMRpYAjvHpD0PUpsTegTPwcl3q6xuf7hOn29UbKMQ830g41YLcebahuisLBYKRkpSoM5KPF5XbENZ2jBR-nuYbjzARYd_vptZ72A3cluuRjSTHjqek5UaP9wWkAxCh06bcStQzJjHN1Y44HiPooj4blJaXlKraN6uKiOmNj5KMOyxolCP7ANaf45-29yZ4Paw7mtH1zEboCu10XFoDBm8L4QEQjZZP5ds4dY5IQX731rnkohqerlNz4BywB9Mvb5c8IvkPwDzJ98lNci5VC4LfX0aINdEozlnhXCIwjDkXT1LwjOwAaE5smSQP4-g52zxMIFBcIQuaA5OOgA5IYyucMAfvpFfAWPyhKrS2dBlMWlYWbB6p70tPU")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// GetTractionUserTasks Traction tools
func GetTractionUserTasks() []models.TractionTaskModel {
	var result []models.TractionTaskModel
	var filteredResult []models.TractionTaskModel

	url := "https://traction.tools/api/v1/todo/users/mine?"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer nPHWqYNCFmVOWRQA8r8XT-jX6Du2UHQr3MfLD5JuaU27i52vYSwCLKR3UImGpHQezted0VBCDL1pMaCy2j3t2-spmfQeMCZC2IWL0k8S6nfEsfdl2Gs8EiB2xzdOKJRtE4CpFUDiEDc16lEtY1LT3mJoDRk_L8xL-z6JAXm_Kq7YNWqen0QYK6pccOY1qgNrlU02jFdAIwUs43eC7RWxmJS_mi1bFbwkq67YzQydmh8x48wYUMJERBVUsvu1TreW2WYOK6t0t-agSuQ8tDc-q0Sh_9aAMKGhKKB8Dmzo9x4vrA7Y3g5rljiTE-CbfTUxG_J0kTpQIpAh0UkZNMi_LTwiqx_-dffhiLe4mcsiRWAcb6IUD_f2YDv4rD6iwDEEoFsY5sJYTBRSunplJez-dkGHo9y7oIUzfbtM5NvHPkxwShhSCVsCyCvCrtKOrei9yLMNKN9vvTJherjaSHnvEY1orWi5vMt3BtCGY1i9RG7nxq1GWOuWx5GHAyFPoI7e")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	e := json.Unmarshal(body, &result)

	if e != nil {
		fmt.Println(e)
		return nil
	}

	for _, v := range result {
		if !v.Complete {
			filteredResult = append(filteredResult, v)
		}

	}

	return filteredResult
}

// GetAllMitingsFromTractionTool traction tools
func GetAllMitingsFromTractionTool() {
	url := "https://traction.tools/api/v1/L10/list"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer aYxZvEt0fX0U1WmMmejKSgyyjoyh5ekwDo4m6UtRzxpYxd10bK5mlonFzR2ww33N_5JgCuFQxHdPolhOy33ISYTJgexcrsZe2IDRav8U9hEJxVWH_UKl1nsIMYbSLIEaiVmNM6_MuL0K9BWCane6r4DA0NVG44SKUEphRaE_RPfFj7bkBoJzIhydt2KIEXVQOiCocD7VBYcE1DioLjIsDIDAEFZiFmW9_gyD4uiH8gRse7og8XSp1E5N-upraPLuSRgZmwhrDSVqOCPus7NoI_eShyz72mntsnQCDdmnBDPcipTwOFLrQFCacU5pl-JrQ7gNUKq9htht0LqpZ3pDQE1esugM08rpxgypATDCTh4_UjI8QgbWkEZnY2Nu_Q-CuZb_Cwc50gD4IrJg6m58F7ocpOlVCUXbki3-c8vjSKl4SfVVoSD8lORzKqy2Ygo8eZWs3SBsSKeGajrbOLhDD3vPOHUO8gIS7-h8WwJWFwfCMt4-3BfN-hxjO-C5jhRw")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// SearchTractionTool traction tools
func SearchTractionTool() {
	url := "https://traction.tools/api/v1/search/user?term=Vladimir&20Malinovskiy"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer nPHWqYNCFmVOWRQA8r8XT-jX6Du2UHQr3MfLD5JuaU27i52vYSwCLKR3UImGpHQezted0VBCDL1pMaCy2j3t2-spmfQeMCZC2IWL0k8S6nfEsfdl2Gs8EiB2xzdOKJRtE4CpFUDiEDc16lEtY1LT3mJoDRk_L8xL-z6JAXm_Kq7YNWqen0QYK6pccOY1qgNrlU02jFdAIwUs43eC7RWxmJS_mi1bFbwkq67YzQydmh8x48wYUMJERBVUsvu1TreW2WYOK6t0t-agSuQ8tDc-q0Sh_9aAMKGhKKB8Dmzo9x4vrA7Y3g5rljiTE-CbfTUxG_J0kTpQIpAh0UkZNMi_LTwiqx_-dffhiLe4mcsiRWAcb6IUD_f2YDv4rD6iwDEEoFsY5sJYTBRSunplJez-dkGHo9y7oIUzfbtM5NvHPkxwShhSCVsCyCvCrtKOrei9yLMNKN9vvTJherjaSHnvEY1orWi5vMt3BtCGY1i9RG7nxq1GWOuWx5GHAyFPoI7e")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// GetUserToken Traction tools
func GetUserToken(username string, password string) {
	url := "https://traction.tools/Token"

	payload := strings.NewReader("grant_type=password&userName=" + username + "&password=" + password)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	// req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}



