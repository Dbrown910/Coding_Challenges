package main

import ("fmt"
		"bytes"
		"bufio"
		"os"
		"io/ioutil"
		"net/http"
		"strings"
		"encoding/json"
		"sort"
		)

type ImageList struct {
	Images []Image
}

type Image struct {
	url string
	concepts []Concept
}

type Concept struct {
	Name string `json:"name"`
	Prob float32 `json:"value"`
}

type ImageConcept struct {
	url string
	name string
	prob float32
}

type ClarifaiApiResponse struct {
	Outputs []struct {
		Status struct {
			Code int32 `json:"code"`
			Desc string `json:"description"`
		} `json:"status"`
		Created_at string `json:"created_at"`
		Input struct {
			Id string `json:"id"`
			Data struct {
				Image struct {
					Url string `json:"url"`
				} `json:"image"`
			} `json:"data"`
		} `json:"input"`
		Data struct {
			ConceptList []Concept `json:"concepts"`
		}`json:"data"`
	} `json:"outputs"`
}

func main() {
	fmt.Println("Processing image inputs. This will take a few minutes. \nPlease Wait...")

	var userInput string
	imageDataList := ImageList{}

	imageRequests := parseImageUrlsForRequest()

	for _, urlSet := range imageRequests {
		var apiResponse *ClarifaiApiResponse = getImageData(urlSet)
		imageDataList = parseResponseInfo(apiResponse, imageDataList)
	}

	// var apiResponse *ClarifaiApiResponse = getImageData(imageRequests[0])
	// imageDataList = parseResponseInfo(apiResponse, imageDataList)

	fmt.Println("Processing Complete\n")

	// Prompt for input on STDIN
	for true {
		fmt.Print("Enter an image tag to search or 0 to exit: ")
    	fmt.Scanln(&userInput)    
    	if (userInput == "0") {
    		break
    	} else {
    		searchAllImagesForMatch(strings.TrimRight(userInput, "\n"), imageDataList)
    	}
	}
}

func searchAllImagesForMatch(searchString string, allImageData ImageList) {
	// Cycle through all image data and find matching concepts. Store urls with concepts for output
	var outputList []ImageConcept

	for _, img := range allImageData.Images {
		for j := range img.concepts {
			if img.concepts[j].Name == searchString {
				hybrid := ImageConcept{img.url, img.concepts[j].Name, img.concepts[j].Prob}
				if (len(outputList) < 10) {
					// Just add the new object to the list and sort it
					outputList = append(outputList, hybrid)
					sort.Slice(outputList, func(i, j int) bool {
						return outputList[i].prob > outputList[j].prob
					})
				} else {
					// List is already sorted so tack the newest value to the end and sort the slice of the first ten objects
					outputList[9] = hybrid
					sort.Slice(outputList[:10], func(i, j int) bool {
						return outputList[i].prob > outputList[j].prob
					})
				}
			}
		}
	}

	outputResults(outputList, searchString) 
}

func parseResponseInfo(response *ClarifaiApiResponse, allImageData ImageList) ImageList {
	// For each output, add response data to image struct and then append it to allImageData
	for _, currImageData := range response.Outputs {
		var newImage Image = Image{}
		newImage.concepts = currImageData.Data.ConceptList
		newImage.url = currImageData.Input.Data.Image.Url

		allImageData.Images = append(allImageData.Images, newImage)
	}

	return allImageData
}

func getImageData(reqData string) *ClarifaiApiResponse {
	// Hit Clarifai API for data on image set
	base_url := "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
	api_key := "cf2575dbba734dba838698b51a5620e5"

	var jsonBytes = []byte(reqData)

	req, err := http.NewRequest("POST", base_url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Key " + api_key)
	req.Header.Add("cache-control", "no-cache")

	// fmt.Println(formatRequest(req))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	//fmt.Println(string(body))

	var apiData = new(ClarifaiApiResponse)

	err = json.Unmarshal(body, &apiData)

	if err != nil {
        fmt.Println(err)
    }

    return apiData	
}

func parseImageUrlsForRequest() [8]string {
	// 1000 image urls divided up in to JSON object strings containing 128 or less urls each
	var allImageUrls [8]string
	nextArrIndex := 0

	file, err := os.Open(".\\ClarifaiCodingChallenge\\assets\\images.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Single Image Req 
	// jsonReq := `{"inputs": [{"data": {"image": {"url": "https://farm1.staticflickr.com/2934/14439122755_d4af7552d1_o.jpg"}}}]}`
	
	// Build a list of 128 image links to send in one request, and add to array
	reqStr := `{"inputs": [`
	i := 0
	for scanner.Scan() {
		if i == 128 {
			// remove the last comma and cap the JSON object
			reqLen := len(reqStr)

			if reqLen > 0 && reqStr[reqLen-1] == ',' {
			    reqStr = reqStr[:reqLen-1]
			}
			reqStr += `]}`

			allImageUrls[nextArrIndex] = reqStr
			nextArrIndex++
			i = 0 
			reqStr = `{"inputs": [`
		}
		newData := `{"data": {"image": {"url": "`+scanner.Text()+`"}}},`
		reqStr += newData
		i++
	}

	// naive way of checking if there are still urls to add at the EOF. Should change to a Reader and make if EOF check
	if i > 0 {
		// remove the last comma and cap the JSON object
		reqLen := len(reqStr)

		if reqLen > 0 && reqStr[reqLen-1] == ',' {
		    reqStr = reqStr[:reqLen-1]
		}
		reqStr += `]}`

		allImageUrls[nextArrIndex] = reqStr
	}

	if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    return allImageUrls
}

func outputResults(results []ImageConcept, query string) {
	if len(results) == 0 {
		fmt.Println("No images matched your search query")
	} else {
		fmt.Println("\nTop Matches")
		fmt.Println("===========")
		for i, r := range results {
			fmt.Printf("%v.) %s \n", i+1, r.url)
			fmt.Printf("Likelihood of %s = %.2f\n", query, r.prob)
		}	
		fmt.Println();
	}
}

// Function Credit = https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
	name = strings.ToLower(name)
	for _, h := range headers {
	 request = append(request, fmt.Sprintf("%v: %v", name, h))
	}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
	r.ParseForm()
	request = append(request, "\n")
	request = append(request, r.Form.Encode())
	} 
	// Return the request as a string
	return strings.Join(request, "\n")
}