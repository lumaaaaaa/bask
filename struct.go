package main

import "time"

type CreateResponse struct {
	ConversationID        string `json:"conversationId"`
	ClientID              string `json:"clientId"`
	ConversationSignature string `json:"conversationSignature"`
	Result                struct {
		Value   string `json:"value"`
		Message string `json:"message"`
	} `json:"result"`
}

type BingMessage struct {
	Type      int    `json:"type"`
	Target    string `json:"target"`
	Arguments []struct {
		Messages []struct {
			Text          string    `json:"text"`
			HiddenText    string    `json:"hiddenText"`
			Author        string    `json:"author"`
			CreatedAt     time.Time `json:"createdAt"`
			Timestamp     time.Time `json:"timestamp"`
			MessageID     string    `json:"messageId"`
			MessageType   string    `json:"messageType"`
			Offense       string    `json:"offense"`
			AdaptiveCards []struct {
				Type    string `json:"type"`
				Version string `json:"version"`
				Body    []struct {
					Type    string `json:"type"`
					Inlines []struct {
						Type     string `json:"type"`
						IsSubtle bool   `json:"isSubtle"`
						Italic   bool   `json:"italic"`
						Text     string `json:"text"`
					} `json:"inlines"`
				} `json:"body"`
			} `json:"adaptiveCards"`
			Feedback struct {
				Tag       interface{} `json:"tag"`
				UpdatedOn interface{} `json:"updatedOn"`
				Type      string      `json:"type"`
			} `json:"feedback"`
			ContentOrigin string      `json:"contentOrigin"`
			Privacy       interface{} `json:"privacy"`
			SpokenText    string      `json:"spokenText"`
		} `json:"messages"`
		RequestID string `json:"requestId"`
		Result    struct {
			Value          string `json:"value"`
			Message        string `json:"message"`
			Error          string `json:"error"`
			ServiceVersion string `json:"serviceVersion"`
		} `json:"result"`
	} `json:"arguments"`
}

type BingMessageType2 struct {
	Type         int    `json:"type"`
	InvocationID string `json:"invocationId"`
	Item         struct {
		Messages []struct {
			Text   string `json:"text,omitempty"`
			Author string `json:"author"`
			From   struct {
				ID   string      `json:"id"`
				Name interface{} `json:"name"`
			} `json:"from,omitempty"`
			CreatedAt     time.Time `json:"createdAt"`
			Timestamp     string    `json:"timestamp"`
			Locale        string    `json:"locale,omitempty"`
			Market        string    `json:"market,omitempty"`
			Region        string    `json:"region,omitempty"`
			Location      string    `json:"location,omitempty"`
			LocationHints []struct {
				Country           string `json:"country"`
				CountryConfidence int    `json:"countryConfidence"`
				State             string `json:"state"`
				City              string `json:"city"`
				CityConfidence    int    `json:"cityConfidence"`
				ZipCode           string `json:"zipCode"`
				TimeZoneOffset    int    `json:"timeZoneOffset"`
				Dma               int    `json:"dma"`
				SourceType        int    `json:"sourceType"`
				Center            struct {
					Latitude  float64     `json:"latitude"`
					Longitude float64     `json:"longitude"`
					Height    interface{} `json:"height"`
				} `json:"center"`
				RegionType int `json:"regionType"`
			} `json:"locationHints,omitempty"`
			MessageID string `json:"messageId"`
			RequestID string `json:"requestId"`
			Offense   string `json:"offense"`
			Feedback  struct {
				Tag       interface{} `json:"tag"`
				UpdatedOn interface{} `json:"updatedOn"`
				Type      string      `json:"type"`
			} `json:"feedback"`
			ContentOrigin string      `json:"contentOrigin"`
			Privacy       interface{} `json:"privacy"`
			InputMethod   string      `json:"inputMethod,omitempty"`
			HiddenText    string      `json:"hiddenText,omitempty"`
			MessageType   string      `json:"messageType,omitempty"`
			AdaptiveCards []struct {
				Type    string `json:"type"`
				Version string `json:"version"`
				Body    []struct {
					Type    string `json:"type"`
					Inlines []struct {
						Type     string `json:"type"`
						IsSubtle bool   `json:"isSubtle"`
						Italic   bool   `json:"italic"`
						Text     string `json:"text"`
					} `json:"inlines"`
				} `json:"body"`
			} `json:"adaptiveCards,omitempty"`
			GroundingInfo struct {
				QuestionAnsweringResults []struct {
					Index           string      `json:"index"`
					Title           string      `json:"title"`
					Snippets        []string    `json:"snippets"`
					Data            interface{} `json:"data"`
					Context         interface{} `json:"context"`
					URL             string      `json:"url"`
					LastUpdatedDate interface{} `json:"lastUpdatedDate"`
				} `json:"question_answering_results"`
				WebSearchResults []struct {
					Index    string   `json:"index"`
					Title    string   `json:"title"`
					Snippets []string `json:"snippets"`
					Data     struct {
						Date string `json:"Date"`
					} `json:"data"`
					Context         interface{} `json:"context"`
					URL             string      `json:"url"`
					LastUpdatedDate interface{} `json:"lastUpdatedDate"`
				} `json:"web_search_results"`
			} `json:"groundingInfo,omitempty"`
			SourceAttributions []struct {
				ProviderDisplayName string `json:"providerDisplayName"`
				SeeMoreURL          string `json:"seeMoreUrl"`
				SearchQuery         string `json:"searchQuery"`
				ImageLink           string `json:"imageLink,omitempty"`
				ImageWidth          string `json:"imageWidth,omitempty"`
				ImageHeight         string `json:"imageHeight,omitempty"`
				ImageFavicon        string `json:"imageFavicon,omitempty"`
			} `json:"sourceAttributions,omitempty"`
			SuggestedResponses []struct {
				Text        string    `json:"text"`
				Author      string    `json:"author"`
				CreatedAt   time.Time `json:"createdAt"`
				Timestamp   time.Time `json:"timestamp"`
				MessageID   string    `json:"messageId"`
				MessageType string    `json:"messageType"`
				Offense     string    `json:"offense"`
				Feedback    struct {
					Tag       interface{} `json:"tag"`
					UpdatedOn interface{} `json:"updatedOn"`
					Type      string      `json:"type"`
				} `json:"feedback"`
				ContentOrigin string      `json:"contentOrigin"`
				Privacy       interface{} `json:"privacy"`
			} `json:"suggestedResponses,omitempty"`
			SpokenText string `json:"spokenText,omitempty"`
		} `json:"messages"`
		FirstNewMessageIndex   int         `json:"firstNewMessageIndex"`
		SuggestedResponses     interface{} `json:"suggestedResponses"`
		ConversationID         string      `json:"conversationId"`
		RequestID              string      `json:"requestId"`
		ConversationExpiryTime time.Time   `json:"conversationExpiryTime"`
		Telemetry              struct {
			Metrics   interface{} `json:"metrics"`
			StartTime time.Time   `json:"startTime"`
		} `json:"telemetry"`
		ShouldInitiateConversation bool `json:"shouldInitiateConversation"`
		Result                     struct {
			Value          string `json:"value"`
			Message        string `json:"message"`
			Error          string `json:"error"`
			ServiceVersion string `json:"serviceVersion"`
		} `json:"result"`
	} `json:"item"`
}

type CookieData struct {
	Cookie string `json:"cookie"`
}
