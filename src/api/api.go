package api

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "github.com/alok87/goutils/pkg/random"
)

func CatFact() (string) {
    URL := "https://catfact.ninja/fact?max_length=2000"
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	map_ := map[string]string{}
	json.Unmarshal(body, &map_)
	data := map_["fact"];
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }
    return data
}

func GeekJoke() (string) {
	
    URL := "https://geek-jokes.sameerkumar.website/api"
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	var map_ string
	json.Unmarshal(body, &map_)
	data := map_;
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }
    return data

}

func FoxPic() (string) {
	
    URL := "https://randomfox.ca/floof/"
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	map_ := map[string]string{}
	json.Unmarshal(body, &map_)
	data := map_["image"];
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }
    return data
}

func Fact() (string) {
	
    URL := "https://uselessfacts.jsph.pl/random.json?language=en"
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	map_ := map[string]string{}
	json.Unmarshal(body, &map_)
	data := map_["text"];
	source_url := map_["source_url"];
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }
    
    return data + "\nSource: " + source_url
}

func Xkcd() (string) {
	
	randNum := random.RangeInt(1, 100, 1)[0]
	
    URL := fmt.Sprintf("https://xkcd.com/%d/info.0.json", randNum)
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	map_ := map[string]string{}
	json.Unmarshal(body, &map_)
	data := map_["img"];
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }
    
    return data
}

func TechQuote() (string) {
	
    URL := "http://quotes.stormconsultancy.co.uk/random.json"
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	map_ := map[string]string{}
	json.Unmarshal(body, &map_)
	data := map_["quote"];
	fmt.Println(data);
	author := map_["author"];
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }

    return data + "\n- " + author
}

func StartupQuote() (string) {
	
    URL := "https://wisdomapi.herokuapp.com/v1/random"
    
    resp, err := http.Get(URL)
    if err != nil {
        panic(err)
    }
    
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	map_ := map[string]interface{}{}
	json.Unmarshal(body, &map_)
	data := map_["content"].(string);
	fmt.Println(data);
	
	author := map_["author"];
	author_map := author.(map[string]interface{})
	author_name := author_map["name"].(string)
	author_company := author_map["company"].(string)
	fmt.Println(author_name);
	fmt.Println(author_company);
	
    err = ioutil.WriteFile("output.txt", body, 0644)
    if err != nil {
        panic(err)
    }
    
    return data + "\n- " + author_name + " (" + author_company + ")"
}