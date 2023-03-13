package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	MAX_REQUESTS = 10 // Bing server-side limit
)

var (
	client                = &http.Client{}
	dialer                = &websocket.Dialer{}
	conversationId        string
	clientId              string
	conversationSignature string

	requestCount = 0
)

func createChat() (CreateResponse, error) {
	cookie, err := getCookie()
	if err != nil {
		return CreateResponse{}, err
	}

	request, err := http.NewRequest("GET", "https://www.bing.com/turing/conversation/create", nil)
	if err != nil {
		return CreateResponse{}, err
	}

	request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Referer", "https://www.bing.com/search?q=I need to throw a dinner party for 6 people who are vegetarian. Can you suggest a 3-course menu with a chocolate dessert?&iscopilotedu=1&form=MA13G7")
	request.Header.Add("x-ms-useragent", "azsdk-js-api-client-factory/1.0.0-beta.1 core-rest-pipeline/1.10.0 OS/Linuxx86_64")
	request.Header.Add("x-edge-shopping-flag", "1")
	request.Header.Add("x-ms-client-request-id", uuid.NewString())
	request.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Microsoft Edge\";v=\"110\"")
	request.Header.Add("sec-ch-ua-arch", "\"x86\"")
	request.Header.Add("sec-ch-ua-bitness", "\"64\"")
	request.Header.Add("sec-ch-ua-full-version", "\"110.0.1587.50\"")
	request.Header.Add("sec-ch-ua-full-version-list", "\"Chromium\";v=\"110.0.5481.104\", \"Not A(Brand\";v=\"24.0.0.0\", \"Microsoft Edge\";v=\"110.0.1587.50\"")
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("sec-ch-ua-model", "")
	request.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	request.Header.Add("sec-ch-ua-platform-version", "\"6.1.11\"")
	request.Header.Add("sec-fetch-dest", "empty")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("accept", "application/json")
	request.Header.Add("accept-language", "en-US,en;q=0.9")

	response, err := client.Do(request)
	if err != nil {
		return CreateResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return CreateResponse{}, err
	}

	var createResponse CreateResponse
	err = json.Unmarshal(body, &createResponse)
	if err != nil {
		return CreateResponse{}, err
	}

	switch createResponse.Result.Value {
	case "Success":
		return createResponse, nil
	default:
		fmt.Println(string(body))
		return CreateResponse{}, errors.New("failed to create a new conversation")
	}
}

func search(query string) {
	conversation, err := createChat()
	if err != nil {
		log.Fatal(err)
	}

	conversationId = conversation.ConversationID
	clientId = conversation.ClientID
	conversationSignature = conversation.ConversationSignature

	u := url.URL{Scheme: "wss", Host: "sydney.bing.com", Path: "/sydney/ChatHub"}

	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: The server refused the connection.")
		os.Exit(2)
	}
	defer c.Close()

	// initial message
	message := []byte("{\"protocol\":\"json\",\"version\":1}\x1E")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// the server should say something
	_, _, err = c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	// we want to tell the server some more info about the type of request
	message = []byte("{\"type\":6}\x1E")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// get time for required timestamp parameter
	currentTime := time.Now()
	timeZone := currentTime.Format("-08:00")
	formattedTime := currentTime.Format("2006-01-02T15:04:05") + timeZone

	// send off our query
	//message = []byte("{\"arguments\":[{\"source\":\"cib\",\"optionsSets\":[\"deepleo\",\"enable_debug_commands\",\"disable_emoji_spoken_text\",\"enablemm\"],\"allowedMessageTypes\":[\"Chat\",\"InternalSearchQuery\",\"InternalSearchResult\",\"InternalLoaderMessage\",\"RenderCardRequest\",\"AdsQuery\",\"SemanticSerp\"],\"sliceIds\":[\"214dv3sc\",\"0113dllog\",\"215boep\",\"216dloffstream\",\"0213retry\"],\"traceId\":\"63f0373a2c4b4e7696a396f61d9f1131\",\"isStartOfSession\":true,\"message\":{\"locale\":\"en-US\",\"market\":\"en-US\",\"region\":\"US\",\"location\":\"lat:47.639557;long:-122.128159;re=1000m;\",\"locationHints\":[{\"country\":\"United States\",\"state\":\"California\",\"city\":\"Oakland\",\"zipcode\":\"94613\",\"timezoneoffset\":-8,\"dma\":807,\"countryConfidence\":9,\"cityConfidence\":8,\"Center\":{\"Latitude\":37.7852,\"Longitude\":-122.186},\"RegionType\":2,\"SourceType\":1}],\"timestamp\":\"" + formattedTime + "\",\"author\":\"user\",\"inputMethod\":\"Keyboard\",\"text\":\"" + query + "\",\"messageType\":\"Chat\"},\"conversationSignature\":\"" + conversation.ConversationSignature + "\",\"participant\":{\"id\":\"" + conversation.ClientID + "\"},\"conversationId\":\"" + conversation.ConversationID + "\"}],\"invocationId\":\"0\",\"target\":\"chat\",\"type\":4}\x1E")
	message = []byte("{\"arguments\":[{\"source\":\"cib\",\"optionsSets\":[\"nlu_direct_response_filter\",\"deepleo\",\"enable_debug_commands\",\"disable_emoji_spoken_text\",\"responsible_ai_policy_235\",\"enablemm\",\"multidcspcv\",\"dlislog\",\"bof106\",\"dloffstream\",\"offenseretry2\",\"caplg\",\"dv3sugg\"],\"allowedMessageTypes\":[\"Chat\",\"InternalSearchQuery\",\"InternalSearchResult\",\"InternalLoaderMessage\",\"RenderCardRequest\",\"AdsQuery\",\"SemanticSerp\"],\"sliceIds\":[\"216spacev\",\"214dv3sc\",\"0113dllog\",\"217bof106\",\"216dloffstream\",\"0213retry\",\"217fcr\"],\"traceId\":\"63f1cd8665a243dbbb537af7c2398e82\",\"isStartOfSession\":true,\"message\":{\"locale\":\"en-US\",\"market\":\"en-US\",\"region\":\"US\",\"location\":\"lat:47.639557;long:-122.128159;re=1000m;\",\"locationHints\":[{\"country\":\"United States\",\"state\":\"New York\",\"city\":\"New York\",\"zipcode\":\"10013\",\"timezoneoffset\":-5,\"dma\":501,\"Center\":{\"Latitude\":40.7206,\"Longitude\":-74.003},\"RegionType\":2,\"SourceType\":1}],\"timestamp\":\"" + formattedTime + "\",\"author\":\"user\",\"inputMethod\":\"Keyboard\",\"text\":\"" + query + "\",\"messageType\":\"Chat\"},\"conversationSignature\":\"" + conversationSignature + "\",\"participant\":{\"id\":\"" + clientId + "\"},\"conversationId\":\"" + conversationId + "\"}],\"invocationId\":\"0\",\"target\":\"chat\",\"type\":4}\x1E")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	requestCount++

	previousMessage := ""
	next := false

	// read the server data!
	for {
		_, response, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		msg := strings.Split(string(response), "\x1e")[0]

		var bingMessage BingMessage
		err = json.Unmarshal([]byte(msg), &bingMessage)
		if err != nil {
			fmt.Println(string(response))
			log.Fatal(err)
			return
		}

		previousMessage, next = handleMessage(previousMessage, response, bingMessage)
		if next {
			break
		}
	}
	if next {
		c.Close()

		fmt.Println()
		fmt.Print("> ")

		in := bufio.NewReader(os.Stdin)

		nextQuery, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// remove the newline from input
		nextQuery = strings.Split(nextQuery, "\n")[0]

		continueConversation(nextQuery, formattedTime)
	}
}

func handleMessage(previousMessage string, response []byte, bingMessage BingMessage) (string, bool) {
	switch bingMessage.Type {
	// 1 -> more info coming!
	case 1:
		if previousMessage == "" {
			if strings.Index(bingMessage.Arguments[0].Messages[0].Text, "Searching the web for:") != -1 {
				fmt.Println(" >", bingMessage.Arguments[0].Messages[0].Text)
			} else {
				fmt.Print(bingMessage.Arguments[0].Messages[0].Text)
				return bingMessage.Arguments[0].Messages[0].Text, false
			}
		} else {
			if len(strings.Split(bingMessage.Arguments[0].Messages[0].Text, previousMessage)) == 2 {
				fmt.Print(strings.Split(bingMessage.Arguments[0].Messages[0].Text, previousMessage)[1])
				return bingMessage.Arguments[0].Messages[0].Text, false
			} else if bingMessage.Arguments[0].Messages[0].MessageType != "RenderCardRequest" && bingMessage.Arguments[0].Messages[0].Offense == "Unknown" {
				fmt.Println()

				//fmt.Println(string(response))

				fmt.Print(bingMessage.Arguments[0].Messages[0].Text)
				return bingMessage.Arguments[0].Messages[0].Text, false
			} else {
				//fmt.Println()
			}
		}
		//fmt.Println(bingMessage.Arguments[0].Messages[0].Text)
	case 2:
		//fmt.Println()
		messages := strings.Split(string(response), "\x1e")
		if len(messages) > 1 {
			// I'll do this better in the future I promise! -> make the loop interpret these messages instead
			for _, m := range messages {
				//fmt.Println(m)
				var type2Message BingMessageType2
				err := json.Unmarshal([]byte(m), &type2Message)
				if err != nil {
					fmt.Println(m)
					log.Fatal(err)
				}

				if type2Message.Type == 2 {
					// we can do something with the sources and meta info? might be good...
					// TODO: use the metadata Bing gives us
					if type2Message.Item.Result.Error != "" {
						fmt.Println("Error: " + type2Message.Item.Result.Message)
						os.Exit(1)
					}
				} else if type2Message.Type == 3 {
					// user input, next query!
					return previousMessage, true
				}
			}
		} else {
			var type2Message BingMessageType2
			err := json.Unmarshal([]byte(messages[0]), &type2Message)
			if err != nil {
				fmt.Println(messages[0])
				log.Fatal(err)
			}

			fmt.Println("Error: " + type2Message.Item.Result.Message)
			os.Exit(1)
		}
	case 6:
		// ignore
	case 7:
		fmt.Println(string(response))
		//os.Exit(0)
	default:
		fmt.Println(string(response))
		os.Exit(1)
	}
	//fmt.Println()
	//fmt.Println(string(response))
	return previousMessage, false
}

func continueConversation(query string, formattedTime string) {
	// we need a new ChatHub connection
	u := url.URL{Scheme: "wss", Host: "sydney.bing.com", Path: "/sydney/ChatHub"}

	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println()
		fmt.Println("Error: The server refused the connection.")
		os.Exit(2)
	}
	defer c.Close()

	// initial message
	message := []byte("{\"protocol\":\"json\",\"version\":1}\x1E")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// the server should say something
	_, _, err = c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	// we want to tell the server some more info about the type of request
	message = []byte("{\"type\":6}\x1E")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// create our message content
	message = []byte("{\"arguments\":[{\"source\":\"cib\",\"optionsSets\":[\"nlu_direct_response_filter\",\"deepleo\",\"enable_debug_commands\",\"disable_emoji_spoken_text\",\"responsible_ai_policy_235\",\"enablemm\",\"multidcspcv\",\"dlislog\",\"bof106\",\"dloffstream\",\"offenseretry2\",\"caplg\",\"dv3sugg\"],\"allowedMessageTypes\":[\"Chat\",\"InternalSearchQuery\",\"InternalSearchResult\",\"InternalLoaderMessage\",\"RenderCardRequest\",\"AdsQuery\",\"SemanticSerp\"],\"sliceIds\":[\"216spacev\",\"214dv3sc\",\"0113dllog\",\"217bof106\",\"216dloffstream\",\"0213retry\",\"217fcr\"],\"traceId\":\"63f1cd8665a243dbbb537af7c2398e82\",\"isStartOfSession\":false,\"message\":{\"locale\":\"en-US\",\"market\":\"en-US\",\"region\":\"US\",\"location\":\"lat:47.639557;long:-122.128159;re=1000m;\",\"locationHints\":[{\"country\":\"United States\",\"state\":\"New York\",\"city\":\"New York\",\"zipcode\":\"10013\",\"timezoneoffset\":-5,\"dma\":501,\"Center\":{\"Latitude\":40.7206,\"Longitude\":-74.003},\"RegionType\":2,\"SourceType\":1}],\"timestamp\":\"" + formattedTime + "\",\"author\":\"user\",\"inputMethod\":\"Keyboard\",\"text\":\"" + query + "\",\"messageType\":\"Chat\"},\"conversationSignature\":\"" + conversationSignature + "\",\"participant\":{\"id\":\"" + clientId + "\"},\"conversationId\":\"" + conversationId + "\"}],\"invocationId\":\"1\",\"target\":\"chat\",\"type\":4}\x1E")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	requestCount++

	previousMessage := ""
	next := false

	// read the server data!
	for {
		_, response, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		msg := strings.Split(string(response), "\x1e")[0]

		var bingMessage BingMessage
		err = json.Unmarshal([]byte(msg), &bingMessage)
		if err != nil {
			log.Fatal(err)
			return
		}

		previousMessage, next = handleMessage(previousMessage, response, bingMessage)
		if next {
			break
		}
	}
	if next {
		c.Close()

		if requestCount == MAX_REQUESTS {
			os.Exit(0)
		}

		fmt.Println()
		fmt.Print("> ")

		in := bufio.NewReader(os.Stdin)

		nextQuery, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// remove the newline from input
		nextQuery = strings.Split(nextQuery, "\n")[0]

		continueConversation(nextQuery, formattedTime)
	}
}
