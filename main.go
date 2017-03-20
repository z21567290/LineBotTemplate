// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	//var err2 error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	//events, err2 := strings.Contains(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	/*
		if err2 != nil {
			if err2 == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
	*/

	var IP [2]string

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				//----------------回聲範例---------------------
				/*if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" TextOK!")).Do(); err != nil {
					//發送訊息的格式
					log.Print(err)
				}*/
				//----------------------------------------------------------------------
				//----------------關鍵字回復--------------------
				if strings.Contains(message.Text, "愛你") {
					//IP := event.ReplyToken
					out := fmt.Sprintf("謝謝愛我 ，但LINEBOT依然可惡")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.ReplyToken+" : "+out)).Do()
				} else if strings.Contains(message.Text, "你好") || strings.Contains(message.Text, "妳好") {
					out := fmt.Sprintf("你好")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					//bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+message.PreviewImageURL)).Do()
				} else if strings.Contains(message.Text, "全頻廣播") {
					IP[0] = "6c65c8b36882491faa32493bfeba736"  //飲茶
					IP[1] = "b4c929b7ceec4e21912e6e16304ff0ee" //台南吃吃吃
					//IP[2] ={"6c65c8b36882491faa32493bfeba736","b4c929b7ceec4e21912e6e16304ff0ee"}
					IPP := IP[2]
					bot.PushMessage(IPP, linebot.NewTextMessage("hello")).Do()
				}

				/* else if strings.Contains(message.Text, "rdrrJC") {
					//type:= "image",
					OURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
					PURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(OURL+PURL)).Do()
				}*/
			/*else if strings.Contains(message.Text, "rdrrJC") {
				message.OriginalContentURL = "https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png"
				bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+"ImageOK!")).Do()
			}*/
			/*else {      //回聲功能
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Lucy:"+message.Text+" Aye")).Do() //message.ID
			}*/
			//----------------------------------------------------------------------

			case *linebot.ImageMessage:
				//if message.ID == "RS232.jpg" {
				out := fmt.Sprintf("這是圖片")
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out+message.OriginalContentURL+message.PreviewImageURL)).Do()
				//}
				//out := fmt.Sprintf("這是圖片")
				//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				/*if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.ID+":"+message.OriginalContentURL+message.PreviewImageURL+"ImageOK!")).Do(); err != nil {
					log.Print(err)
				}*/
				//message.OriginalContentURL = "https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png"
				//message.PreviewImageURL = "https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png"
				/*if bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+message.PreviewImageURL+"ImageOK!")).Do() {
					out := fmt.Sprintf("這是圖片")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				}*/

				/*if strings.Contains(message.Text, "rdrrJC") {
				//type:= "image",
				OURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
				PURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
				bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(OURL+PURL+"ImageOK!")).Do()

				}*/

				//--------------------------------------------------------------
				/*
					case *linebot.ImageMessage:
						if _, err2 = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.ID+":"+message.OriginalContentURL+message.PreviewImageURL+"ImageOK!")).Do(); err2 != nil {
							log.Print(err2)
						}
				*/
			}
			//--------------------------------------------------------------- + message.PreviewImageURL
		}
	}
	/*
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.ImageMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.ID+":/n"+message.OriginalContentURL+"/n"+message.PreviewImageURL+"/n OK!")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	*/
}
